package smappee

import "net/http"

type Smappee struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string

	client *http.Client
}

func NewSmappee(clientID string, clientSecret string, username string, password string) *Smappee {
	return &Smappee{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		client:       http.DefaultClient,
	}
}
