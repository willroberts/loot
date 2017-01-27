package stash

import "github.com/willroberts/loot/stash/items"

type StashesResponse struct {
	NextChangeId string `json:"next_change_id"`
	Stashes      []Stash
}

type Stash struct {
	AccountName       string
	LastCharacterName string
	Id                string
	Stash             string // Stash tab label, where the global price can be stored.
	Items             []items.Item
	Public            bool
}
