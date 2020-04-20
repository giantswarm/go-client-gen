// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V5AddNodePoolRequestNodeSpecAwsInstanceDistribution Attributes defining the instance distribution in the node pool being created.
//
// swagger:model v5AddNodePoolRequestNodeSpecAwsInstanceDistribution
type V5AddNodePoolRequestNodeSpecAwsInstanceDistribution struct {

	// Base capacity of on-demand EC2 instances to use for worker nodes in this pools.
	// When this is larger than 0, this value defines a number of worker nodes that
	// will be created using on-demand EC2 instances, regardless of the value
	// configured as `on_demand_percentage_above_base_capacity`.
	//
	// Minimum: 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty"`

	// Percentage of on-demand EC2 instances to use for worker nodes, instead of spot
	// instances, for instances exceeding `on_demand_base_capacity`. For example, to
	// have half of the worker nodes use spot instances and half use on-demand, set this
	// value to 50.
	//
	// Maximum: 100
	// Minimum: 0
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty"`
}

// Validate validates this v5 add node pool request node spec aws instance distribution
func (m *V5AddNodePoolRequestNodeSpecAwsInstanceDistribution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOnDemandBaseCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnDemandPercentageAboveBaseCapacity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5AddNodePoolRequestNodeSpecAwsInstanceDistribution) validateOnDemandBaseCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.OnDemandBaseCapacity) { // not required
		return nil
	}

	if err := validate.MinimumInt("on_demand_base_capacity", "body", int64(*m.OnDemandBaseCapacity), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *V5AddNodePoolRequestNodeSpecAwsInstanceDistribution) validateOnDemandPercentageAboveBaseCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.OnDemandPercentageAboveBaseCapacity) { // not required
		return nil
	}

	if err := validate.MinimumInt("on_demand_percentage_above_base_capacity", "body", int64(*m.OnDemandPercentageAboveBaseCapacity), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("on_demand_percentage_above_base_capacity", "body", int64(*m.OnDemandPercentageAboveBaseCapacity), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpecAwsInstanceDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpecAwsInstanceDistribution) UnmarshalBinary(b []byte) error {
	var res V5AddNodePoolRequestNodeSpecAwsInstanceDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
