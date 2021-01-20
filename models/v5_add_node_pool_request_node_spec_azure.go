// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V5AddNodePoolRequestNodeSpecAzure Attributes specific to the Azure provider
//
// swagger:model v5AddNodePoolRequestNodeSpecAzure
type V5AddNodePoolRequestNodeSpecAzure struct {

	// spot instances
	SpotInstances *V5AddNodePoolRequestNodeSpecAzureSpotInstances `json:"spot_instances,omitempty"`

	// Azure VM size to use for all nodes in the node pool. _(Validated against available VM sizes.)_
	//
	VMSize string `json:"vm_size,omitempty"`
}

// Validate validates this v5 add node pool request node spec azure
func (m *V5AddNodePoolRequestNodeSpecAzure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpotInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5AddNodePoolRequestNodeSpecAzure) validateSpotInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.SpotInstances) { // not required
		return nil
	}

	if m.SpotInstances != nil {
		if err := m.SpotInstances.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spot_instances")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpecAzure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpecAzure) UnmarshalBinary(b []byte) error {
	var res V5AddNodePoolRequestNodeSpecAzure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
