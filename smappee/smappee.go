package smappee

import "net/http"

type Smappee struct {
	clientID     string
	clientSecret string
	username     string
	password     string

	accessToken  string
	refreshToken string

	client *http.Client
}

func NewSmappee(clientID string, clientSecret string, username string, password string) *Smappee {
	smappee := &Smappee{
		clientID:     clientID,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
		client:       http.DefaultClient,
	}

	smappee.authenticate()

	return smappee
}
