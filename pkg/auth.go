package qtapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	authURL         = "https://login.questrade.com/oauth2/token?grant_type=refresh_token&refresh_token="
	practiceAuthURL = "https://practicelogin.questrade.com/oauth2/token?grant_type=refresh_token&refresh_token="
)

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	APIServer    string `json:"api_server"`
	Expiry       int32  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func (q Questrade) RefreshToken() (Auth, error) {
	client := &http.Client{}

	url := fmt.Sprintf(authURL, q.token)
	resp, err := client.Get(url)
	if err != nil {
		return Auth{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Auth{}, fmt.Errorf("Questrade token refresh failure.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Auth{}, err
	}

	qauth := Auth{}
	err = json.Unmarshal(body, &qauth)
	if err != nil {
		return Auth{}, err
	}

	return qauth, nil

}
