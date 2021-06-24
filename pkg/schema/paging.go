package schema

import (
	"encoding/json"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Paging struct {
	Next string `json:"next"`
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *Paging) String() string {
	if data, err := json.MarshalIndent(this, "", "  "); err != nil {
		panic(err)
	} else {
		return "<workable.paging " + string(data) + ">"
	}
}
