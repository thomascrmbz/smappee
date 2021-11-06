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

// NewSmappee creates a new smappee client
func NewSmappee(clientID string, clientSecret string, username string, password string) (*Smappee, error) {
	smappee := &Smappee{
		clientID:     clientID,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
		client:       http.DefaultClient,
	}

	return smappee, smappee.authenticate()
}
