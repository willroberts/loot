package forum

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/willroberts/loot/items"
	"golang.org/x/net/html"
)

// Given an HTML string, parses and returns []items.ForumItem
func Parse(h string) ([]items.ForumItem, error) {
	scripts, err := extractScripts(h)
	if err != nil {
		return []items.ForumItem{}, err
	}
	itemScript, err := findItems(scripts)
	if err != nil {
		return []items.ForumItem{}, err
	}
	j, err := extractJson(itemScript)
	if err != nil {
		return []items.ForumItem{}, err
	}
	i, err := parseItems(j)
	if err != nil {
		return []items.ForumItem{}, err
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

// Unmarshal items JSON into items.ForumItem structs
func parseItems(j string) ([]items.ForumItem, error) {
	var v []interface{}
	var itemSlice []items.ForumItem
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range v {
		item := items.ForumItem{}
		switch assertedItem := data.(type) {
		case []interface{}:
			//for i := 0; i < 3; i++ {
			for i, _ := range assertedItem {
				switch assertedAttr := assertedItem[i].(type) {
				case float64:
					item.Id = int64(assertedAttr)
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
func populateAttributes(i items.ForumItem) items.ForumItem {
	i.Attributes = items.ItemAttributes{
		Name:              i.Data["name"].(string),
		Corrupted:         i.Data["corrupted"].(bool),
		ExplicitMods:      toStrings(i.Data["explicitMods"]),
		FlavorText:        toStrings(i.Data["flavourText"]),
		FrameType:         int64(i.Data["frameType"].(float64)),
		Height:            int64(i.Data["h"].(float64)),
		Icon:              i.Data["icon"].(string),
		Identified:        i.Data["identified"].(bool),
		League:            i.Data["league"].(string),
		LockedToCharacter: i.Data["lockedToCharacter"].(bool),
		Properties:        parseProperties(i.Data["properties"]),
		Requirements:      parseRequirements(i.Data["requirements"]),
		SocketedItems:     i.Data["socketedItems"],
		Sockets:           parseSockets(i.Data["sockets"]),
		TalismanTier:      int64(i.Data["talismanTier"].(float64)),
		TypeLine:          i.Data["typeLine"].(string),
		Verified:          i.Data["verified"].(bool),
		Width:             int64(i.Data["w"].(float64)),
	}
	return i
}

// Converts a slice of map[string]interface{} to a slice of items.Socket
func parseSockets(i interface{}) []items.Socket {
	switch ii := i.(type) {
	case []interface{}:
		s := []items.Socket{}
		for _, socketData := range ii {
			socketMap := socketData.(map[string]interface{})
			socket := items.Socket{
				Attribute: socketMap["attr"].(string),
				Group:     int64(socketMap["group"].(float64)),
			}
			s = append(s, socket)
		}
		return s
	default:
		return []items.Socket{}
	}
}

// Converts a slice of map[string]interface{} to a slice of items.Property
func parseProperties(i interface{}) []items.Property {
	if i == nil {
		return []items.Property{}
	}
	switch ii := i.(type) {
	case []interface{}:
		p := []items.Property{}
		for _, PropertyData := range ii {
			propertyMap := PropertyData.(map[string]interface{})
			property := items.Property{
				Name:        propertyMap["name"].(string),
				DisplayMode: int64(propertyMap["displayMode"].(float64)),
				Value:       parseValues(propertyMap["values"]),
			}
			p = append(p, property)
		}
		return p
	default:
		return []items.Property{}
	}
}

// Converts a slice of map[string]interface{} to a slice of items.Requirement
func parseRequirements(i interface{}) []items.Requirement {
	if i == nil {
		return []items.Requirement{}
	}
	switch ii := i.(type) {
	case []interface{}:
		p := []items.Requirement{}
		for _, RequirementData := range ii {
			requirementMap := RequirementData.(map[string]interface{})
			requirement := items.Requirement{
				Name:        requirementMap["name"].(string),
				DisplayMode: int64(requirementMap["displayMode"].(float64)),
				Value:       parseValues(requirementMap["values"]),
			}
			p = append(p, requirement)
		}
		return p
	default:
		return []items.Requirement{}
	}
}

// Converts a slice of slices with varying contents to items.ItemValue
func parseValues(i interface{}) items.ItemValue {
	switch ii := i.(type) {
	case []interface{}:
		for _, z := range ii {
			switch zz := z.(type) {
			case []interface{}:
				return items.ItemValue{
					Text: zz[0].(string),
					Flag: int64(zz[1].(float64)),
				}
			default:
			}
		}
		return items.ItemValue{}
	default:
		return items.ItemValue{}
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
