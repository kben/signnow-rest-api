// Code generated by go-swagger; DO NOT EDIT.

package folder

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

// NewDeleteUserFolderFolderIDParams creates a new DeleteUserFolderFolderIDParams object
// with the default values initialized.
func NewDeleteUserFolderFolderIDParams() *DeleteUserFolderFolderIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteUserFolderFolderIDParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserFolderFolderIDParamsWithTimeout creates a new DeleteUserFolderFolderIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserFolderFolderIDParamsWithTimeout(timeout time.Duration) *DeleteUserFolderFolderIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteUserFolderFolderIDParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewDeleteUserFolderFolderIDParamsWithContext creates a new DeleteUserFolderFolderIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserFolderFolderIDParamsWithContext(ctx context.Context) *DeleteUserFolderFolderIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteUserFolderFolderIDParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewDeleteUserFolderFolderIDParamsWithHTTPClient creates a new DeleteUserFolderFolderIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserFolderFolderIDParamsWithHTTPClient(client *http.Client) *DeleteUserFolderFolderIDParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &DeleteUserFolderFolderIDParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*DeleteUserFolderFolderIDParams contains all the parameters to send to the API endpoint
for the delete user folder folder ID operation typically these are written to a http.Request
*/
type DeleteUserFolderFolderIDParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*FolderID*/
	FolderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) WithTimeout(timeout time.Duration) *DeleteUserFolderFolderIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) WithContext(ctx context.Context) *DeleteUserFolderFolderIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) WithHTTPClient(client *http.Client) *DeleteUserFolderFolderIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) WithAuthorization(authorization string) *DeleteUserFolderFolderIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFolderID adds the folderID to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) WithFolderID(folderID string) *DeleteUserFolderFolderIDParams {
	o.SetFolderID(folderID)
	return o
}

// SetFolderID adds the folderId to the delete user folder folder ID params
func (o *DeleteUserFolderFolderIDParams) SetFolderID(folderID string) {
	o.FolderID = folderID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserFolderFolderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param folder_id
	if err := r.SetPathParam("folder_id", o.FolderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
