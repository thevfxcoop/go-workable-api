package client

import (
	"net/http"

	// Modules
	schema "github.com/thevfxcoop/go-workable-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Payload interface {
	Method() string
	Accept() string
}

type payload struct {
	method string
	accept string
}

type createpayload struct {
	payload
	*schema.Candidate
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewGetPayload(accept string) Payload {
	this := new(payload)
	this.method = http.MethodGet
	this.accept = accept
	return this
}

func NewCreateCandidatePayload(value *schema.Candidate) Payload {
	this := new(createpayload)
	this.method = http.MethodPost
	this.accept = ContentTypeJson
	this.Candidate = value
	return this
}

///////////////////////////////////////////////////////////////////////////////
// PAYLOAD METHODS

func (this *payload) Method() string {
	return this.method
}

func (this *payload) Accept() string {
	return this.accept
}
