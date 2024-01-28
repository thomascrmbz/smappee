package unlicensed

import "net/http"

type Client struct {
	Username string
	Password string

	HttpClient *http.Client

	accessToken  string
	refreshToken string
}
