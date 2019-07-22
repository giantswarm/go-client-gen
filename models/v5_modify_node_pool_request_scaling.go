// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5ModifyNodePoolRequestScaling Attributes specific to node pool scaling. To have full control of
// the cluster size, min and max can be set to the same value. If only
// `min` or only `max` is specified, `min` and `max` will be set equally.
//
// swagger:model v5ModifyNodePoolRequestScaling
type V5ModifyNodePoolRequestScaling struct {

	// Maximum number of nodes in the pool
	Max int64 `json:"max,omitempty"`

	// Minimum number of nodes in the pool
	Min int64 `json:"min,omitempty"`
}

// Validate validates this v5 modify node pool request scaling
func (m *V5ModifyNodePoolRequestScaling) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5ModifyNodePoolRequestScaling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5ModifyNodePoolRequestScaling) UnmarshalBinary(b []byte) error {
	var res V5ModifyNodePoolRequestScaling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
