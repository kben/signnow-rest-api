// Code generated by go-swagger; DO NOT EDIT.

package template

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

// NewPostTemplateParams creates a new PostTemplateParams object
// with the default values initialized.
func NewPostTemplateParams() *PostTemplateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostTemplateParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTemplateParamsWithTimeout creates a new PostTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTemplateParamsWithTimeout(timeout time.Duration) *PostTemplateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostTemplateParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPostTemplateParamsWithContext creates a new PostTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTemplateParamsWithContext(ctx context.Context) *PostTemplateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostTemplateParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPostTemplateParamsWithHTTPClient creates a new PostTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTemplateParamsWithHTTPClient(client *http.Client) *PostTemplateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostTemplateParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*PostTemplateParams contains all the parameters to send to the API endpoint
for the post template operation typically these are written to a http.Request
*/
type PostTemplateParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*Body*/
	Body PostTemplateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post template params
func (o *PostTemplateParams) WithTimeout(timeout time.Duration) *PostTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post template params
func (o *PostTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post template params
func (o *PostTemplateParams) WithContext(ctx context.Context) *PostTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post template params
func (o *PostTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post template params
func (o *PostTemplateParams) WithHTTPClient(client *http.Client) *PostTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post template params
func (o *PostTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post template params
func (o *PostTemplateParams) WithAuthorization(authorization string) *PostTemplateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post template params
func (o *PostTemplateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post template params
func (o *PostTemplateParams) WithBody(body PostTemplateBody) *PostTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post template params
func (o *PostTemplateParams) SetBody(body PostTemplateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
