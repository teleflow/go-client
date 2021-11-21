package teleflow

import (
	"fmt"
	"net/http"
	"time"
)

func (c *Client) GetCampaigns() ([]*AutocallCampaign, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/campaigns", nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallCampaigns{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) GetCampaign(campaignId int64) (*AutocallCampaign, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("auto-call/campaigns/%d", campaignId), nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallCampaign{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) AddCampaign(campaign *AutocallCampaign) (*AutocallCampaign, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, "auto-call/campaigns", campaign)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallCampaign{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) UpdateCampaign(campaign *AutocallCampaign) (*AutocallCampaign, *ApiError, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("auto-call/campaigns/%d", campaign.Id), campaign)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallCampaign{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) PartialUpdateCampaign(campaign *AutocallCampaign) (*AutocallCampaign, *ApiError, error) {
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("auto-call/campaigns/%d", campaign.Id), campaign)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallCampaign{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) DeleteCampaign(campaign *AutocallCampaign) (*AutocallCampaign, *ApiError, error) {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("auto-call/campaigns/%d", campaign.Id), campaign)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseAutocallCampaign{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) SetWebhook(campaignId int64, url string) (*WebHook, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("auto-call/campaigns/%d/webhook", campaignId), WebHook{
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
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("auto-call/campaigns/%d/webhook", campaignId), nil)
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
	Id                        int64         `json:"id,omitempty"`
	Name                      string        `json:"name,omitempty"`
	CreatedAt                 time.Time     `json:"created_at"`
	Flow                      *Flow         `json:"flow,omitempty"`
	IsActive                  bool          `json:"is_active,omitempty"`
	Description               string        `json:"description,omitempty"`
	DateStartAt               time.Time     `json:"date_start_at"`
	DateFinishAt              time.Time     `json:"date_finish_at"`
	CompletedPercent          int           `json:"completed_percent,omitempty"`
	NumberDialingAttempts     int64         `json:"number_dialing_attempts,omitempty"`
	IntervalBetweenAttempts   int64         ` json:"interval_between_attempts,omitempty"`
	DialingTimeout            int64         `json:"dialing_timeout,omitempty"`
	IntervalForListenedStatus int64         `json:"interval_for_listened_status,omitempty"`
	IncreaseIntervalDialing   int64         `json:"increase_interval_dialing,omitempty"`
	TaskLifetime              int64         `json:"task_lifetime,omitempty"`
	PhoneLine                 string        `json:"phone_line,omitempty"`
	TimeZone                  *TimeZone     `json:"time_zone,omitempty"`
	WebHook                   string        `json:"web_hook,omitempty"`
	TimePeriods               []*TimePeriod `json:"time_periods,omitempty"`
}

type TimeZone struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Offset int64  `json:"offset,omitempty"`
}

var (
	ResultComment = map[int64]string{
		1: "Answered",
		2: "Busy",
		3: "No answered",
		4: "Failed",
	}
)
