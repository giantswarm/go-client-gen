// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4InfoResponseFeaturesHaMasters Support for multiple master nodes.
// swagger:model v4InfoResponseFeaturesHaMasters
type V4InfoResponseFeaturesHaMasters struct {

	// The minimum release version number required to have support for multiple master nodes.
	ReleaseVersionMinimum string `json:"release_version_minimum,omitempty"`
}

// Validate validates this v4 info response features ha masters
func (m *V4InfoResponseFeaturesHaMasters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4InfoResponseFeaturesHaMasters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4InfoResponseFeaturesHaMasters) UnmarshalBinary(b []byte) error {
	var res V4InfoResponseFeaturesHaMasters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
