// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V5GetNodePoolsResponseItemsNodeSpec v5 get node pools response items node spec
// swagger:model v5GetNodePoolsResponseItemsNodeSpec
type V5GetNodePoolsResponseItemsNodeSpec struct {

	// aws
	Aws *V5GetNodePoolsResponseItemsNodeSpecAws `json:"aws,omitempty"`

	// azure
	Azure *V5GetNodePoolsResponseItemsNodeSpecAzure `json:"azure,omitempty"`

	// volume sizes gb
	VolumeSizesGb *V5GetNodePoolsResponseItemsNodeSpecVolumeSizesGb `json:"volume_sizes_gb,omitempty"`
}

// Validate validates this v5 get node pools response items node spec
func (m *V5GetNodePoolsResponseItemsNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeSizesGb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5GetNodePoolsResponseItemsNodeSpec) validateAws(formats strfmt.Registry) error {

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

func (m *V5GetNodePoolsResponseItemsNodeSpec) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *V5GetNodePoolsResponseItemsNodeSpec) validateVolumeSizesGb(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeSizesGb) { // not required
		return nil
	}

	if m.VolumeSizesGb != nil {
		if err := m.VolumeSizesGb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume_sizes_gb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5GetNodePoolsResponseItemsNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5GetNodePoolsResponseItemsNodeSpec) UnmarshalBinary(b []byte) error {
	var res V5GetNodePoolsResponseItemsNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
