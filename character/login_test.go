package character

import (
	"fmt"
	"testing"
)

var (
	creds Credentials
)

func TestReadCredsFromFile(t *testing.T) {
	fmt.Print("Testing credential reading...")
	var err error
	creds, err = readCredsFromFile()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("You'll need to create a .credentials file.")
		fmt.Println("See the template at credentials.template.")
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestGetToken(t *testing.T) {
	fmt.Print("Testing CSRF token retrieval...")
	_, err := getToken()
	if err != nil {
		fmt.Println("error:", err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestAuthenticate(t *testing.T) {
	fmt.Print("Testing authentication...")
	err := authenticate(creds)
	if err != nil {
		t.FailNow()
	}
	fmt.Println("OK")
}
