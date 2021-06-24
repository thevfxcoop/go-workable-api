package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-workable-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Jobs struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewJobs(client *client.Client) Command {
	return &Jobs{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Jobs) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "jobs" && len(args) == 1 {
		return this.RunJobs, params
	}
	if args[0] == "job" && len(args) == 2 {
		params.Set("job", args[1])
		return this.RunJob, params
	}
	return nil, nil
}

func (this *Jobs) RunJobs(params url.Values) error {
	if jobs, err := this.Client.GetJobs(); err != nil {
		return err
	} else {
		return this.output(jobs)
	}
}

func (this *Jobs) RunJob(params url.Values) error {
	if job, err := this.Client.GetJobWithShortcode(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(job)
	}
}
