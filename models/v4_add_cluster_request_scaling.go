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

// V4AddClusterRequestScaling Attributes specific to cluster node scaling on KVM and Azure. To have
// full control of the cluster size, min and max can be set to the same
// value. If only `min` or only `max` is specified, `min` and `max` will
// be set equally. Not available on AWS.
//
// swagger:model v4AddClusterRequestScaling
type V4AddClusterRequestScaling struct {

	// Maximum number of cluster nodes
	//
	// Maximum: 999
	// Minimum: 1
	Max int64 `json:"max,omitempty"`

	// Minimum number of cluster nodes
	//
	// Maximum: 999
	// Minimum: 1
	Min int64 `json:"min,omitempty"`
}

// Validate validates this v4 add cluster request scaling
func (m *V4AddClusterRequestScaling) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4AddClusterRequestScaling) validateMax(formats strfmt.Registry) error {

	if swag.IsZero(m.Max) { // not required
		return nil
	}

	if err := validate.MinimumInt("max", "body", int64(m.Max), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max", "body", int64(m.Max), 999, false); err != nil {
		return err
	}

	return nil
}

func (m *V4AddClusterRequestScaling) validateMin(formats strfmt.Registry) error {

	if swag.IsZero(m.Min) { // not required
		return nil
	}

	if err := validate.MinimumInt("min", "body", int64(m.Min), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("min", "body", int64(m.Min), 999, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4AddClusterRequestScaling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddClusterRequestScaling) UnmarshalBinary(b []byte) error {
	var res V4AddClusterRequestScaling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
