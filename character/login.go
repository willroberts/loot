package character

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	loginUrl = "https://www.pathofexile.com/login"
)

// Credentials contains an email and password for pathofexile.com.
// Authentication via session ID is not yet supported.
// TODO: Try SessionID instead of username/password.
type Credentials struct {
	Email    string
	Password string
}

// authenticate creates a session on pathofexile.com.
func authenticate(c Credentials) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	log.Println("Token:", token)

	queryParams := url.Values{}
	queryParams.Add("login_email", c.Email)
	queryParams.Add("login_password", c.Password)
	queryParams.Add("hash", token)
	queryParams.Add("login", "Login")
	paramReader := bytes.NewBufferString(queryParams.Encode())

	req, err := http.NewRequest("POST", loginUrl, paramReader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	errorMessage := strings.Split(strings.SplitAfter(string(b), "<ul class=\"errors\"><li>")[1], "</li></ul></td></tr>")[0]

	if resp.Status != string(http.StatusFound) {
		log.Println("Error:", errorMessage)
		return errors.New(fmt.Sprintf("authentication failed: %s", "reason"))
	}

	return nil
}

// readCredsFromFile loads an email and password from the .credentials file.
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

// getToken extracts the CSRF token from the login form HTML.
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
	token := strings.SplitAfter(string(b), "name=\"hash\" value=\"")[1][:32]

	return token, nil
}
