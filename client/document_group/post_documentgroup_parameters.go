// Code generated by go-swagger; DO NOT EDIT.

package document_group

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

// NewPostDocumentgroupParams creates a new PostDocumentgroupParams object
// with the default values initialized.
func NewPostDocumentgroupParams() *PostDocumentgroupParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDocumentgroupParamsWithTimeout creates a new PostDocumentgroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDocumentgroupParamsWithTimeout(timeout time.Duration) *PostDocumentgroupParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPostDocumentgroupParamsWithContext creates a new PostDocumentgroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDocumentgroupParamsWithContext(ctx context.Context) *PostDocumentgroupParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPostDocumentgroupParamsWithHTTPClient creates a new PostDocumentgroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDocumentgroupParamsWithHTTPClient(client *http.Client) *PostDocumentgroupParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*PostDocumentgroupParams contains all the parameters to send to the API endpoint
for the post documentgroup operation typically these are written to a http.Request
*/
type PostDocumentgroupParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*Body*/
	Body PostDocumentgroupBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post documentgroup params
func (o *PostDocumentgroupParams) WithTimeout(timeout time.Duration) *PostDocumentgroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post documentgroup params
func (o *PostDocumentgroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post documentgroup params
func (o *PostDocumentgroupParams) WithContext(ctx context.Context) *PostDocumentgroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post documentgroup params
func (o *PostDocumentgroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post documentgroup params
func (o *PostDocumentgroupParams) WithHTTPClient(client *http.Client) *PostDocumentgroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post documentgroup params
func (o *PostDocumentgroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post documentgroup params
func (o *PostDocumentgroupParams) WithAuthorization(authorization string) *PostDocumentgroupParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post documentgroup params
func (o *PostDocumentgroupParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post documentgroup params
func (o *PostDocumentgroupParams) WithBody(body PostDocumentgroupBody) *PostDocumentgroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post documentgroup params
func (o *PostDocumentgroupParams) SetBody(body PostDocumentgroupBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDocumentgroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
