package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-workable-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Stages struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewStages(client *client.Client) Command {
	return &Stages{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Stages) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "stages" && len(args) == 1 {
		return this.RunStages, params
	}
	return nil, nil
}

func (this *Stages) RunStages(params url.Values) error {
	if stages, err := this.Client.GetStages(); err != nil {
		return err
	} else {
		return this.output(stages)
	}
}
