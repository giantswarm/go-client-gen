// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4NodeDefinitionResponseCPU v4 node definition response Cpu
// swagger:model v4NodeDefinitionResponseCpu
type V4NodeDefinitionResponseCPU struct {

	// Number of CPU cores
	Cores int64 `json:"cores,omitempty"`
}

// Validate validates this v4 node definition response Cpu
func (m *V4NodeDefinitionResponseCPU) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4NodeDefinitionResponseCPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4NodeDefinitionResponseCPU) UnmarshalBinary(b []byte) error {
	var res V4NodeDefinitionResponseCPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
