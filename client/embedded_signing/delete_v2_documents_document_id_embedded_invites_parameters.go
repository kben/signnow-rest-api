// Code generated by go-swagger; DO NOT EDIT.

package embedded_signing

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

// NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParams creates a new DeleteV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized.
func NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParams() *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParamsWithTimeout creates a new DeleteV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParamsWithTimeout(timeout time.Duration) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParamsWithContext creates a new DeleteV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParamsWithContext(ctx context.Context) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParamsWithHTTPClient creates a new DeleteV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV2DocumentsDocumentIDEmbeddedInvitesParamsWithHTTPClient(client *http.Client) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*DeleteV2DocumentsDocumentIDEmbeddedInvitesParams contains all the parameters to send to the API endpoint
for the delete v2 documents document ID embedded invites operation typically these are written to a http.Request
*/
type DeleteV2DocumentsDocumentIDEmbeddedInvitesParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*DocumentID*/
	DocumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) WithTimeout(timeout time.Duration) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) WithContext(ctx context.Context) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) WithHTTPClient(client *http.Client) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) WithAuthorization(authorization string) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDocumentID adds the documentID to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) WithDocumentID(documentID string) *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the delete v2 documents document ID embedded invites params
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV2DocumentsDocumentIDEmbeddedInvitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param document_id
	if err := r.SetPathParam("document_id", o.DocumentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
