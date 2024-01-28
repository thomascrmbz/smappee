package unlicensed

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Client) newRequest(method string, url string, data interface{}, parameters ...url.Values) (*http.Response, error) {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(data)

	req, _ := http.NewRequest(method, url, body)

	for _, param := range parameters {
		req.URL.RawQuery = param.Encode()
	}

	req.Header.Add("Token", c.accessToken)
	req.Header.Add("Content-Type", "application/json")

	return c.HttpClient.Do(req)
}
