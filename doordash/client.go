// Centralized DoorDash client functionality
package doordash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://openapi.doordash.com/"
)

type (
	Client struct {
		BaseURL *url.URL
		token   string
		client  *http.Client
	}
)

// TODO: Refactor accessKey to a JSON string and process during client creation
func NewClient(token string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		BaseURL: baseURL,
		token:   token,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) NewRequest(method string, subPath string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	url, err := c.BaseURL.Parse(subPath)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Authorization": []string{"Bearer " + c.token},
		"Content-Type":  []string{"application/json"},
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if v != nil {
		decErr := json.NewDecoder(res.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return err
}

func (c *Client) makeRequest(method string, endpoint string, params url.Values, body interface{}, res interface{}) error {
	req, err := c.NewRequest(method, endpoint, body)
	if err != nil {
		return err
	}

	if err = c.Do(req, res); err != nil {
		return err
	}

	return nil
}
