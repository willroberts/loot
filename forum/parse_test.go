package forum

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLoadFixture(t *testing.T) {
	fmt.Printf("Testing fixture loading...")
	_, err := ioutil.ReadFile("fixtures/shop.html")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestExtractScripts(t *testing.T) {
	fmt.Printf("Testing script extraction...")
	htmlFixture, _ := ioutil.ReadFile("fixtures/shop.html")
	_, err := extractScripts(string(htmlFixture)) // can find 0 scripts, but should not fail
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestNoScripts(t *testing.T) {
	fmt.Printf("Testing script extraction with bad input...")
	htmlFixture, err := ioutil.ReadFile("fixtures/noscripts.html")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	_, err = extractScripts(string(htmlFixture))
	if err != nil {
		if err.Error() == "no scripts found" {
			fmt.Println("OK") // detected properly
			return
		} else {
			fmt.Println(err)
			t.FailNow()
		}
	}
	t.FailNow() // failed to detect
}

func TestFindItems(t *testing.T) {
	fmt.Printf("Testing items script finder...")
	htmlFixture, err := ioutil.ReadFile("fixtures/shop.html")
	scripts, _ := extractScripts(string(htmlFixture))
	_, err = findItems(scripts)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestNoItems(t *testing.T) {
	fmt.Printf("Testing items script finder with bad input...")
	htmlFixture, err := ioutil.ReadFile("fixtures/noitems.html")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	scripts, _ := extractScripts(string(htmlFixture))
	_, err = findItems(scripts)
	if err != nil {
		if err.Error() == "no items found" {
			fmt.Println("OK") // detected properly
			return
		} else {
			fmt.Println(err)
			t.FailNow()
		}
	}
	t.FailNow() // failed to detect
}

func TestExtractJson(t *testing.T) {
	fmt.Printf("Testing JSON extraction...")
	htmlFixture, _ := ioutil.ReadFile("fixtures/shop.html")
	scripts, _ := extractScripts(string(htmlFixture))
	s, _ := findItems(scripts)
	_, err := extractJson(s)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestInvalidJson(t *testing.T) {
	fmt.Printf("Testing JSON extraction with bad input...")
	htmlFixture, _ := ioutil.ReadFile("fixtures/shop.html")
	scripts, _ := extractScripts(string(htmlFixture))
	s, _ := findItems(scripts)
	s = fmt.Sprintf("%sasdf,,,,", s)
	_, err := extractJson(s)
	if err != nil {
		fmt.Println("OK") // detected properly
		return
	}
	t.FailNow() // failed to detect
}

func TestExtractItems(t *testing.T) {
	fmt.Printf("Testing item parsing...")
	htmlFixture, _ := ioutil.ReadFile("fixtures/shop.html")
	scripts, _ := extractScripts(string(htmlFixture))
	s, _ := findItems(scripts)
	j, _ := extractJson(s)
	_, err := parseItems(j)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}
