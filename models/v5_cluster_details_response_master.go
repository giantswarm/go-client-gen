// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5ClusterDetailsResponseMaster Information about the master node
//
// swagger:model v5ClusterDetailsResponseMaster
type V5ClusterDetailsResponseMaster struct {

	// Name of the availability zone the master node is placed in
	//
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

// Validate validates this v5 cluster details response master
func (m *V5ClusterDetailsResponseMaster) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5ClusterDetailsResponseMaster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5ClusterDetailsResponseMaster) UnmarshalBinary(b []byte) error {
	var res V5ClusterDetailsResponseMaster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
