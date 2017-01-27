# forum

The forum package contains code for retrieving items for a specific forum
thread. It's no longer being developed now that the stash API is available.

## usage

Here's a sample program which retrieves and displays the items in a specific
thread:

```go
package main

import (
	"github.com/willroberts/loot/forum"
)

func main() {
	html, err := forum.Retrieve(1566069)
	// check err

	items, err := forum.Parse(html)
	// check err

	for _, i := range items {
		i.PrintAttributes()
	}
}
```
