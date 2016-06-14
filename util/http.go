package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONClient struct {
	authToken string
	endpoint  string
	client    *http.Client
}

func NewJSONClient(endpoint, token string) *JSONClient {
	return &JSONClient{
		authToken: token,
		endpoint:  endpoint,
		client:    new(http.Client),
	}
}

func (c *JSONClient) Get(uri string, output interface{}) error {
	req, err := http.NewRequest("GET", c.endpoint+uri, nil)
	if err != nil {
		return err
	}

	return c.execute(req, output)
}

func (c *JSONClient) Post(uri string, input interface{}, output interface{}) error {
	body, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.endpoint+uri, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return c.execute(req, output)
}

func (c *JSONClient) execute(req *http.Request, output interface{}) error {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.authToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Unexpected server reply : %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(output)
}
