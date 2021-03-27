package smappee

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (s *Smappee) newRequest(method string, path string) (*http.Response, error) {
	req, _ := http.NewRequest(method, "https://"+host+path, nil)
	res, err := s.client.Do(req)
	return res, err
}

func (s *Smappee) authenticate() {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", s.clientID)
	data.Set("client_secret", s.clientSecret)
	data.Set("username", s.username)
	data.Set("password", s.password)

	res, _ := s.client.PostForm("https://"+host+"/dev/v1/oauth2/token", data)

	tokenResponse := getTokenResponse{}
	json.NewDecoder(res.Body).Decode(&tokenResponse)

	s.accessToken = tokenResponse.AccessToken
	s.refreshToken = tokenResponse.RefreshToken
}

type getTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
