package client

import (
	// Modules
	schema "github.com/thevfxcoop/go-workable-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// GetAccounts returns a collection of accounts you have access
func (this *Client) GetAccounts(opts ...RequestOpt) ([]*schema.Account, error) {
	var response struct {
		Accounts []*schema.Account `json:"accounts"`
	}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &response, append(opts, OptPath("accounts"))...); err != nil {
		return nil, err
	} else {
		return response.Accounts, nil
	}
}
