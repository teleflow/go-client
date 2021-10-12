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
	Id            int64             `json:"id,omitempty"`
	Phone         string            `json:"phone,omitempty"`
	CampaignId    int64             `json:"campaign_id,omitempty"`
	Params        map[string]string `json:"params,omitempty"`
	ResultVars    map[string]string `json:"result_vars,omitempty"`
	ExtId         string            `json:"ext_id,omitempty"`
	Result        int64             `json:"result,omitempty"`
	IsCompleted   bool              `json:"is_completed,omitempty"`
	IsListened    bool              `json:"is_listened,omitempty"`
	InProcess     bool              `json:"in_process,omitempty"`
	CompletedAt   time.Time         `json:"completed_at,omitempty"`
	LastCall      *Call             `json:"last_call,omitempty"`
	NextCallTime  time.Time         `json:"next_call_time,omitempty"`
	AttemptsCount int64             `json:"attempts_count,omitempty"`
}

type Call struct {
	Id              int64         `json:"id,omitempty"`
	Task            *AutocallTask `json:"task,omitempty"`
	Status          int64         `json:"status,omitempty"`
	StatusChangedAt time.Time     `json:"status_changed_at"`
	Result          int64         `json:"result,omitempty"`
	Duration        float64       `json:"duration,omitempty"`
	Comment         string        `json:"comment,omitempty"`
	Priority        int64         `json:"priority,omitempty"`
	ResultVariables string        `json:"result_variables,omitempty"`
	CompletedAt     time.Time     `json:"completed_at"`
}
