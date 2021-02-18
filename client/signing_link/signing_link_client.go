// Code generated by go-swagger; DO NOT EDIT.

package signing_link

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new signing link API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for signing link API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PostLink(params *PostLinkParams) (*PostLinkOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PostLink creates signing link

  Creates a Signing Link for a specific document.
*/
func (a *Client) PostLink(params *PostLinkParams) (*PostLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLink",
		Method:             "POST",
		PathPattern:        "/link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
