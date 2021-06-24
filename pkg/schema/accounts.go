package schema

import (
	"encoding/json"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Subdomain   string `json:"subdomain,omitempty"`
	Description string `json:"description,omitempty"`
	Summary     string `json:"summary,omitempty"`
	WebsiteUrl  string `json:"website_url,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *Account) String() string {
	if data, err := json.MarshalIndent(this, "", "  "); err != nil {
		panic(err)
	} else {
		return "<workable.account " + string(data) + ">"
	}
}
