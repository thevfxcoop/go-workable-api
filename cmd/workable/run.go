package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"

	"github.com/thevfxcoop/go-workable-api"
	"github.com/thevfxcoop/go-workable-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Command interface {
	Matches([]string) (fn, url.Values)
}

type command struct {
	*client.Client
}

type fn func(url.Values) error

///////////////////////////////////////////////////////////////////////////////
// METHODS

func Run(args []string, client *client.Client) error {
	// Register commands
	commands := []Command{
		NewAccounts(client),
		NewCandidates(client),
		NewJobs(client),
		NewStages(client),
	}

	// If no arguments, then use "accounts"
	if len(args) == 0 {
		args = []string{"accounts"}
	}

	// Cycle through commands and run first which matches
	for _, cmd := range commands {
		if fn, params := cmd.Matches(args); fn != nil && params != nil {
			return fn(params)
		}
	}

	// No command found
	return workable.ErrNotFound.With(args[0])
}

func Usage(flags *flag.FlagSet) {
	fmt.Fprintf(flags.Output(), "Usage of %v:\n", flags.Name())
	fmt.Fprintf(flags.Output(), "  workable <flags> <command> (<args>)\n")
	fmt.Fprintf(flags.Output(), "\nFlags:\n")
	flags.PrintDefaults()
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *command) output(v interface{}) error {
	if data, err := json.MarshalIndent(v, "", "  "); err != nil {
		return err
	} else {
		fmt.Println(string(data))
	}
	return nil
}
