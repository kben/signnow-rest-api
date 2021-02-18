// Code generated by go-swagger; DO NOT EDIT.

package document_group

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
)

// GetUserDocumentgroupsReader is a Reader for the GetUserDocumentgroups structure.
type GetUserDocumentgroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserDocumentgroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserDocumentgroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserDocumentgroupsOK creates a GetUserDocumentgroupsOK with default headers values
func NewGetUserDocumentgroupsOK() *GetUserDocumentgroupsOK {
	return &GetUserDocumentgroupsOK{}
}

/*GetUserDocumentgroupsOK handles this case with default header values.

GetUserDocumentgroupsOK get user documentgroups o k
*/
type GetUserDocumentgroupsOK struct {
	Payload *GetUserDocumentgroupsOKBody
}

func (o *GetUserDocumentgroupsOK) Error() string {
	return fmt.Sprintf("[GET /user/documentgroups][%d] getUserDocumentgroupsOK  %+v", 200, o.Payload)
}

func (o *GetUserDocumentgroupsOK) GetPayload() *GetUserDocumentgroupsOKBody {
	return o.Payload
}

func (o *GetUserDocumentgroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUserDocumentgroupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUserDocumentgroupsOKBody get user documentgroups o k body
swagger:model GetUserDocumentgroupsOKBody
*/
type GetUserDocumentgroupsOKBody struct {

	// document group total count
	DocumentGroupTotalCount int64 `json:"document_group_total_count,omitempty"`

	// document groups
	DocumentGroups []*GetUserDocumentgroupsOKBodyDocumentGroupsItems0 `json:"document_groups"`

	// originator organization settings
	OriginatorOrganizationSettings []interface{} `json:"originator_organization_settings"`
}

// Validate validates this get user documentgroups o k body
func (o *GetUserDocumentgroupsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocumentGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserDocumentgroupsOKBody) validateDocumentGroups(formats strfmt.Registry) error {

	if swag.IsZero(o.DocumentGroups) { // not required
		return nil
	}

	for i := 0; i < len(o.DocumentGroups); i++ {
		if swag.IsZero(o.DocumentGroups[i]) { // not required
			continue
		}

		if o.DocumentGroups[i] != nil {
			if err := o.DocumentGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUserDocumentgroupsOK" + "." + "document_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserDocumentgroupsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUserDocumentgroupsOKBodyDocumentGroupsItems0 get user documentgroups o k body document groups items0
swagger:model GetUserDocumentgroupsOKBodyDocumentGroupsItems0
*/
type GetUserDocumentgroupsOKBodyDocumentGroupsItems0 struct {

	// documents
	Documents []*GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0 `json:"documents"`

	// group id
	GroupID string `json:"group_id,omitempty"`

	// group name
	GroupName string `json:"group_name,omitempty"`

	// invite id
	InviteID string `json:"invite_id,omitempty"`

	// invite status
	InviteStatus string `json:"invite_status,omitempty"`

	// last updated
	LastUpdated string `json:"last_updated,omitempty"`
}

// Validate validates this get user documentgroups o k body document groups items0
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0) validateDocuments(formats strfmt.Registry) error {

	if swag.IsZero(o.Documents) { // not required
		return nil
	}

	for i := 0; i < len(o.Documents); i++ {
		if swag.IsZero(o.Documents[i]) { // not required
			continue
		}

		if o.Documents[i] != nil {
			if err := o.Documents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0) UnmarshalBinary(b []byte) error {
	var res GetUserDocumentgroupsOKBodyDocumentGroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0 get user documentgroups o k body document groups items0 documents items0
swagger:model GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0
*/
type GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0 struct {

	// has credit card number
	HasCreditCardNumber bool `json:"has_credit_card_number,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// roles
	Roles []string `json:"roles"`

	// thumbnail
	Thumbnail *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail `json:"thumbnail,omitempty"`
}

// Validate validates this get user documentgroups o k body document groups items0 documents items0
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateThumbnail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0) validateThumbnail(formats strfmt.Registry) error {

	if swag.IsZero(o.Thumbnail) { // not required
		return nil
	}

	if o.Thumbnail != nil {
		if err := o.Thumbnail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumbnail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0) UnmarshalBinary(b []byte) error {
	var res GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail get user documentgroups o k body document groups items0 documents items0 thumbnail
swagger:model GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail
*/
type GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail struct {

	// large
	Large string `json:"large,omitempty"`

	// medium
	Medium string `json:"medium,omitempty"`

	// small
	Small string `json:"small,omitempty"`
}

// Validate validates this get user documentgroups o k body document groups items0 documents items0 thumbnail
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail) UnmarshalBinary(b []byte) error {
	var res GetUserDocumentgroupsOKBodyDocumentGroupsItems0DocumentsItems0Thumbnail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
