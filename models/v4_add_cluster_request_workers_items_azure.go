// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V4AddClusterRequestWorkersItemsAzure Attributes specific to nodes running on Microsoft Azure
//
// swagger:model v4AddClusterRequestWorkersItemsAzure
type V4AddClusterRequestWorkersItemsAzure struct {

	// Azure Virtual Machine size. Must be the same for all worker nodes
	// of a cluster.
	//
	VMSize string `json:"vm_size,omitempty"`
}

// Validate validates this v4 add cluster request workers items azure
func (m *V4AddClusterRequestWorkersItemsAzure) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4AddClusterRequestWorkersItemsAzure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddClusterRequestWorkersItemsAzure) UnmarshalBinary(b []byte) error {
	var res V4AddClusterRequestWorkersItemsAzure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
