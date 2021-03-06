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

// V5AddClusterRequest v5 add cluster request
// swagger:model v5AddClusterRequest
type V5AddClusterRequest struct {

	// master
	Master *V5AddClusterRequestMaster `json:"master,omitempty"`

	// master nodes
	MasterNodes *V5AddClusterRequestMasterNodes `json:"master_nodes,omitempty"`

	// Cluster name
	Name string `json:"name,omitempty"`

	// Name of the organization owning the cluster
	//
	// Required: true
	Owner *string `json:"owner"`

	// The [release](https://docs.giantswarm.io/api/#tag/releases) version
	// to use in the new cluster. If not given, the latest release will be
	// used.
	//
	ReleaseVersion string `json:"release_version,omitempty"`
}

// Validate validates this v5 add cluster request
func (m *V5AddClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5AddClusterRequest) validateMaster(formats strfmt.Registry) error {

	if swag.IsZero(m.Master) { // not required
		return nil
	}

	if m.Master != nil {
		if err := m.Master.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

func (m *V5AddClusterRequest) validateMasterNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.MasterNodes) { // not required
		return nil
	}

	if m.MasterNodes != nil {
		if err := m.MasterNodes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master_nodes")
			}
			return err
		}
	}

	return nil
}

func (m *V5AddClusterRequest) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5AddClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5AddClusterRequest) UnmarshalBinary(b []byte) error {
	var res V5AddClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
