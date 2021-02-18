// Code generated by go-swagger; DO NOT EDIT.

package document_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostDocumentgroupReader is a Reader for the PostDocumentgroup structure.
type PostDocumentgroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDocumentgroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDocumentgroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDocumentgroupOK creates a PostDocumentgroupOK with default headers values
func NewPostDocumentgroupOK() *PostDocumentgroupOK {
	return &PostDocumentgroupOK{}
}

/*PostDocumentgroupOK handles this case with default header values.

PostDocumentgroupOK post documentgroup o k
*/
type PostDocumentgroupOK struct {
	Payload interface{}
}

func (o *PostDocumentgroupOK) Error() string {
	return fmt.Sprintf("[POST /documentgroup][%d] postDocumentgroupOK  %+v", 200, o.Payload)
}

func (o *PostDocumentgroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostDocumentgroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostDocumentgroupBody post documentgroup body
swagger:model PostDocumentgroupBody
*/
type PostDocumentgroupBody struct {

	// document ids
	DocumentIds []string `json:"document_ids"`

	// group name
	GroupName string `json:"group_name,omitempty"`
}

// Validate validates this post documentgroup body
func (o *PostDocumentgroupBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDocumentgroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDocumentgroupBody) UnmarshalBinary(b []byte) error {
	var res PostDocumentgroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}