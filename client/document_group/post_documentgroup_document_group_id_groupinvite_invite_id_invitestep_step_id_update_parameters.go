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

// NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams creates a new PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams object
// with the default values initialized.
func NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams() *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParamsWithTimeout creates a new PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParamsWithTimeout(timeout time.Duration) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParamsWithContext creates a new PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParamsWithContext(ctx context.Context) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParamsWithHTTPClient creates a new PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParamsWithHTTPClient(client *http.Client) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	var (
		authorizationDefault = string("Bearer {{access_token}}")
	)
	return &PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams contains all the parameters to send to the API endpoint
for the post documentgroup document group ID groupinvite invite ID invitestep step ID update operation typically these are written to a http.Request
*/
type PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams struct {

	/*Authorization
	  Bearer followed by access_token

	*/
	Authorization string
	/*Body*/
	Body PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateBody
	/*DocumentGroupID*/
	DocumentGroupID string
	/*InviteID*/
	InviteID string
	/*StepID*/
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithTimeout(timeout time.Duration) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithContext(ctx context.Context) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithHTTPClient(client *http.Client) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithAuthorization(authorization string) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithBody(body PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateBody) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetBody(body PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateBody) {
	o.Body = body
}

// WithDocumentGroupID adds the documentGroupID to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithDocumentGroupID(documentGroupID string) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetDocumentGroupID(documentGroupID)
	return o
}

// SetDocumentGroupID adds the documentGroupId to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetDocumentGroupID(documentGroupID string) {
	o.DocumentGroupID = documentGroupID
}

// WithInviteID adds the inviteID to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithInviteID(inviteID string) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetInviteID(inviteID)
	return o
}

// SetInviteID adds the inviteId to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetInviteID(inviteID string) {
	o.InviteID = inviteID
}

// WithStepID adds the stepID to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WithStepID(stepID string) *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the post documentgroup document group ID groupinvite invite ID invitestep step ID update params
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDocumentgroupDocumentGroupIDGroupinviteInviteIDInvitestepStepIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param document_group_id
	if err := r.SetPathParam("document_group_id", o.DocumentGroupID); err != nil {
		return err
	}

	// path param invite_id
	if err := r.SetPathParam("invite_id", o.InviteID); err != nil {
		return err
	}

	// path param step_id
	if err := r.SetPathParam("step_id", o.StepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
