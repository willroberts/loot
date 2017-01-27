package forum

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Given an HTML string, parses and returns []Item
func Parse(h string) ([]Item, error) {
	scripts, err := extractScripts(h)
	if err != nil {
		return []Item{}, err
	}
	itemScript, err := findItems(scripts)
	if err != nil {
		return []Item{}, err
	}
	j, err := extractJson(itemScript)
	if err != nil {
		return []Item{}, err
	}
	i, err := parseItems(j)
	if err != nil {
		return []Item{}, err
	}
	return i, nil
}

// Given an HTML string, find all <script> blocks and return []string
// Returns an error when no script blocks are found
func extractScripts(h string) ([]string, error) {
	var scripts []string
	z := html.NewTokenizer(bytes.NewBufferString(h))
	for {
		tokenType := z.Next()
		if tokenType == html.ErrorToken {
			break
		}
		if tokenType == html.StartTagToken {
			if z.Token().Data == "script" {
				z.Next()
				scripts = append(scripts, z.Token().Data)
			}
		}
	}
	if len(scripts) == 0 {
		return nil, errors.New("no scripts found")
	}
	return scripts, nil
}

// Given a slice of <script> strings, return the one containing "PoE/Item/DeferredItemRenderer"
func findItems(scripts []string) (string, error) {
	if len(scripts) == 0 {
		return "", errors.New("no scripts supplied")
	}
	for _, s := range scripts {
		if strings.Contains(s, "PoE/Item/DeferredItemRenderer") {
			return s, nil
		}
	}
	return "", errors.New("no items found")
}

// Given a <script> block containing items, extract and return the JSON items block
// Uses hardcoded strings to anticipate the text before and after the JSON block
// Pretty terrible, but avoids regular expressions for now
// Will update when the updated PoE API is released
// Returns an error when invalid JSON is generated
func extractJson(script string) (string, error) {
	script = strings.Replace(script, "//<!--", "", 1)
	script = strings.Replace(script, "require([\"main\"], function() {", "", 1)
	script = strings.Replace(script, "require([\"PoE/Item/DeferredItemRenderer\"], function(R) { (new R(", "", 1)
	script = strings.Replace(script, ")).run(); });", "", 1)
	script = strings.Replace(script, "});", "", 1)
	script = strings.Replace(script, "//-->", "", 1)
	j := strings.Trim(script, "\n ")
	var v interface{}
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		return "", err
	}
	return j, nil
}

// Unmarshal items JSON into Item structs
func parseItems(j string) ([]Item, error) {
	var v []interface{}
	var itemSlice []Item
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range v {
		item := Item{}
		switch assertedItem := data.(type) {
		case []interface{}:
			//for i := 0; i < 3; i++ {
			for i, _ := range assertedItem {
				switch assertedAttr := assertedItem[i].(type) {
				case float64:
					item.Id = int(assertedAttr)
				case map[string]interface{}:
					item.Data = assertedAttr
				case []interface{}:
					item.Extra = assertedAttr
				default:
				}
			}
		}
		itemSlice = append(itemSlice, populateAttributes(item))
	}
	return itemSlice, nil
}

// Replace map[string]interface{} with ItemAttributes in Item structs
func populateAttributes(i Item) Item {
	i.Attributes = ItemAttributes{
		Name:              i.Data["name"].(string),
		Corrupted:         i.Data["corrupted"].(bool),
		ExplicitMods:      toStrings(i.Data["explicitMods"]),
		FlavorText:        toStrings(i.Data["flavourText"]),
		FrameType:         int(i.Data["frameType"].(float64)),
		Height:            int(i.Data["h"].(float64)),
		Icon:              i.Data["icon"].(string),
		Identified:        i.Data["identified"].(bool),
		League:            i.Data["league"].(string),
		LockedToCharacter: i.Data["lockedToCharacter"].(bool),
		Properties:        parseProperties(i.Data["properties"]),
		Requirements:      parseRequirements(i.Data["requirements"]),
		SocketedItems:     i.Data["socketedItems"],
		Sockets:           parseSockets(i.Data["sockets"]),
		TalismanTier:      int(i.Data["talismanTier"].(float64)),
		TypeLine:          i.Data["typeLine"].(string),
		Verified:          i.Data["verified"].(bool),
		Width:             int(i.Data["w"].(float64)),
	}
	return i
}

// Converts a slice of map[string]interface{} to a slice of Sockets
func parseSockets(i interface{}) []Socket {
	switch ii := i.(type) {
	case []interface{}:
		s := []Socket{}
		for _, socketData := range ii {
			socketMap := socketData.(map[string]interface{})
			socket := Socket{
				Attribute: socketMap["attr"].(string),
				Group:     int(socketMap["group"].(float64)),
			}
			s = append(s, socket)
		}
		return s
	default:
		return []Socket{}
	}
}

// Converts a slice of map[string]interface{} to a slice of Properties
func parseProperties(i interface{}) []Property {
	if i == nil {
		return []Property{}
	}
	switch ii := i.(type) {
	case []interface{}:
		p := []Property{}
		for _, PropertyData := range ii {
			propertyMap := PropertyData.(map[string]interface{})
			property := Property{
				Name:        propertyMap["name"].(string),
				DisplayMode: int(propertyMap["displayMode"].(float64)),
				Value:       parseValues(propertyMap["values"]),
			}
			p = append(p, property)
		}
		return p
	default:
		return []Property{}
	}
}

// Converts a slice of map[string]interface{} to a slice of Requirements
func parseRequirements(i interface{}) []Requirement {
	if i == nil {
		return []Requirement{}
	}
	switch ii := i.(type) {
	case []interface{}:
		p := []Requirement{}
		for _, RequirementData := range ii {
			requirementMap := RequirementData.(map[string]interface{})
			requirement := Requirement{
				Name:        requirementMap["name"].(string),
				DisplayMode: int(requirementMap["displayMode"].(float64)),
				Value:       parseValues(requirementMap["values"]),
			}
			p = append(p, requirement)
		}
		return p
	default:
		return []Requirement{}
	}
}

// Converts a slice of slices with varying contents to ItemValues
func parseValues(i interface{}) ItemValue {
	switch ii := i.(type) {
	case []interface{}:
		for _, z := range ii {
			switch zz := z.(type) {
			case []interface{}:
				return ItemValue{
					Text: zz[0].(string),
					Flag: int(zz[1].(float64)),
				}
			default:
			}
		}
		return ItemValue{}
	default:
		return ItemValue{}
	}
}

// Converts a slice of empty interfaces to a slice of strings
// Returns an empty slice of strings when an invalid slice is given
func toStrings(i interface{}) []string {
	s := []string{}
	switch ii := i.(type) {
	case []interface{}:
		for _, u := range ii {
			s = append(s, string(u.(string)))
		}
		return s
	default:
		return []string{}
	}
}
