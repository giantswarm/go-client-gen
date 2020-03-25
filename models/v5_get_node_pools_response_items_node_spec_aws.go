// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5GetNodePoolsResponseItemsNodeSpecAws Attributes specific to the AWS provider
//
// swagger:model v5GetNodePoolsResponseItemsNodeSpecAws
type V5GetNodePoolsResponseItemsNodeSpecAws struct {

	// EC2 instance type used by all nodes in this pool
	//
	InstanceType string `json:"instance_type,omitempty"`

	// If true, the node pool will be filled with spot instances.
	//
	SpotInstanceEnabled bool `json:"spot_instance_enabled,omitempty"`
}

// Validate validates this v5 get node pools response items node spec aws
func (m *V5GetNodePoolsResponseItemsNodeSpecAws) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5GetNodePoolsResponseItemsNodeSpecAws) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5GetNodePoolsResponseItemsNodeSpecAws) UnmarshalBinary(b []byte) error {
	var res V5GetNodePoolsResponseItemsNodeSpecAws
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
