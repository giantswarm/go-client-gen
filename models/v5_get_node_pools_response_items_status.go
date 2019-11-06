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

// V5GetNodePoolsResponseItemsStatus Information on the current size and status of the node pool
// swagger:model v5GetNodePoolsResponseItemsStatus
type V5GetNodePoolsResponseItemsStatus struct {

	// Desired number of nodes in the pool
	// Minimum: 0
	Nodes *int64 `json:"nodes,omitempty"`

	// Number of nodes in state NodeReady
	// Minimum: 0
	NodesReady *int64 `json:"nodes_ready,omitempty"`
}

// Validate validates this v5 get node pools response items status
func (m *V5GetNodePoolsResponseItemsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodesReady(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5GetNodePoolsResponseItemsStatus) validateNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	if err := validate.MinimumInt("nodes", "body", int64(*m.Nodes), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *V5GetNodePoolsResponseItemsStatus) validateNodesReady(formats strfmt.Registry) error {

	if swag.IsZero(m.NodesReady) { // not required
		return nil
	}

	if err := validate.MinimumInt("nodes_ready", "body", int64(*m.NodesReady), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5GetNodePoolsResponseItemsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5GetNodePoolsResponseItemsStatus) UnmarshalBinary(b []byte) error {
	var res V5GetNodePoolsResponseItemsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
