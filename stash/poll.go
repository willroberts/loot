package stash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	// Base URL for the official public stash tab API.
	baseUrl string = "http://www.pathofexile.com/api/public-stash-tabs"

	// Third-party poe.ninja API for retrieving the most recent change ID.
	ninjaUrl string = "http://poeninja.azureedge.net/api/Data/GetStats"
)

// ninjaResponse contains the response from the poe.ninja API.
type ninjaResponse struct {
	Id                 int
	NextChangeId       string
	ApiBytesDownloaded int
	StashTabsProcessed int
	ApiCalls           int
}

// getLatestChangeId retrieves the most recent change ID from the poe.ninja API
// so we can skip to the present. If we build an indexer, we'll need to start
// from the first page instead.
func getLatestChangeId() (string, error) {
	resp, err := http.Get(ninjaUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var n ninjaResponse
	err = json.Unmarshal(b, &n)
	if err != nil {
		return "", err
	}

	return n.NextChangeId, nil
}

// getStashes retrieves a single set of stashes or changes.
func getStashes(nextChangeId string) (*StashesResponse, error) {
	url := baseUrl
	if nextChangeId != "" {
		url = fmt.Sprintf("%s?id=%s", baseUrl, nextChangeId)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var s StashesResponse
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Poll begins requesting stashes from the API starting from the most recent
// change ID. It does not stop unless it is interrupted.
// TODO: End loop on SIGINT or SIGTERM.
// TODO: Store items in Redis with a short TTL?
func Poll() error {
	var s *StashesResponse
	var err error

	next, err := getLatestChangeId()
	if err != nil {
		return err
	}
	log.Println("Starting change ID:", next)

	for {
		s, err = getStashes(next)
		if err != nil {
			return err
		}
		log.Printf("Found %d items.", countItems(s))

		next = s.NextChangeId
		log.Println("Next change ID:", next)
		if err != nil {
			return err
		}

		// Always wait one second between requests. The API documentation suggests
		// waiting one second when encountering a page with no stashes (indicating
		// that we're caught up), but we wait one second after every page in order
		// to avoid being throttled or blacklisted.
		time.Sleep(1 * time.Second)
	}
}

// countItems iterates over the stashes in a specific response, then sums and
// returns the number of items included in it.
func countItems(sr *StashesResponse) int {
	var count int
	for _, s := range sr.Stashes {
		count += len(s.Items)
	}
	return count
}
