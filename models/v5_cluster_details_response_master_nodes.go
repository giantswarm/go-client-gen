// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5ClusterDetailsResponseMasterNodes Information on the master node(s) of this cluster
//
// swagger:model v5ClusterDetailsResponseMasterNodes
type V5ClusterDetailsResponseMasterNodes struct {

	// Availability zones of the master node(s) of this cluster.
	//
	AvailabilityZones []string `json:"availability_zones"`

	// When true, the cluster has (or should have) three master nodes.
	// Otherwise it should have one.
	//
	HighAvailability bool `json:"high_availability,omitempty"`

	// Number of master nodes that are reported as `Ready`.
	//
	NumReady *int8 `json:"num_ready,omitempty"`
}

// Validate validates this v5 cluster details response master nodes
func (m *V5ClusterDetailsResponseMasterNodes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5ClusterDetailsResponseMasterNodes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5ClusterDetailsResponseMasterNodes) UnmarshalBinary(b []byte) error {
	var res V5ClusterDetailsResponseMasterNodes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
