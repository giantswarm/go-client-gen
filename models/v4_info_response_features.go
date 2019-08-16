// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V4InfoResponseFeatures Information on particular capabilities of the installation.
// swagger:model v4InfoResponseFeatures
type V4InfoResponseFeatures struct {

	// nodepools
	Nodepools *V4InfoResponseFeaturesNodepools `json:"nodepools,omitempty"`
}

// Validate validates this v4 info response features
func (m *V4InfoResponseFeatures) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodepools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4InfoResponseFeatures) validateNodepools(formats strfmt.Registry) error {

	if swag.IsZero(m.Nodepools) { // not required
		return nil
	}

	if m.Nodepools != nil {
		if err := m.Nodepools.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodepools")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4InfoResponseFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4InfoResponseFeatures) UnmarshalBinary(b []byte) error {
	var res V4InfoResponseFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
