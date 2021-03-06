// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V5AddNodePoolRequestNodeSpecAws Attributes specific to the AWS provider
//
// swagger:model v5AddNodePoolRequestNodeSpecAws
type V5AddNodePoolRequestNodeSpecAws struct {

	// instance distribution
	InstanceDistribution *V5AddNodePoolRequestNodeSpecAwsInstanceDistribution `json:"instance_distribution,omitempty"`

	// EC2 instance type to use for all nodes in the node pool. _(Validated against available instance types.)_
	//
	InstanceType string `json:"instance_type,omitempty"`

	// If true, instance types alike the type set via `instance_type` will be used. This can
	// increase the likelihood to get instances for this pool, especially spot instances at
	// a low rate. If false, only the exact type set as `instance_type` is used.
	// Added with AWS release v11.2.0.
	//
	UseAlikeInstanceTypes *bool `json:"use_alike_instance_types,omitempty"`
}

// Validate validates this v5 add node pool request node spec aws
func (m *V5AddNodePoolRequestNodeSpecAws) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceDistribution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5AddNodePoolRequestNodeSpecAws) validateInstanceDistribution(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceDistribution) { // not required
		return nil
	}

	if m.InstanceDistribution != nil {
		if err := m.InstanceDistribution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance_distribution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpecAws) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpecAws) UnmarshalBinary(b []byte) error {
	var res V5AddNodePoolRequestNodeSpecAws
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
