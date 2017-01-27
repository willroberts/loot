package stash

import (
	"log"
	"time"
)

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
