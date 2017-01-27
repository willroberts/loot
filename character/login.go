package character

import (
	"encoding/json"
	"io/ioutil"
)

type Credentials struct {
	Username  string
	Password  string
	SessionId string // Not yet supported.
}

func authenticate(username, password string) {

}

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
