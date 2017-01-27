package character

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	loginUrl = "http://www.pathofexile.com/login"
)

// Credentials contains a username and password for pathofexile.com.
// Authentication via session ID is not yet supported.
type Credentials struct {
	Username string
	Password string
}

// authenticate creates a session at pathofexile.com.
func authenticate(c Credentials) error {
	token, err := getToken()
	if err != nil {
		return err
	}

	queryParams := url.Values{}
	queryParams.Add("login_email", "")
	queryParams.Add("login_password", "")
	queryParams.Add("hash", token)
	queryParams.Add("login", "Login")

	_ = queryParams

	return nil
}

// readCredsFromFile loads a username and password from the .credentials file.
func readCredsFromFile() (Credentials, error) {
	b, err := ioutil.ReadFile(".credentials")
	if err != nil {
		return Credentials{}, err
	}

	var c Credentials
	err = json.Unmarshal(b, &c)
	if err != nil {
		return Credentials{}, err
	}

	return c, nil
}

// getToken extracts the CSRF token from the login form.
func getToken() (string, error) {
	resp, err := http.Get(loginUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	parts := strings.SplitAfter(string(b), "name=\"hash\" value=\"")
	token := parts[1][:32]

	return token, nil
}
