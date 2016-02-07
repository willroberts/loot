package forum

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// Given a forum thread ID, returns the response body as a string
// Returns an error when HTML validation fails
func Retrieve(thread int) (string, error) {
	u, err := formatUrl(thread)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	r := bytes.NewReader(b)
	_, err = html.Parse(r)
	if err != nil {
		fmt.Println("html validation failed")
		return "", err
	}
	return string(b), nil
}

// Given a forum thread ID, returns the pathofexile.com URL
// Returns an error when URL validation fails
func formatUrl(thread int) (string, error) {
	u := fmt.Sprintf("http://pathofexile.com/forum/view-thread/%d", thread)
	_, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return u, nil
}
