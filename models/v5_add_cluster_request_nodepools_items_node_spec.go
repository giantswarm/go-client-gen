// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V5AddClusterRequestNodepoolsItemsNodeSpec v5 add cluster request nodepools items node spec
// swagger:model v5AddClusterRequestNodepoolsItemsNodeSpec
type V5AddClusterRequestNodepoolsItemsNodeSpec struct {

	// aws
	Aws *V5AddClusterRequestNodepoolsItemsNodeSpecAws `json:"aws,omitempty"`
}

// Validate validates this v5 add cluster request nodepools items node spec
func (m *V5AddClusterRequestNodepoolsItemsNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5AddClusterRequestNodepoolsItemsNodeSpec) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5AddClusterRequestNodepoolsItemsNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5AddClusterRequestNodepoolsItemsNodeSpec) UnmarshalBinary(b []byte) error {
	var res V5AddClusterRequestNodepoolsItemsNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}