package smappee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
)

func (s *Smappee) newRequest(method string, path string, data interface{}, parameters ...url.Values) (*http.Response, error) {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(data)

	req, _ := http.NewRequest(method, "https://"+host+path, body)

	for _, param := range parameters {
		req.URL.RawQuery = param.Encode()
	}

	req.Header.Add("Authorization", "Bearer "+s.accessToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := s.client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	return res, err
}

func (s *Smappee) authenticate() error {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", s.clientID)
	data.Set("client_secret", s.clientSecret)
	data.Set("username", s.username)
	data.Set("password", s.password)

	res, err := s.client.PostForm("https://"+host+"/dev/v1/oauth2/token", data)

	if res.StatusCode != 200 {
		return ErrorClientConnection(res.Status)
	}

	tokenResponse := getTokenResponse{}
	json.NewDecoder(res.Body).Decode(&tokenResponse)

	s.accessToken = tokenResponse.AccessToken
	s.refreshToken = tokenResponse.RefreshToken

	return err
}

type getTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func sum(list []float64) float64 {
	total := 0.0
	for _, value := range list {
		total += value
	}
	return round(total)
}

func round(x float64) float64 {
	return math.Round(x*10000) / 10000
}

type context struct {
	Smappee         *Smappee
	ServiceLocation *ServiceLocation
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.EqualFold(a, b) {
			return true
		}
	}
	return false
}
