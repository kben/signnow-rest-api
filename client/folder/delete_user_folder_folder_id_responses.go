// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUserFolderFolderIDReader is a Reader for the DeleteUserFolderFolderID structure.
type DeleteUserFolderFolderIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserFolderFolderIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserFolderFolderIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserFolderFolderIDOK creates a DeleteUserFolderFolderIDOK with default headers values
func NewDeleteUserFolderFolderIDOK() *DeleteUserFolderFolderIDOK {
	return &DeleteUserFolderFolderIDOK{}
}

/*DeleteUserFolderFolderIDOK handles this case with default header values.

DeleteUserFolderFolderIDOK delete user folder folder Id o k
*/
type DeleteUserFolderFolderIDOK struct {
	Payload interface{}
}

func (o *DeleteUserFolderFolderIDOK) Error() string {
	return fmt.Sprintf("[DELETE /user/folder/{folder_id}][%d] deleteUserFolderFolderIdOK  %+v", 200, o.Payload)
}

func (o *DeleteUserFolderFolderIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteUserFolderFolderIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
