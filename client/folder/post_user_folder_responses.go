// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostUserFolderReader is a Reader for the PostUserFolder structure.
type PostUserFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserFolderOK creates a PostUserFolderOK with default headers values
func NewPostUserFolderOK() *PostUserFolderOK {
	return &PostUserFolderOK{}
}

/*PostUserFolderOK handles this case with default header values.

PostUserFolderOK post user folder o k
*/
type PostUserFolderOK struct {
	Payload interface{}
}

func (o *PostUserFolderOK) Error() string {
	return fmt.Sprintf("[POST /user/folder][%d] postUserFolderOK  %+v", 200, o.Payload)
}

func (o *PostUserFolderOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostUserFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostUserFolderBody post user folder body
swagger:model PostUserFolderBody
*/
type PostUserFolderBody struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// parent id
	// Required: true
	ParentID *string `json:"parent_id"`
}

// Validate validates this post user folder body
func (o *PostUserFolderBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostUserFolderBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PostUserFolderBody) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"parent_id", "body", o.ParentID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostUserFolderBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUserFolderBody) UnmarshalBinary(b []byte) error {
	var res PostUserFolderBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
