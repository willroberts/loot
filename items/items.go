package items

import "fmt"

type ForumItem struct {
	// The order in which the item appears in its thread
	Id int64

	// The raw JSON data for the item
	Data map[string]interface{}

	// A parsed version of the JSON data
	Attributes ItemAttributes

	// An empty array; purpose unknown
	Extra []interface{}
}

type ItemAttributes struct {
	// The name of the item
	Name string

	// Whether or not the item is corrupted
	Corrupted bool

	// The mods of the item as a slice of strings
	ExplicitMods []string

	// The flavor text of the item, where each element is a distinct line of text
	FlavorText []string

	// Not sure what this is for yet
	FrameType int64

	// The height of the item in stash spaces
	Height int64

	// The image URL for the item's icon
	Icon string

	// Whether or not this item is identified
	Identified bool

	// The league in which this item resides
	League string

	// Whether or not this item is locked to the character
	LockedToCharacter bool

	// A slice of item.Property
	Properties []Property

	// A slice of item.Requirement
	Requirements []Requirement

	// Not sure what this looks like yet
	SocketedItems interface{}

	// A slice of item.Socket
	Sockets []Socket

	// Whether or not this is a support gem (?)
	Support bool

	// For talismans, displays the tier (1-4)
	TalismanTier int64

	// Item base type
	TypeLine string

	// Whether or not the item is still owned by the user
	Verified bool

	// Width of the item in stash spaces
	Width int64
}

func (i *ForumItem) PrintAttributes() {
	fmt.Println("ID:", i.Id)
	fmt.Println("Name:", i.Attributes.Name)
	fmt.Println("Base Type:", i.Attributes.TypeLine)
	fmt.Println("Properties:", i.Attributes.Properties)
	fmt.Println("Mods:", i.Attributes.ExplicitMods)
	fmt.Println("Requirements:", i.Attributes.Requirements)
	fmt.Println("Sockets:", i.Attributes.Sockets)
	fmt.Println()
}

type Property struct {
	// The name of the property
	Name string

	// Whether or not the requirement is displayed on the website
	DisplayMode int64

	// The value for this property
	Value ItemValue
}

type Requirement struct {
	// The name of the requirement
	Name string

	// Whether or not the requirement is displayed on the website
	DisplayMode int64

	// Contains two integers: the value itself, and an unknown value (usually 0)
	Value ItemValue
}

type Socket struct {
	// One of S, D, or I for Strength, Dexterity, or Intelligence
	Attribute string

	// Numeric socket group. Sockets with the same group ID are linked
	Group int64
}

type ItemValue struct {
	// The actual value
	Text string

	// Purpose unknown
	Flag int64
}
