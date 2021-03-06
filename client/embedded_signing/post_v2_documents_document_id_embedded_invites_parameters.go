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

// NewPostV2DocumentsDocumentIDEmbeddedInvitesParams creates a new PostV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized.
func NewPostV2DocumentsDocumentIDEmbeddedInvitesParams() *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostV2DocumentsDocumentIDEmbeddedInvitesParamsWithTimeout creates a new PostV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostV2DocumentsDocumentIDEmbeddedInvitesParamsWithTimeout(timeout time.Duration) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPostV2DocumentsDocumentIDEmbeddedInvitesParamsWithContext creates a new PostV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostV2DocumentsDocumentIDEmbeddedInvitesParamsWithContext(ctx context.Context) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPostV2DocumentsDocumentIDEmbeddedInvitesParamsWithHTTPClient creates a new PostV2DocumentsDocumentIDEmbeddedInvitesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostV2DocumentsDocumentIDEmbeddedInvitesParamsWithHTTPClient(client *http.Client) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostV2DocumentsDocumentIDEmbeddedInvitesParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*PostV2DocumentsDocumentIDEmbeddedInvitesParams contains all the parameters to send to the API endpoint
for the post v2 documents document ID embedded invites operation typically these are written to a http.Request
*/
type PostV2DocumentsDocumentIDEmbeddedInvitesParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*Body*/
	Body PostV2DocumentsDocumentIDEmbeddedInvitesBody
	/*DocumentID*/
	DocumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) WithTimeout(timeout time.Duration) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) WithContext(ctx context.Context) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) WithHTTPClient(client *http.Client) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) WithAuthorization(authorization string) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) WithBody(body PostV2DocumentsDocumentIDEmbeddedInvitesBody) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) SetBody(body PostV2DocumentsDocumentIDEmbeddedInvitesBody) {
	o.Body = body
}

// WithDocumentID adds the documentID to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) WithDocumentID(documentID string) *PostV2DocumentsDocumentIDEmbeddedInvitesParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the post v2 documents document ID embedded invites params
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV2DocumentsDocumentIDEmbeddedInvitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param document_id
	if err := r.SetPathParam("document_id", o.DocumentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
