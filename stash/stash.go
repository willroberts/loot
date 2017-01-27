package stash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// Base URL for the official public stash tab API.
	baseUrl string = "http://www.pathofexile.com/api/public-stash-tabs"

	// Third-party poe.ninja API for retrieving the most recent change ID.
	ninjaUrl string = "http://poeninja.azureedge.net/api/Data/GetStats"
)

// stashesResponse contains the response from the official public stash tab API.
type stashesResponse struct {
	NextChangeId string `json:"next_change_id"`
	Stashes      []Stash
}

// Stash represents a stash tab containing items.
type Stash struct {
	AccountName       string
	LastCharacterName string
	Id                string
	Stash             string // Stash tab label, where the global price can be stored.
	Items             []Item
	Public            bool
}

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
func getStashes(nextChangeId string) (*stashesResponse, error) {
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

	var s stashesResponse
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// CountItems iterates over the stashes in a specific response, then sums and
// returns the number of items included in it.
func CountItems(sr *stashesResponse) int {
	var count int
	for _, s := range sr.Stashes {
		count += len(s.Items)
	}
	return count
}
