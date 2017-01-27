// Runs the poller for the public stash tab API.
package main

import (
	"log"

	"github.com/willroberts/loot/stash"
)

func main() {
	err := stash.Poll()
	if err != nil {
		log.Fatalln(err)
	}
}
