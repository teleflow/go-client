package teleflow

import (
	"net/http"
)

func (c *Client) GetSpeechSources() ([]*SpeechSource, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/speech-sources", nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseSpeechSources{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

type ResponseSpeechSources struct {
	Data []*SpeechSource `json:"data,omitempty"`
}

type SpeechSource struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
