package teleflow

import (
	"net/http"
	"time"
)

func (c *Client) GetAutocallTasks() ([]*AutocallTask, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/tasks", nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallTasks{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) AddAutocallTask(task *AutocallTask) (*AutocallTask, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, "auto-call/tasks", task)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallTask{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}
func (c *Client) AddAutocallTasks(tasks []*AutocallTask) ([]*AutocallTask, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, "auto-call/tasks-batch", tasks)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallTasks{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

type ResponseAutocallTasks struct {
	Data []*AutocallTask `json:"data,omitempty"`
}

type ResponseAutocallTask struct {
	Data *AutocallTask `json:"data"`
}

type AutocallTask struct {
	Id          int64             `json:"id,omitempty"`
	Phone       string            `json:"phone,omitempty"`
	CampaignId  int64             `json:"campaign_id,omitempty"`
	Params      map[string]string `json:"params,omitempty"`
	ResultVars  map[string]string `json:"result_vars,omitempty"`
	ExtId       string            `json:"ext_id,omitempty"`
	Result      int64             `json:"result"`
	IsCompleted bool              `json:"is_completed"`
	IsListened  bool              `json:"is_listened"`
	CompletedAt time.Time         `json:"completed_at"`
}
