// Code generated by go-swagger; DO NOT EDIT.

package template

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

// PostTemplateReader is a Reader for the PostTemplate structure.
type PostTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTemplateOK creates a PostTemplateOK with default headers values
func NewPostTemplateOK() *PostTemplateOK {
	return &PostTemplateOK{}
}

/*PostTemplateOK handles this case with default header values.

PostTemplateOK post template o k
*/
type PostTemplateOK struct {
	Payload *PostTemplateOKBody
}

func (o *PostTemplateOK) Error() string {
	return fmt.Sprintf("[POST /template][%d] postTemplateOK  %+v", 200, o.Payload)
}

func (o *PostTemplateOK) GetPayload() *PostTemplateOKBody {
	return o.Payload
}

func (o *PostTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTemplateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostTemplateBody post template body
swagger:model PostTemplateBody
*/
type PostTemplateBody struct {

	// document id
	// Required: true
	DocumentID *string `json:"document_id"`

	// document name
	// Required: true
	DocumentName *string `json:"document_name"`
}

// Validate validates this post template body
func (o *PostTemplateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocumentID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDocumentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTemplateBody) validateDocumentID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"document_id", "body", o.DocumentID); err != nil {
		return err
	}

	return nil
}

func (o *PostTemplateBody) validateDocumentName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"document_name", "body", o.DocumentName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTemplateBody) UnmarshalBinary(b []byte) error {
	var res PostTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTemplateOKBody post template o k body
swagger:model PostTemplateOKBody
*/
type PostTemplateOKBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this post template o k body
func (o *PostTemplateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTemplateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTemplateOKBody) UnmarshalBinary(b []byte) error {
	var res PostTemplateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}