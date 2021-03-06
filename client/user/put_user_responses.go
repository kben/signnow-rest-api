// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutUserReader is a Reader for the PutUser structure.
type PutUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserOK creates a PutUserOK with default headers values
func NewPutUserOK() *PutUserOK {
	return &PutUserOK{}
}

/*PutUserOK handles this case with default header values.

PutUserOK put user o k
*/
type PutUserOK struct {
	Payload *PutUserOKBody
}

func (o *PutUserOK) Error() string {
	return fmt.Sprintf("[PUT /user][%d] putUserOK  %+v", 200, o.Payload)
}

func (o *PutUserOK) GetPayload() *PutUserOKBody {
	return o.Payload
}

func (o *PutUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutUserBody put user body
swagger:model PutUserBody
*/
type PutUserBody struct {

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// if "true" - all user tokens with "scope =*" expire
	//
	// if "false" - all user tokens except current one are expired
	//
	// Note: parameter must have string format.
	//
	// Default value - "true"
	LogoutAll string `json:"logout_all,omitempty"`

	// old password
	OldPassword string `json:"old_password,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this put user body
func (o *PutUserBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutUserBody) UnmarshalBinary(b []byte) error {
	var res PutUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutUserOKBody put user o k body
swagger:model PutUserOKBody
*/
type PutUserOKBody struct {

	// first name
	FirstName string `json:"first_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`
}

// Validate validates this put user o k body
func (o *PutUserOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutUserOKBody) UnmarshalBinary(b []byte) error {
	var res PutUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
