// Code generated by go-swagger; DO NOT EDIT.

package smart_fields

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
	"github.com/go-openapi/validate"
)

// PostDocumentDocumentIDIntegrationObjectSmartfieldsReader is a Reader for the PostDocumentDocumentIDIntegrationObjectSmartfields structure.
type PostDocumentDocumentIDIntegrationObjectSmartfieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDocumentDocumentIDIntegrationObjectSmartfieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDocumentDocumentIDIntegrationObjectSmartfieldsOK creates a PostDocumentDocumentIDIntegrationObjectSmartfieldsOK with default headers values
func NewPostDocumentDocumentIDIntegrationObjectSmartfieldsOK() *PostDocumentDocumentIDIntegrationObjectSmartfieldsOK {
	return &PostDocumentDocumentIDIntegrationObjectSmartfieldsOK{}
}

/*PostDocumentDocumentIDIntegrationObjectSmartfieldsOK handles this case with default header values.

PostDocumentDocumentIDIntegrationObjectSmartfieldsOK post document document Id integration object smartfields o k
*/
type PostDocumentDocumentIDIntegrationObjectSmartfieldsOK struct {
	Payload interface{}
}

func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsOK) Error() string {
	return fmt.Sprintf("[POST /document/{document_id}/integration/object/smartfields][%d] postDocumentDocumentIdIntegrationObjectSmartfieldsOK  %+v", 200, o.Payload)
}

func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostDocumentDocumentIDIntegrationObjectSmartfieldsBody post document document ID integration object smartfields body
swagger:model PostDocumentDocumentIDIntegrationObjectSmartfieldsBody
*/
type PostDocumentDocumentIDIntegrationObjectSmartfieldsBody struct {

	// client timestamp
	// Required: true
	ClientTimestamp *string `json:"client_timestamp"`

	// data
	// Required: true
	Data []*PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0 `json:"data"`
}

// Validate validates this post document document ID integration object smartfields body
func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsBody) validateClientTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"client_timestamp", "body", o.ClientTimestamp); err != nil {
		return err
	}

	return nil
}

func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsBody) UnmarshalBinary(b []byte) error {
	var res PostDocumentDocumentIDIntegrationObjectSmartfieldsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0 post document document ID integration object smartfields params body data items0
swagger:model PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0
*/
type PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0 struct {

	// field name1
	FieldName1 string `json:"field_name1,omitempty"`

	// field name2
	FieldName2 string `json:"field_name2,omitempty"`
}

// Validate validates this post document document ID integration object smartfields params body data items0
func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res PostDocumentDocumentIDIntegrationObjectSmartfieldsParamsBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
