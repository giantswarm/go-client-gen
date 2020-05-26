// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4UserListItem v4 user list item
// swagger:model v4UserListItem
type V4UserListItem struct {

	// The date and time that this account was created
	Created string `json:"created,omitempty"`

	// Email address of the user
	Email string `json:"email,omitempty"`

	// The date and time when this account will expire
	Expiry string `json:"expiry,omitempty"`
}

// Validate validates this v4 user list item
func (m *V4UserListItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4UserListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4UserListItem) UnmarshalBinary(b []byte) error {
	var res V4UserListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
