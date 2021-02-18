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

// PutDocumentDocumentIDReader is a Reader for the PutDocumentDocumentID structure.
type PutDocumentDocumentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDocumentDocumentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDocumentDocumentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDocumentDocumentIDOK creates a PutDocumentDocumentIDOK with default headers values
func NewPutDocumentDocumentIDOK() *PutDocumentDocumentIDOK {
	return &PutDocumentDocumentIDOK{}
}

/*PutDocumentDocumentIDOK handles this case with default header values.

PutDocumentDocumentIDOK put document document Id o k
*/
type PutDocumentDocumentIDOK struct {
	Payload *PutDocumentDocumentIDOKBody
}

func (o *PutDocumentDocumentIDOK) Error() string {
	return fmt.Sprintf("[PUT /document/{document_id}][%d] putDocumentDocumentIdOK  %+v", 200, o.Payload)
}

func (o *PutDocumentDocumentIDOK) GetPayload() *PutDocumentDocumentIDOKBody {
	return o.Payload
}

func (o *PutDocumentDocumentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutDocumentDocumentIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutDocumentDocumentIDOKBody put document document ID o k body
swagger:model PutDocumentDocumentIDOKBody
*/
type PutDocumentDocumentIDOKBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this put document document ID o k body
func (o *PutDocumentDocumentIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutDocumentDocumentIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutDocumentDocumentIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutDocumentDocumentIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}