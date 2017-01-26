package stash

import (
	"fmt"
	"testing"
)

func TestStashRetrieval(t *testing.T) {
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
