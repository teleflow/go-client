package teleflow

import (
	"io"
	"net/http"
	"os"
	"time"
)

type RecordOptions struct {
	Src           string    `json:"src,omitempty"`
	Dst           string    `json:"dst,omitempty"`
	LinkedId      string    `json:"linked_id,omitempty"`
	UniqId        string    `json:"uniq_id,omitempty"`
	RecordingFile string    `json:"recording_file,omitempty"`
	TimeStart     time.Time `json:"time_start"`
	TimeEnd       time.Time `json:"time_end"`
}

type ResponseRecords struct {
	Data []*Record `json:"data,omitempty"`
}

type Record struct {
	LinkedId      string    `json:"linkedid,omitempty"`
	UniqId        string    `json:"uniqid,omitempty"`
	RecordingFile string    `json:"recording_file,omitempty"`
	Sec           int32     `json:"sec,omitempty"`
	Src           string    `json:"src,omitempty"`
	Dst           string    `json:"dst,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

// DownloadAudioFile download audio file and save by savePath
func (c *Client) DownloadAudioFile(option *RecordOptions, savePath string) (*ApiError, error) {

	req, err := c.newRequest(http.MethodPost, "record/download", option)
	if err != nil {
		return nil, err
	}
	apiErr, err := c.doRequestWith(req, func(response *http.Response) error {
		if response.StatusCode == 200 {
			file, err := os.Create(savePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, response.Body)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return apiErr, err
}

// FindRecords allows search records by params
func (c *Client) FindRecords(option *RecordOptions) ([]*Record, *ApiError, error) {

	req, err := c.newRequest(http.MethodPost, "record/find", option)
	if err != nil {
		return nil, nil, err
	}
	resp := &ResponseRecords{}
	apiErr, err := c.doRequest(req, resp)
	return resp.Data, apiErr, err
}
