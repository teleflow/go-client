package teleflow

import (
	"net/http"
	"time"
)

func (c *Client) GetAutocallCampaigns() ([]*AutocallCampaign, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/campaigns", nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallCampaigns{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) SetWebhook(campaignId int64, url string) (*WebHook, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, "auto-call/webhook", WebHook{
		CampaignId: campaignId,
		WebhookUrl: url,
	})
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseSetWebhook{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) GetWebhook(campaignId int64) (*WebHook, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/webhook", WebHook{
		CampaignId: campaignId,
	})
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseGetWebhook{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

type ResponseSetWebhook struct {
	Data *WebHook `json:"data,omitempty"`
}

type ResponseGetWebhook struct {
	Data *WebHook `json:"data,omitempty"`
}

type WebHook struct {
	CampaignId int64  `json:"campaign_id"`
	WebhookUrl string `json:"webhook_url"`
}
type ResponseAutocallCampaigns struct {
	Data []*AutocallCampaign `json:"data,omitempty"`
}

type ResponseAutocallCampaign struct {
	Data *AutocallCampaign `json:"data"`
}

type AutocallCampaign struct {
	Id               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	Flow             *Flow     `json:"flow,omitempty"`
	IsActive         bool      `json:"is_active,omitempty"`
	Description      string    `json:"description,omitempty"`
	DateStartAt      time.Time `json:"date_start_at"`
	DateFinishAt     time.Time `json:"date_finish_at"`
	CompletedPercent int       `json:"completed_percent,omitempty"`
}

var (
	ResultComment = map[int64]string{
		1: "Answered",
		2: "Busy",
		3: "No answered",
		4: "Failed",
	}
)
