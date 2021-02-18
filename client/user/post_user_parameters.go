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

// NewPostUserParams creates a new PostUserParams object
// with the default values initialized.
func NewPostUserParams() *PostUserParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserParamsWithTimeout creates a new PostUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUserParamsWithTimeout(timeout time.Duration) *PostUserParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPostUserParamsWithContext creates a new PostUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUserParamsWithContext(ctx context.Context) *PostUserParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPostUserParamsWithHTTPClient creates a new PostUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUserParamsWithHTTPClient(client *http.Client) *PostUserParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*PostUserParams contains all the parameters to send to the API endpoint
for the post user operation typically these are written to a http.Request
*/
type PostUserParams struct {

	/*Authorization
	  The word "Basic" followed by Basic Authorization Token

	*/
	Authorization string
	/*Body*/
	Body PostUserBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post user params
func (o *PostUserParams) WithTimeout(timeout time.Duration) *PostUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user params
func (o *PostUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user params
func (o *PostUserParams) WithContext(ctx context.Context) *PostUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user params
func (o *PostUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user params
func (o *PostUserParams) WithHTTPClient(client *http.Client) *PostUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user params
func (o *PostUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post user params
func (o *PostUserParams) WithAuthorization(authorization string) *PostUserParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post user params
func (o *PostUserParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post user params
func (o *PostUserParams) WithBody(body PostUserBody) *PostUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user params
func (o *PostUserParams) SetBody(body PostUserBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
