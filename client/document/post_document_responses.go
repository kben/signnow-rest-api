// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostDocumentReader is a Reader for the PostDocument structure.
type PostDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDocumentOK creates a PostDocumentOK with default headers values
func NewPostDocumentOK() *PostDocumentOK {
	return &PostDocumentOK{}
}

/*PostDocumentOK handles this case with default header values.

PostDocumentOK post document o k
*/
type PostDocumentOK struct {
	Payload *PostDocumentOKBody
}

func (o *PostDocumentOK) Error() string {
	return fmt.Sprintf("[POST /document][%d] postDocumentOK  %+v", 200, o.Payload)
}

func (o *PostDocumentOK) GetPayload() *PostDocumentOKBody {
	return o.Payload
}

func (o *PostDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostDocumentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostDocumentOKBody post document o k body
swagger:model PostDocumentOKBody
*/
type PostDocumentOKBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this post document o k body
func (o *PostDocumentOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDocumentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDocumentOKBody) UnmarshalBinary(b []byte) error {
	var res PostDocumentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
