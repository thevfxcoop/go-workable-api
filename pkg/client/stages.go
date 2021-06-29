package client

import (
	schema "github.com/thevfxcoop/go-workable-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// GetCandidates returns a collection of candidates
func (this *Client) GetStages(opts ...RequestOpt) ([]*schema.Stage, error) {
	var response struct {
		Stages []*schema.Stage `json:"stages"`
	}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &response, append(opts, OptPath("stages"))...); err != nil {
		return nil, err
	} else {
		return response.Stages, nil
	}
}
