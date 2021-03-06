// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4OrganizationListItem v4 organization list item
// swagger:model v4OrganizationListItem
type V4OrganizationListItem struct {

	// Unique name/identifier of the organization
	ID string `json:"id,omitempty"`
}

// Validate validates this v4 organization list item
func (m *V4OrganizationListItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4OrganizationListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4OrganizationListItem) UnmarshalBinary(b []byte) error {
	var res V4OrganizationListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
