package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sync"
	"time"

	// Modules
	workable "github.com/thevfxcoop/go-workable-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Client struct {
	sync.Mutex
	*http.Client

	req      *http.Request
	endpoint *url.URL
	rate     float32 // number of requests allowed per second
	ua       string
	apikey   string
	strict   bool
	ts       time.Time
}

type ClientOpt func(*Client) error
type RequestOpt func(*http.Request) error

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	DefaultTimeout   = time.Second * 10
	DefaultUserAgent = "github.com/thevfxcoop/go-workable-api"
	PathSeparator    = string(os.PathSeparator)
	ContentTypeJson  = "application/json"
	ContentTypeText  = "text/plain"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewClient(endpoint *url.URL, opts ...ClientOpt) (*Client, error) {
	this := new(Client)

	// Check incoming parameters
	if endpoint == nil {
		return nil, workable.ErrBadParameter
	} else {
		this.endpoint = endpoint
	}

	// Create a standard request
	if req, err := http.NewRequest(http.MethodPost, this.endpoint.String(), nil); err != nil {
		return nil, err
	} else {
		this.req = req
	}

	// Create a HTTP client
	this.Client = &http.Client{
		Timeout:   DefaultTimeout,
		Transport: http.DefaultTransport,
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(this); err != nil {
			return nil, err
		}
	}

	// Return success
	return this, nil
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *Client) String() string {
	str := "<workable.client"
	str += fmt.Sprintf(" endpoint=%q", redactedUrl(this.endpoint))
	if this.Client.Timeout > 0 {
		str += fmt.Sprint(" timeout=", this.Client.Timeout)
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// NewRequest creates a request which can be used to return responses from the API
func (this *Client) NewRequest(method, accept string, body io.Reader) (*http.Request, error) {
	// Make a request
	r, err := http.NewRequest(method, this.endpoint.String(), body)
	if err != nil {
		return nil, err
	}

	// Set the user-agent and credentials
	r.Header.Set("Content-Type", ContentTypeJson)
	r.Header.Set("Accept", accept)
	r.Header.Set("User-Agent", this.ua)

	// Set the bearer token
	if this.apikey != "" {
		r.Header.Set("Authorization", "Bearer "+this.apikey)
	}

	// Return success
	return r, nil
}

// Do will make a JSON request, populate an object with the response and return any errors
func (this *Client) Do(in Payload, out interface{}, opts ...RequestOpt) error {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()

	// Check rate limit
	now := time.Now()
	if this.ts.IsZero() == false && this.rate > 0.0 {
		next := this.ts.Add(time.Duration(float32(time.Second) / this.rate))
		if next.After(now) {
			time.Sleep(next.Sub(now))
		}
	}

	// Set timestamp at return
	defer func(now time.Time) {
		this.ts = now
	}(now)

	// Make a request
	var data []byte
	var err error
	if in != nil {
		if data, err = json.Marshal(in); err != nil {
			return err
		}
	}
	req, err := this.NewRequest(in.Method(), in.Accept(), bytes.NewReader(data))
	if err != nil {
		return err
	}

	if debug, ok := this.Client.Transport.(*logtransport); ok {
		debug.Payload(in)
	}

	// Parse through options
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return err
		}
	}

	// Do the request
	response, err := this.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Decode body - this can be an array or an object, so we read the whole body
	// and choose happy and sad paths
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Check status code
	if response.StatusCode < 200 || response.StatusCode > 299 {
		// Read any information from the body
		var err string
		if err := decodeString(&err, string(data)); err != nil {
			return err
		}
		return workable.ErrUnexpectedResponse.With(response.Status, ": ", err)
	}

	// When in strict mode, check content type returned is as expected
	if this.strict {
		contenttype := response.Header.Get("Content-Type")
		if mimetype, _, err := mime.ParseMediaType(contenttype); err != nil {
			return workable.ErrUnexpectedResponse.With(contenttype)
		} else if mimetype != in.Accept() {
			return workable.ErrUnexpectedResponse.With(contenttype)
		}
	}

	// Return success if out is nil
	if out == nil {
		return nil
	}

	// If JSON, then decode body
	if in.Accept() == ContentTypeJson {
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(out); err == nil {
			return nil
		} else {
			return err
		}
	} else if in.Accept() == ContentTypeText {
		// Decode as text
		if err := decodeString(out, string(data)); err != nil {
			return err
		}
	}

	// Return success
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Remove any usernames and passwords before printing out
func redactedUrl(url *url.URL) string {
	url_ := *url // make a copy
	url_.User = nil
	return url_.String()
}

// Set string from data
func decodeString(v interface{}, data string) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return workable.ErrInternalAppError.With("DecodeString")
	} else {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.String {
		return workable.ErrInternalAppError.With("DecodeString")
	}
	rv.Set(reflect.ValueOf(data))
	return nil
}
