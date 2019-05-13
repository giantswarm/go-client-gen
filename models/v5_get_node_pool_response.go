// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V5GetNodePoolResponse v5 get node pool response
// swagger:model v5GetNodePoolResponse
type V5GetNodePoolResponse struct {

	// Names of the availability zones used by the nodes of this pool.
	//
	AvailabilityZones []string `json:"availability_zones"`

	// Node pool identifier. Unique within a tenant cluster.
	ID string `json:"id,omitempty"`

	// Node pool name
	Name string `json:"name,omitempty"`

	// node spec
	NodeSpec *V5GetNodePoolResponseNodeSpec `json:"node_spec,omitempty"`

	// scaling
	Scaling *V5GetNodePoolResponseScaling `json:"scaling,omitempty"`
}

// Validate validates this v5 get node pool response
func (m *V5GetNodePoolResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScaling(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5GetNodePoolResponse) validateNodeSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeSpec) { // not required
		return nil
	}

	if m.NodeSpec != nil {
		if err := m.NodeSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_spec")
			}
			return err
		}
	}

	return nil
}

func (m *V5GetNodePoolResponse) validateScaling(formats strfmt.Registry) error {

	if swag.IsZero(m.Scaling) { // not required
		return nil
	}

	if m.Scaling != nil {
		if err := m.Scaling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scaling")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5GetNodePoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5GetNodePoolResponse) UnmarshalBinary(b []byte) error {
	var res V5GetNodePoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}