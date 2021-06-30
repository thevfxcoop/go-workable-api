package client

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	// Modules
	workable "github.com/thevfxcoop/go-workable-api"
)

// OptTimeout sets the timeout on any request. By default, a timeout
// of 10 seconds is used if OptTimeout is not set
func OptTimeout(value time.Duration) ClientOpt {
	return func(client *Client) error {
		client.Client.Timeout = value
		return nil
	}
}

// OptUserAgent sets the user agent string on each API request
// It is set to the default if empty string is passed
func OptUserAgent(value string) ClientOpt {
	return func(client *Client) error {
		value = strings.TrimSpace(value)
		if value == "" {
			client.ua = DefaultUserAgent
		} else {
			client.ua = value
		}
		return nil
	}
}

// OptTrace allows you to be the "man in the middle" on any
// requests so you can see traffic move back and forth.
// Setting verbose to true also displays the JSON response
func OptTrace(w io.Writer, verbose bool) ClientOpt {
	return func(client *Client) error {
		client.Client.Transport = NewLogTransport(w, client.Client.Transport, verbose)
		return nil
	}
}

// OptStrict turns on strict content type checking on anything returned
// from the API
func OptStrict() ClientOpt {
	return func(client *Client) error {
		client.strict = true
		return nil
	}
}

// OptAPIKey adds an authorization header
func OptAPIKey(value string) ClientOpt {
	return func(client *Client) error {
		client.apikey = value
		return nil
	}
}

// OptPath appends path elements onto a request
func OptPath(value ...string) RequestOpt {
	return func(r *http.Request) error {
		// Make a copy
		url := *r.URL
		// Clean up and append path
		url.Path = PathSeparator + filepath.Join(strings.Trim(url.Path, PathSeparator), strings.Join(value, PathSeparator))
		// Set new path
		r.URL = &url
		return nil
	}
}

// OptShortcode appends shortcode query onto the request
func OptShortcode(value string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("shortcode", value)
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// OptStage appends stage query onto the request
func OptStage(value string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("stage", value)
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// OptLimit appends limit query onto the request
func OptLimit(value uint) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("limit", fmt.Sprint(value))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// OptCreatedAfter appends created_after query onto the request
func OptCreatedAfter(value time.Time) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("created_after", value.Format(time.RFC3339))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// OptUpdatedAfter appends created_after query onto the request
func OptUpdatedAfter(value time.Time) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("updated_after", value.Format(time.RFC3339))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// OptState appends state query onto the request
func OptState(value string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("state", value)
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// OptRateLimit sets the limit on number of requests per second
// and the API will sleep when exceeded. For account tokens this is 1 per second
func OptRateLimit(value float32) ClientOpt {
	return func(client *Client) error {
		if value < 0.0 {
			return workable.ErrBadParameter.With("OptRateLimit")
		} else {
			client.rate = value
			return nil
		}
	}
}
