package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/willroberts/loot/items"
)

type StashesResponse struct {
	NextChangeId string `json:"next_change_id"`
	Stashes      []Stash
}

type Stash struct {
	AccountName       string
	LastCharacterName string
	Id                string
	// Note for the Stash, where global price is stored
	Stash  string
	Items  []items.Item
	Public bool
}

func getPublicStashTabs(nextChangeId string) (*StashesResponse, error) {
	url := "http://www.pathofexile.com/api/public-stash-tabs"
	if nextChangeId != "" {
		url = fmt.Sprintf("http://www.pathofexile.com/api/public-stash-tabs?id=%s", nextChangeId)
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

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

// When a change is made to a stash, the entire stash is sent in an update.
// If you wish to track historical items, you will need to compare the previous
// items in the stash to the current items in the stash, otherwise you can
// simply delete any items matching the stash id and insert the new items.
// Items don't currently have UIDs. Calculate UIDs based on the stash tab
// location.
func updateStash() {

}

// If there are no changes, this page will show as empty
// You need to simply retry this same page no quicker than every one second
// Eventually, when someone makes a change, this page will return content
// You should process this content just like the original page, load up the
// next_change_id, then start watching the next one for changes
func pollPublicStashTabs() {

}

func main() {
	sr, err := getPublicStashTabs("foo")
	if err != nil {
		panic(err)
	}

	for _, s := range sr.Stashes {
		// debug printing -- replace with actual storage
		for _, i := range s.Items {
			i.Print()
			fmt.Println()
		}
	}
}
