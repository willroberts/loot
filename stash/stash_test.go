package stash

import (
	"fmt"
	"testing"
)

var (
	id string
	sr *stashesResponse
)

func TestGetLatestChangeId(t *testing.T) {
	fmt.Print("Retrieving latest change ID from poe.ninja...")
	var err error
	id, err = getLatestChangeId()
	if err != nil {
		fmt.Println("failed to retrieve latest change ID:", err)
		t.FailNow()
	}
	if len(id) < 1 {
		fmt.Println("error: next change ID is empty.")
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestGetStashesPage(t *testing.T) {
	fmt.Print("Retrieving one page of stashes...")
	var err error
	sr, err = getStashes(id)
	if err != nil {
		fmt.Println("failed to retrieve stashes:", err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestItemParsing(t *testing.T) {
	fmt.Print("Testing item parsing...")
	_ = CountItems(sr)
	fmt.Println("OK")
}
