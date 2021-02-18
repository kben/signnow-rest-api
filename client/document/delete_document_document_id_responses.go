// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDocumentDocumentIDReader is a Reader for the DeleteDocumentDocumentID structure.
type DeleteDocumentDocumentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDocumentDocumentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDocumentDocumentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDocumentDocumentIDOK creates a DeleteDocumentDocumentIDOK with default headers values
func NewDeleteDocumentDocumentIDOK() *DeleteDocumentDocumentIDOK {
	return &DeleteDocumentDocumentIDOK{}
}

/*DeleteDocumentDocumentIDOK handles this case with default header values.

DeleteDocumentDocumentIDOK delete document document Id o k
*/
type DeleteDocumentDocumentIDOK struct {
	Payload interface{}
}

func (o *DeleteDocumentDocumentIDOK) Error() string {
	return fmt.Sprintf("[DELETE /document/{document_id}][%d] deleteDocumentDocumentIdOK  %+v", 200, o.Payload)
}

func (o *DeleteDocumentDocumentIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDocumentDocumentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}