package stash

import (
	"fmt"
	"testing"
)

var (
	sr *StashesResponse // Reuse stashes response in successive tests.
)

func TestGetLatestChangeId(t *testing.T) {
	fmt.Print("Retrieving latest change ID from poe.ninja...")
	id, err := getLatestChangeId()
	if err != nil {
		fmt.Println("failed to retrieve latest change ID:", err)
		t.Fail()
	}
	if len(id) < 1 {
		fmt.Println("error: next change ID is empty.")
		t.Fail()
	}
	fmt.Println("OK")
}

func TestGetStashesPage(t *testing.T) {
	fmt.Print("Retrieving one page of stashes...")
	testChangeId := "40866453-43597087-40520850-47274086-44021477"
	var err error
	sr, err = getStashes(testChangeId)
	if err != nil {
		fmt.Println("failed to retrieve stashes:", err)
		t.Fail()
	}
	fmt.Println("OK")
}

func TestItemParsing(t *testing.T) {
	fmt.Print("Testing item parsing...")
	_ = countItems(sr)
	fmt.Println("OK")
}

// TODO: Limit this test so it doesn't run forever. 10 seconds?
func TestPolling(t *testing.T) {
	fmt.Print("Testing polling...")
	//err := poll()
	//if err != nil {
	//	log.Println("Error:", err)
	//	t.Fail()
	//}
	fmt.Println("OK")
}
