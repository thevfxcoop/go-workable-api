package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	// Modules
	"github.com/thevfxcoop/go-workable-api"
	"github.com/thevfxcoop/go-workable-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

// Command-line flags
var (
	flagEndpoint *string
	flagKey      *string
	flagDebug    *bool
)

///////////////////////////////////////////////////////////////////////////////
// MAIN

func main() {
	// Create flagset
	flags := flag.NewFlagSet("workable", flag.ContinueOnError)
	defineFlags(flags)

	// Parse flags, if no command then ping for deadline version
	if err := flags.Parse(os.Args[1:]); err == flag.ErrHelp {
		os.Exit(0)
	} else if err != nil {
		fmt.Fprintln(flags.Output(), err)
		os.Exit(-1)
	}

	// Set client options
	opts := []client.ClientOpt{}
	if *flagDebug {
		opts = append(opts, client.OptTrace(os.Stderr, true))
	}

	// Create client, ping then run command
	if endpoint, err := getEndpoint(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if key, err := getKey(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if client, err := client.NewClient(endpoint, append(opts, client.OptAPIKey(key))...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if err := Run(flags.Args(), client); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func getEndpoint() (*url.URL, error) {
	if endpoint := os.Getenv("WORKABLE_ENDPOINT"); endpoint != "" {
		return url.Parse(endpoint)
	} else if *flagEndpoint != "" {
		return url.Parse(*flagEndpoint)
	} else {
		return nil, workable.ErrBadParameter.With("-endpoint")
	}
}

func getKey() (string, error) {
	if key := os.Getenv("WORKABLE_APIKEY"); key != "" {
		return key, nil
	} else if *flagKey != "" {
		return *flagKey, nil
	} else {
		return "", workable.ErrBadParameter.With("-key")
	}
}

func defineFlags(flags *flag.FlagSet) {
	flags.Usage = func() { Usage(flags) }
	flagEndpoint = flags.String("endpoint", "", "Endpoint URL, can be overridden with WORKABLE_ENDPOINT environment variable")
	flagKey = flags.String("key", "", "API Key, can be overridden with WORKABLE_APIKEY environment variable")
	flagDebug = flags.Bool("debug", false, "Trace request and reponse with API")
}
