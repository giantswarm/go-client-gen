// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4GetClusterAppsResponseItemsMetadata v4 get cluster apps response items metadata
// swagger:model v4GetClusterAppsResponseItemsMetadata
type V4GetClusterAppsResponseItemsMetadata struct {

	// The labels that are set on this App
	Labels interface{} `json:"labels,omitempty"`

	// The identifier you set when creating this app
	Name string `json:"name,omitempty"`
}

// Validate validates this v4 get cluster apps response items metadata
func (m *V4GetClusterAppsResponseItemsMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4GetClusterAppsResponseItemsMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4GetClusterAppsResponseItemsMetadata) UnmarshalBinary(b []byte) error {
	var res V4GetClusterAppsResponseItemsMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
