// Code generated by go-swagger; DO NOT EDIT.

package signing_link

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

// PostLinkReader is a Reader for the PostLink structure.
type PostLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLinkOK creates a PostLinkOK with default headers values
func NewPostLinkOK() *PostLinkOK {
	return &PostLinkOK{}
}

/*PostLinkOK handles this case with default header values.

PostLinkOK post link o k
*/
type PostLinkOK struct {
	Payload *PostLinkOKBody
}

func (o *PostLinkOK) Error() string {
	return fmt.Sprintf("[POST /link][%d] postLinkOK  %+v", 200, o.Payload)
}

func (o *PostLinkOK) GetPayload() *PostLinkOKBody {
	return o.Payload
}

func (o *PostLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostLinkOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLinkBody post link body
swagger:model PostLinkBody
*/
type PostLinkBody struct {

	// document id
	// Required: true
	DocumentID *string `json:"document_id"`
}

// Validate validates this post link body
func (o *PostLinkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocumentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLinkBody) validateDocumentID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"document_id", "body", o.DocumentID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLinkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLinkBody) UnmarshalBinary(b []byte) error {
	var res PostLinkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLinkOKBody post link o k body
swagger:model PostLinkOKBody
*/
type PostLinkOKBody struct {

	// url
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// url no signup
	// Format: uri
	URLNoSignup strfmt.URI `json:"url_no_signup,omitempty"`
}

// Validate validates this post link o k body
func (o *PostLinkOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURLNoSignup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLinkOKBody) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(o.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("postLinkOK"+"."+"url", "body", "uri", o.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostLinkOKBody) validateURLNoSignup(formats strfmt.Registry) error {

	if swag.IsZero(o.URLNoSignup) { // not required
		return nil
	}

	if err := validate.FormatOf("postLinkOK"+"."+"url_no_signup", "body", "uri", o.URLNoSignup.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLinkOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLinkOKBody) UnmarshalBinary(b []byte) error {
	var res PostLinkOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
