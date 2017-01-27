package stash

// If I decide to build an indexer, this is where it will live. For now, this
// just contains some relevant quotes from the official documentation.

// From the docs:
// When a change is made to a stash, the entire stash is sent in an update.
// If you wish to track historical items, you will need to compare the previous
// items in the stash to the current items in the stash, otherwise you can
// simply delete any items matching the stash id and insert the new items.
// Items don't currently have UIDs. Calculate UIDs based on the stash tab
// location.

func updateItems(sr *StashResponse) error {
	return nil
}
