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

// NewGetDocumentgroupDocumentGroupIDParams creates a new GetDocumentgroupDocumentGroupIDParams object
// with the default values initialized.
func NewGetDocumentgroupDocumentGroupIDParams() *GetDocumentgroupDocumentGroupIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetDocumentgroupDocumentGroupIDParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentgroupDocumentGroupIDParamsWithTimeout creates a new GetDocumentgroupDocumentGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDocumentgroupDocumentGroupIDParamsWithTimeout(timeout time.Duration) *GetDocumentgroupDocumentGroupIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetDocumentgroupDocumentGroupIDParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewGetDocumentgroupDocumentGroupIDParamsWithContext creates a new GetDocumentgroupDocumentGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDocumentgroupDocumentGroupIDParamsWithContext(ctx context.Context) *GetDocumentgroupDocumentGroupIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetDocumentgroupDocumentGroupIDParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewGetDocumentgroupDocumentGroupIDParamsWithHTTPClient creates a new GetDocumentgroupDocumentGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDocumentgroupDocumentGroupIDParamsWithHTTPClient(client *http.Client) *GetDocumentgroupDocumentGroupIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &GetDocumentgroupDocumentGroupIDParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*GetDocumentgroupDocumentGroupIDParams contains all the parameters to send to the API endpoint
for the get documentgroup document group ID operation typically these are written to a http.Request
*/
type GetDocumentgroupDocumentGroupIDParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*DocumentGroupID*/
	DocumentGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) WithTimeout(timeout time.Duration) *GetDocumentgroupDocumentGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) WithContext(ctx context.Context) *GetDocumentgroupDocumentGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) WithHTTPClient(client *http.Client) *GetDocumentgroupDocumentGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) WithAuthorization(authorization string) *GetDocumentgroupDocumentGroupIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDocumentGroupID adds the documentGroupID to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) WithDocumentGroupID(documentGroupID string) *GetDocumentgroupDocumentGroupIDParams {
	o.SetDocumentGroupID(documentGroupID)
	return o
}

// SetDocumentGroupID adds the documentGroupId to the get documentgroup document group ID params
func (o *GetDocumentgroupDocumentGroupIDParams) SetDocumentGroupID(documentGroupID string) {
	o.DocumentGroupID = documentGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentgroupDocumentGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param document_group_id
	if err := r.SetPathParam("document_group_id", o.DocumentGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}