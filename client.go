package teleflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	clientVersion = "0.1a"
)

type Client struct {
	config     *Config
	server     *url.URL
	httpClient *http.Client
	apiVersion string
	timeout    time.Duration
}

type Config struct {
	Server      string
	Username    string
	Password    string
	AccessToken string
}

type ApiResponse struct {
	Error *ApiError   `json:"err"`
	Data  interface{} `json:"data"`
}

type ApiError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// newRequest returns new HTTP Request with ..
func (c *Client) newRequest(method, resource string, body interface{}) (*http.Request, error) {

	endpoint := fmt.Sprintf("%s://%s/flow/api/%s/%s/", c.server.Scheme, c.server.Host, c.apiVersion, resource)

	var buf io.Reader
	if body != nil {
		bts, err := json.Marshal(&body)
		if err != nil {
			return nil, fmt.Errorf("marshal request data: %s", err)
		}
		buf = bytes.NewBuffer(bts)
	}

	r, err := http.NewRequest(method, endpoint, buf)
	if err != nil {
		return nil, fmt.Errorf("create http request: %s", err)
	}

	// Specify request timeout
	if c.timeout > 0 {
		ctx, cancel := context.WithTimeout(r.Context(), c.timeout)
		defer cancel()
		r = r.WithContext(ctx)
	}
	return r, nil
}

func (c *Client) execRequest(req *http.Request) (*http.Response, *ApiError, error) {

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "teleflow-client-go/"+clientVersion)
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+c.config.AccessToken)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	// check errors from API
	if response.StatusCode > 299 || response.StatusCode < 200 {
		defer response.Body.Close()

		apiResp := &ApiResponse{}
		err := json.NewDecoder(response.Body).Decode(apiResp)
		if err != nil {
		}

		if apiResp.Error != nil {
			return nil, apiResp.Error, nil
		}
		return nil, nil, fmt.Errorf("some error with http code: %d", response.StatusCode)
	}

	return response, nil, nil
}

func (c *Client) doRequest(req *http.Request, responseObj interface{}) (*ApiError, error) {

	response, apiErr, err := c.execRequest(req)
	if err != nil || apiErr != nil {
		return apiErr, err
	}
	defer response.Body.Close()

	// else JSON decode
	err = json.NewDecoder(response.Body).Decode(responseObj)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// doRequestWith make request and run function after response received
func (c *Client) doRequestWith(req *http.Request, onResponse func(response *http.Response) error) (*ApiError, error) {

	response, apiErr, err := c.execRequest(req)
	if err != nil || apiErr != nil {
		return apiErr, err
	}
	defer response.Body.Close()

	err = onResponse(response)
	return nil, err
}

func NewClient(config *Config) (*Client, error) {

	serverUrl, err := url.Parse(config.Server)
	if err != nil {
		return nil, fmt.Errorf("server param is required: %s", err)
	}

	return &Client{
		config:     config,
		server:     serverUrl,
		apiVersion: "v1",
		httpClient: &http.Client{},
	}, nil
}
