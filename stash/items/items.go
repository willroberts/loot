package items

import (
	"fmt"
	"regexp"
)

type Item struct {
	// The name of the item
	Name string

	// The level of the item
	ItemLevel int64 `json:"ilvl"`

	// Whether or not the item is corrupted
	Corrupted bool

	// The mods of the item as a slice of strings
	ExplicitMods []string

	// The flavor text of the item, where each element is a distinct line of text
	FlavorText []string

	// Not sure what this is for yet
	FrameType int64

	// The height of the item in stash spaces
	Height int64 `json:"h"`

	// The image URL for the item's icon
	Icon string

	// Whether or not this item is identified
	Identified bool

	// The league in which this item resides
	League string

	// Whether or not this item is locked to the character
	LockedToCharacter bool

	// Custom note for the item, where price is stored
	Note string

	// A slice of item.Property
	Properties []Property

	// A slice of item.Requirement
	Requirements []Requirement

	// Not sure what this looks like yet
	// https://github.com/willroberts/loot/blob/master/forum/fixtures/items.json#L1238
	SocketedItems interface{}

	// A slice of item.Socket
	Sockets []Socket

	// Unknown
	Support bool

	// For talismans, displays the tier (1-4)
	TalismanTier int64

	// Item base type
	TypeLine string

	// Whether or not the item is still owned by the user
	Verified bool

	// Width of the item in stash spaces
	Width int64 `json:"w"`

	// Position in Stash
	X int64
	Y int64

	// Tab Number?
	InventoryId string
}

type Property struct {
	// The name of the property
	Name string

	// Whether or not the requirement is displayed on the website
	//DisplayMode int64

	// The value for this property
	Values [][]interface{}
}

type Requirement struct {
	// The name of the requirement
	Name string

	// Whether or not the requirement is displayed on the website
	//DisplayMode int64

	// The value for this requirement
	Values [][]interface{}
}

type Socket struct {
	// One of S, D, or I for Strength, Dexterity, or Intelligence
	Attribute string `json:"attr"`

	// Numeric socket group. Sockets with the same group ID are linked
	Group int64
}

func (i *Item) Print() {
	if i.Name != "" {
		fmt.Printf("%s, ", trimName(i.Name))
	}
	fmt.Println(trimName(i.TypeLine))

	fmt.Println("Note:", i.Note)

	if len(i.Properties) > 0 {
		fmt.Println("Properties:")
		for _, p := range i.Properties {
			fmt.Printf("  %s: ", p.Name)
			if len(p.Values) > 0 {
				fmt.Printf("%s\n", p.Values[0][0])
			}
		}
	}

	if len(i.ExplicitMods) > 0 {
		fmt.Println("Mods:")
		for _, m := range i.ExplicitMods {
			fmt.Printf("  %s\n", m)
		}
	}

	if len(i.Requirements) > 0 {
		fmt.Println("Requirements:")
		for _, r := range i.Requirements {
			fmt.Printf("  %s: ", r.Name)
			if len(r.Values) > 0 {
				fmt.Printf("%s\n", r.Values[0][0])
			}
		}
	}

	if len(i.Sockets) > 0 {
		fmt.Println("Sockets:")
		for _, p := range i.Sockets {
			fmt.Printf("  %d: %s\n", p.Group, p.Attribute)
		}
	}
}

// Some strings contain tags surrounded by << >>. These are tags for
// localization, and they can be discarded when using English.
// Source: https://www.reddit.com/r/pathofexiledev/comments/48i4s1/information_on_the_new_stash_tab_api/d0kib1h/
var localizationTagPattern = regexp.MustCompile("<<.*>>")

func trimName(name string) string {
	b := localizationTagPattern.ReplaceAll([]byte(name), []byte(""))
	return string(b)
}
