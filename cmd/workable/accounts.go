package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-workable-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Accounts struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewAccounts(client *client.Client) Command {
	return &Accounts{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Accounts) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "accounts" && len(args) == 1 {
		return this.RunAccounts, params
	}
	return nil, nil
}

func (this *Accounts) RunAccounts(params url.Values) error {
	if accounts, err := this.Client.GetAccounts(); err != nil {
		return err
	} else {
		return this.output(accounts)
	}
}
