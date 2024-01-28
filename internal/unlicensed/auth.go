package unlicensed

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type authenticateRequest struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type authResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (c *Client) Authenticate() error {
	b, _ := json.Marshal(authenticateRequest{
		Username: c.Username,
		Password: c.Password,
	})
	res, err := c.HttpClient.Post("https://dashboard.smappee.net/dashapi/login", "application/json", bytes.NewReader(b))

	if err := c.handleAuthResponse(res); err != nil {
		return err
	}

	go func() {
		for range time.NewTicker(4 * time.Minute).C {
			c.RefreshToken()
		}
	}()

	return err
}

func (c *Client) RefreshToken() error {
	res, err := c.newRequest(http.MethodPost, "https://dashboard.smappee.net/dashapi/refreshToken", refreshTokenRequest{
		RefreshToken: c.refreshToken,
	})
	if err != nil {
		return err
	}
	return c.handleAuthResponse(res)
}

func (c *Client) handleAuthResponse(res *http.Response) error {
	authResponse := authResponse{}
	err := json.NewDecoder(res.Body).Decode(&authResponse)

	if res.StatusCode != 200 {
		return errors.Join(
			errors.New("could not refresh token"),
			errors.New(res.Status),
		)
	}

	c.accessToken = authResponse.Token
	c.refreshToken = authResponse.RefreshToken

	return err
}
