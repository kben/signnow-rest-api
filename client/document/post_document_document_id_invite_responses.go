// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostDocumentDocumentIDInviteReader is a Reader for the PostDocumentDocumentIDInvite structure.
type PostDocumentDocumentIDInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDocumentDocumentIDInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDocumentDocumentIDInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDocumentDocumentIDInviteOK creates a PostDocumentDocumentIDInviteOK with default headers values
func NewPostDocumentDocumentIDInviteOK() *PostDocumentDocumentIDInviteOK {
	return &PostDocumentDocumentIDInviteOK{}
}

/*PostDocumentDocumentIDInviteOK handles this case with default header values.

PostDocumentDocumentIDInviteOK post document document Id invite o k
*/
type PostDocumentDocumentIDInviteOK struct {
	Payload interface{}
}

func (o *PostDocumentDocumentIDInviteOK) Error() string {
	return fmt.Sprintf("[POST /document/{document_id}/invite][%d] postDocumentDocumentIdInviteOK  %+v", 200, o.Payload)
}

func (o *PostDocumentDocumentIDInviteOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostDocumentDocumentIDInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostDocumentDocumentIDInviteBody post document document ID invite body
swagger:model PostDocumentDocumentIDInviteBody
*/
type PostDocumentDocumentIDInviteBody struct {

	// cc
	Cc []interface{} `json:"cc"`

	// from
	From string `json:"from,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// to
	To []*PostDocumentDocumentIDInviteParamsBodyToItems0 `json:"to"`
}

// Validate validates this post document document ID invite body
func (o *PostDocumentDocumentIDInviteBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDocumentDocumentIDInviteBody) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(o.To) { // not required
		return nil
	}

	for i := 0; i < len(o.To); i++ {
		if swag.IsZero(o.To[i]) { // not required
			continue
		}

		if o.To[i] != nil {
			if err := o.To[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDocumentDocumentIDInviteBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDocumentDocumentIDInviteBody) UnmarshalBinary(b []byte) error {
	var res PostDocumentDocumentIDInviteBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDocumentDocumentIDInviteParamsBodyToItems0 post document document ID invite params body to items0
swagger:model PostDocumentDocumentIDInviteParamsBodyToItems0
*/
type PostDocumentDocumentIDInviteParamsBodyToItems0 struct {

	// email
	Email string `json:"email,omitempty"`

	// order
	Order int64 `json:"order,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// role id
	RoleID string `json:"role_id,omitempty"`
}

// Validate validates this post document document ID invite params body to items0
func (o *PostDocumentDocumentIDInviteParamsBodyToItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDocumentDocumentIDInviteParamsBodyToItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDocumentDocumentIDInviteParamsBodyToItems0) UnmarshalBinary(b []byte) error {
	var res PostDocumentDocumentIDInviteParamsBodyToItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
