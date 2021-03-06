// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V4InfoResponseWorkers Information related to worker nodes
// swagger:model v4InfoResponseWorkers
type V4InfoResponseWorkers struct {

	// count per cluster
	CountPerCluster *V4InfoResponseWorkersCountPerCluster `json:"count_per_cluster,omitempty"`

	// instance type
	InstanceType *V4InfoResponseWorkersInstanceType `json:"instance_type,omitempty"`

	// vm size
	VMSize *V4InfoResponseWorkersVMSize `json:"vm_size,omitempty"`
}

// Validate validates this v4 info response workers
func (m *V4InfoResponseWorkers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountPerCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4InfoResponseWorkers) validateCountPerCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.CountPerCluster) { // not required
		return nil
	}

	if m.CountPerCluster != nil {
		if err := m.CountPerCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("count_per_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *V4InfoResponseWorkers) validateInstanceType(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceType) { // not required
		return nil
	}

	if m.InstanceType != nil {
		if err := m.InstanceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance_type")
			}
			return err
		}
	}

	return nil
}

func (m *V4InfoResponseWorkers) validateVMSize(formats strfmt.Registry) error {

	if swag.IsZero(m.VMSize) { // not required
		return nil
	}

	if m.VMSize != nil {
		if err := m.VMSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4InfoResponseWorkers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4InfoResponseWorkers) UnmarshalBinary(b []byte) error {
	var res V4InfoResponseWorkers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
