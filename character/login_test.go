package character

import (
	"fmt"
	"testing"
)

func TestReadCredsFromFile(t *testing.T) {
	fmt.Print("Testing credential reading...")
	_, err := readCredsFromFile()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("You'll need to create a .credentials file.")
		fmt.Println("See the template at credentials.template.")
		t.Fail()
	}
	fmt.Println("OK")
}
