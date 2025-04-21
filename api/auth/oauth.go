package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const GH_API_URL = "https://github.com/login/oauth"

type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Error       string `json:"error,omitempty"`
}

func getAccessTokenGithub(clientId, clientSecret, code string) (*GithubAccessTokenResponse, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/access_token?client_id=%s&client_secret=%s&code=%s", GH_API_URL, clientId, clientSecret, code),
		nil,
	)
	req.Header.Add("Accept", "application/json")
	if err != nil {
		log.Println("Failed to create request: ", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Failed to send request: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response: ", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println("Failed to get access token: ", string(buf))
		return nil, errors.New("unknown error at getting access token")
	}
	var tresp GithubAccessTokenResponse
	err = json.Unmarshal(buf, &tresp)
	if err != nil {
		return nil, err
	}
	if tresp.Error != "" {
		return nil, errors.New(tresp.Error)
	}
	return &tresp, nil
}
