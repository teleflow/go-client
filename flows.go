package teleflow

import (
	"fmt"
	"net/http"
	"time"
)

func (c *Client) GetFlows() ([]*Flow, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/flows", nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlows{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) GetFlow(flowId int64) (*Flow, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("auto-call/flows/%d", flowId), nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlow{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) AddFlow(flow *Flow) (*Flow, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, "auto-call/flows", flow)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlow{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) UpdateFlow(flow *Flow) (*Flow, *ApiError, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("auto-call/flows/%d", flow.Id), flow)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlow{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) PartialUpdateFlow(flow *Flow) (*Flow, *ApiError, error) {
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("auto-call/flows/%d", flow.Id), flow)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlow{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) DeleteFlow(flow *Flow) (*Flow, *ApiError, error) {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("auto-call/flows/%d", flow.Id), flow)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlow{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

/////////////// steps

func (c *Client) GetFlowSteps(flowId int64) ([]*FlowStep, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("auto-call/flows/%d/steps", flowId), nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlowSteps{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) GetFlowStep(stepId int64) (*FlowStep, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("auto-call/flows/steps/%d", stepId), nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlowStep{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) AddFlowStep(flowStep *FlowStep) (*FlowStep, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("auto-call/flows/%d/steps", flowStep.Flow.Id), flowStep)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlowStep{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) UpdateFlowStep(flowStep *FlowStep) (*FlowStep, *ApiError, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("auto-call/flows/steps/%d", flowStep.Id), flowStep)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlowStep{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) PartialUpdateFlowStep(flowStep *FlowStep) (*FlowStep, *ApiError, error) {
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("auto-call/flows/steps/%d", flowStep.Id), flowStep)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlowStep{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) MoveFlowStep(flowStep *FlowStep, parentFlowStep *FlowStep) (*FlowStep, *ApiError, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("auto-call/flows/steps/%d/move-to/%d", flowStep.Id, parentFlowStep.Id), nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlowStep{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) DeleteFlowStep(flowStep *FlowStep) (*FlowStep, *ApiError, error) {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("auto-call/flows/steps/%d", flowStep.Id), flowStep)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseFlowStep{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

type ResponseFlows struct {
	Data []*Flow `json:"data,omitempty"`
}
type ResponseFlowSteps struct {
	Data []*FlowStep `json:"data,omitempty"`
}

type ResponseFlow struct {
	Data *Flow `json:"data"`
}
type ResponseFlowStep struct {
	Data *FlowStep `json:"data"`
}

type Flow struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsActive  bool      `json:"is_active,omitempty"`
}

type FlowStep struct {
	Id              int64      `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Flow            *Flow      `json:"flow,omitempty"`
	Parent          *FlowStep  `json:"parent,omitempty"`
	Root            *FlowStep  `json:"root,omitempty"`
	StepType        int64      `json:"step_type,omitempty"`
	Sound           *MediaFile `json:"sound,omitempty"`
	TtsText         string     `json:"tts_text,omitempty"`
	SpeechSourceId  int64      `json:"speech_source_id,omitempty"`
	DigitsCount     int64      `json:"digits_count,omitempty"`
	WaitSeconds     int64      `json:"wait_seconds,omitempty"`
	ApiUrl          string     `json:"api_url,omitempty"`
	GoToExtension   string     `json:"go_to_extension,omitempty"`
	GoToContext     string     `json:"go_to_context,omitempty"`
	GoToPriority    string     `json:"go_to_priority,omitempty"`
	GoToDescription string     `json:"go_to_description,omitempty"`
	GoToFlowStep    *FlowStep  `json:"go_to_flow_step,omitempty"`
	Position        int64      `json:"position,omitempty"`
	Condition       string     `json:"condition,omitempty"`
	Variable        string     `json:"variable,omitempty"`
}
