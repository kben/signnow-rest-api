// Code generated by go-swagger; DO NOT EDIT.

package signing_link

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

// NewPostLinkParams creates a new PostLinkParams object
// with the default values initialized.
func NewPostLinkParams() *PostLinkParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostLinkParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLinkParamsWithTimeout creates a new PostLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLinkParamsWithTimeout(timeout time.Duration) *PostLinkParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostLinkParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPostLinkParamsWithContext creates a new PostLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLinkParamsWithContext(ctx context.Context) *PostLinkParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostLinkParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPostLinkParamsWithHTTPClient creates a new PostLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLinkParamsWithHTTPClient(client *http.Client) *PostLinkParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostLinkParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*PostLinkParams contains all the parameters to send to the API endpoint
for the post link operation typically these are written to a http.Request
*/
type PostLinkParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*Body*/
	Body PostLinkBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post link params
func (o *PostLinkParams) WithTimeout(timeout time.Duration) *PostLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post link params
func (o *PostLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post link params
func (o *PostLinkParams) WithContext(ctx context.Context) *PostLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post link params
func (o *PostLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post link params
func (o *PostLinkParams) WithHTTPClient(client *http.Client) *PostLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post link params
func (o *PostLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post link params
func (o *PostLinkParams) WithAuthorization(authorization string) *PostLinkParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post link params
func (o *PostLinkParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post link params
func (o *PostLinkParams) WithBody(body PostLinkBody) *PostLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post link params
func (o *PostLinkParams) SetBody(body PostLinkBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}