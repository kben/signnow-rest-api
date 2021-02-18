// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetUserDocumentsv2Params creates a new GetUserDocumentsv2Params object
// with the default values initialized.
func NewGetUserDocumentsv2Params() *GetUserDocumentsv2Params {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetUserDocumentsv2Params{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserDocumentsv2ParamsWithTimeout creates a new GetUserDocumentsv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserDocumentsv2ParamsWithTimeout(timeout time.Duration) *GetUserDocumentsv2Params {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetUserDocumentsv2Params{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewGetUserDocumentsv2ParamsWithContext creates a new GetUserDocumentsv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserDocumentsv2ParamsWithContext(ctx context.Context) *GetUserDocumentsv2Params {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetUserDocumentsv2Params{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewGetUserDocumentsv2ParamsWithHTTPClient creates a new GetUserDocumentsv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserDocumentsv2ParamsWithHTTPClient(client *http.Client) *GetUserDocumentsv2Params {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetUserDocumentsv2Params{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*GetUserDocumentsv2Params contains all the parameters to send to the API endpoint
for the get user documentsv2 operation typically these are written to a http.Request
*/
type GetUserDocumentsv2Params struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) WithTimeout(timeout time.Duration) *GetUserDocumentsv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) WithContext(ctx context.Context) *GetUserDocumentsv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) WithHTTPClient(client *http.Client) *GetUserDocumentsv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) WithAuthorization(authorization string) *GetUserDocumentsv2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get user documentsv2 params
func (o *GetUserDocumentsv2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserDocumentsv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
