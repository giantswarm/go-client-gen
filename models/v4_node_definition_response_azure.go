// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4NodeDefinitionResponseAzure v4 node definition response azure
// swagger:model v4NodeDefinitionResponseAzure
type V4NodeDefinitionResponseAzure struct {

	// vm size
	VMSize string `json:"vm_size,omitempty"`
}

// Validate validates this v4 node definition response azure
func (m *V4NodeDefinitionResponseAzure) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4NodeDefinitionResponseAzure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4NodeDefinitionResponseAzure) UnmarshalBinary(b []byte) error {
	var res V4NodeDefinitionResponseAzure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}