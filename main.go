// Sample app showing retrieval and display of items in the given thread
package main

import (
	"flag"

	"github.com/willroberts/loot/forum"
)

var (
	threadId int
)

func init() {
	flag.IntVar(&threadId, "thread", 1566069, "Thread ID to show items from")
	flag.Parse()
}

func main() {
	h, err := forum.Retrieve(threadId)
	if err != nil {
		panic(err)
	}
	f, err := forum.Parse(h)
	if err != nil {
		panic(err)
	}
	for _, i := range f {
		i.PrintAttributes()
	}
}
