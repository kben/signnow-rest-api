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

// NewPostUserVerifyemailParams creates a new PostUserVerifyemailParams object
// with the default values initialized.
func NewPostUserVerifyemailParams() *PostUserVerifyemailParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserVerifyemailParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserVerifyemailParamsWithTimeout creates a new PostUserVerifyemailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUserVerifyemailParamsWithTimeout(timeout time.Duration) *PostUserVerifyemailParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserVerifyemailParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPostUserVerifyemailParamsWithContext creates a new PostUserVerifyemailParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUserVerifyemailParamsWithContext(ctx context.Context) *PostUserVerifyemailParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserVerifyemailParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPostUserVerifyemailParamsWithHTTPClient creates a new PostUserVerifyemailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUserVerifyemailParamsWithHTTPClient(client *http.Client) *PostUserVerifyemailParams {
	var (
		authorizationDefault = string("Basic {{basic_authorization_token}}")
	)
	return &PostUserVerifyemailParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*PostUserVerifyemailParams contains all the parameters to send to the API endpoint
for the post user verifyemail operation typically these are written to a http.Request
*/
type PostUserVerifyemailParams struct {

	/*Authorization
	  The word "Basic" followed by Basic Authorization Token

	*/
	Authorization string
	/*Body*/
	Body PostUserVerifyemailBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post user verifyemail params
func (o *PostUserVerifyemailParams) WithTimeout(timeout time.Duration) *PostUserVerifyemailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user verifyemail params
func (o *PostUserVerifyemailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user verifyemail params
func (o *PostUserVerifyemailParams) WithContext(ctx context.Context) *PostUserVerifyemailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user verifyemail params
func (o *PostUserVerifyemailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user verifyemail params
func (o *PostUserVerifyemailParams) WithHTTPClient(client *http.Client) *PostUserVerifyemailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user verifyemail params
func (o *PostUserVerifyemailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post user verifyemail params
func (o *PostUserVerifyemailParams) WithAuthorization(authorization string) *PostUserVerifyemailParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post user verifyemail params
func (o *PostUserVerifyemailParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post user verifyemail params
func (o *PostUserVerifyemailParams) WithBody(body PostUserVerifyemailBody) *PostUserVerifyemailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user verifyemail params
func (o *PostUserVerifyemailParams) SetBody(body PostUserVerifyemailBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserVerifyemailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
