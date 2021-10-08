package teleflow

import (
	"fmt"
	"net/http"
	"time"
)

func (c *Client) GetMediaFiles() ([]*MediaFile, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, "auto-call/media-files", nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseMediaFiles{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) GetMediaFile(mediaFileId int64) (*MediaFile, *ApiError, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("auto-call/media-files/%d", mediaFileId), nil)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseMediaFile{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) AddMediaFile(mediaFile *MediaFile) (*MediaFile, *ApiError, error) {
	req, err := c.newRequest(http.MethodPost, "auto-call/media-files", mediaFile)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseMediaFile{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) UpdateMediaFile(mediaFile *MediaFile) (*MediaFile, *ApiError, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("auto-call/media-files/%d", mediaFile.Id), mediaFile)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseMediaFile{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) UploadMediaFile(mediaFile *MediaFile, filePath string) (*MediaFile, *ApiError, error) {
	req, err := c.newUploadRequest(http.MethodPost,
		fmt.Sprintf("auto-call/media-files/%d/upload", mediaFile.Id),
		filePath,
		"file",
	)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseMediaFile{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

func (c *Client) DeleteMediaFile(mediaFile *MediaFile) (*MediaFile, *ApiError, error) {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("auto-call/media-files/%d", mediaFile.Id), mediaFile)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseMediaFile{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}

type ResponseMediaFiles struct {
	Data []*MediaFile `json:"data,omitempty"`
}

type ResponseMediaFile struct {
	Data *MediaFile `json:"data"`
}

type MediaFile struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Path      string    `json:"path,omitempty"`
	Size      int64     `json:"size,omitempty"`
	Seconds   int64     `json:"seconds,omitempty"`
	Ext       string    `json:"ext,omitempty"`
}
