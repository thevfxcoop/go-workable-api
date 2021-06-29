package schema

import (
	"encoding/json"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Stage struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	Position uint   `json:"position"`
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *Stage) String() string {
	if data, err := json.MarshalIndent(this, "", "  "); err != nil {
		panic(err)
	} else {
		return "<workable.stage " + string(data) + ">"
	}
}
