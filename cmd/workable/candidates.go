package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-workable-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Candidates struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewCandidates(client *client.Client) Command {
	return &Candidates{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Candidates) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "candidates" && len(args) == 1 {
		return this.RunCandidates, params
	}
	if args[0] == "candidate" && len(args) == 2 {
		params.Set("candidate", args[1])
		return this.RunCandidate, params
	}
	return nil, nil
}

func (this *Candidates) RunCandidates(params url.Values) error {
	if candidates, err := this.Client.GetCandidates(); err != nil {
		return err
	} else {
		return this.output(candidates)
	}
}

func (this *Candidates) RunCandidate(params url.Values) error {
	if candidate, err := this.Client.GetCandidate(params.Get("candidate")); err != nil {
		return err
	} else {
		return this.output(candidate)
	}
}
