package client

import (
	// Modules
	schema "github.com/thevfxcoop/go-workable-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// GetJobs returns a collection of jobs
func (this *Client) GetJobs(opts ...RequestOpt) ([]*schema.Job, error) {
	var response struct {
		Jobs   []*schema.Job  `json:"jobs"`
		Paging *schema.Paging `json:"paging"`
	}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &response, append(opts, OptPath("jobs"))...); err != nil {
		return nil, err
	} else {
		return response.Jobs, nil
	}
}

// GetJobWithShortcode returns a single job
func (this *Client) GetJobWithShortcode(value string, opts ...RequestOpt) (*schema.Job, error) {
	var response schema.Job
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &response, append(opts, OptPath("jobs", value))...); err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}
