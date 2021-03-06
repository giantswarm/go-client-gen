// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4ClusterDetailsResponseScaling Attributes specific to cluster node scaling as set during cluster
// creation or cluster scaling.
//
// swagger:model v4ClusterDetailsResponseScaling
type V4ClusterDetailsResponseScaling struct {

	// Maximum number of cluster nodes as configured
	//
	Max int64 `json:"max,omitempty"`

	// Minimum number of cluster nodes as configured
	//
	Min *int64 `json:"min,omitempty"`
}

// Validate validates this v4 cluster details response scaling
func (m *V4ClusterDetailsResponseScaling) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4ClusterDetailsResponseScaling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4ClusterDetailsResponseScaling) UnmarshalBinary(b []byte) error {
	var res V4ClusterDetailsResponseScaling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
