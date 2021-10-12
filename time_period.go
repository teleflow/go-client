package teleflow

import (
	"net/http"
)

func (c *Client) GetTimePeriods() ([]*TimePeriod, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/time-periods", nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseTimePeriods{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

type ResponseTimePeriods struct {
	Data []*TimePeriod `json:"data,omitempty"`
}

type TimePeriod struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
}
