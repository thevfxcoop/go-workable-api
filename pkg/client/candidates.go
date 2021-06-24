package client

import (
	// Modules
	workable "github.com/thevfxcoop/go-workable-api"
	schema "github.com/thevfxcoop/go-workable-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// GetCandidates returns a collection of candidates
func (this *Client) GetCandidates(opts ...RequestOpt) ([]*schema.Candidate, error) {
	var response struct {
		Candidates []*schema.Candidate `json:"candidates"`
		Paging     *schema.Paging      `json:"paging"`
	}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &response, append(opts, OptPath("candidates"))...); err != nil {
		return nil, err
	} else {
		return response.Candidates, nil
	}
}

// GetCandidate will return the full job JSON object of a specific candidate
func (this *Client) GetCandidate(value string, opts ...RequestOpt) (*schema.Candidate, error) {
	var response struct {
		Candidate *schema.Candidate `json:"candidate"`
	}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &response, append(opts, OptPath("candidates", value))...); err != nil {
		return nil, err
	} else {
		return response.Candidate, nil
	}
}

// CreateCandidate will create a candidate at the specified job
func (this *Client) CreateCandidate(shortcode string, candidate *schema.Candidate, opts ...RequestOpt) (*schema.Candidate, error) {
	var response struct {
		Status    string            `json:"status"`
		Candidate *schema.Candidate `json:"candidate"`
	}
	payload := NewCreateCandidatePayload(candidate)
	if err := this.Do(payload, &response, append(opts, OptPath("jobs", shortcode, "candidates"))...); err != nil {
		return nil, err
	} else if response.Status != "created" {
		return nil, workable.ErrUnexpectedResponse.With(response.Status)
	} else {
		return response.Candidate, nil
	}
}
