// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V4AddClusterRequest Request model for creating a new cluster
// swagger:model v4AddClusterRequest
type V4AddClusterRequest struct {

	// Number of availability zones a cluster should be spread across. The default is provided via the [info](#operation/getInfo) endpoint.
	AvailabilityZones int64 `json:"availability_zones,omitempty"`

	// Cluster name
	Name string `json:"name,omitempty"`

	// Name of the organization owning the cluster
	// Required: true
	Owner *string `json:"owner"`

	// The [release](https://docs.giantswarm.io/api/#tag/releases) version
	// to use in the new cluster
	//
	ReleaseVersion string `json:"release_version,omitempty"`

	// scaling
	Scaling *V4AddClusterRequestScaling `json:"scaling,omitempty"`

	// Worker node definition. If present, the first item of the array is
	// expected to contain the specification for all worker nodes.
	//
	Workers []*V4AddClusterRequestWorkersItems `json:"workers"`
}

// Validate validates this v4 add cluster request
func (m *V4AddClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScaling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4AddClusterRequest) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *V4AddClusterRequest) validateScaling(formats strfmt.Registry) error {

	if swag.IsZero(m.Scaling) { // not required
		return nil
	}

	if m.Scaling != nil {
		if err := m.Scaling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scaling")
			}
			return err
		}
	}

	return nil
}

func (m *V4AddClusterRequest) validateWorkers(formats strfmt.Registry) error {

	if swag.IsZero(m.Workers) { // not required
		return nil
	}

	for i := 0; i < len(m.Workers); i++ {
		if swag.IsZero(m.Workers[i]) { // not required
			continue
		}

		if m.Workers[i] != nil {
			if err := m.Workers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4AddClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddClusterRequest) UnmarshalBinary(b []byte) error {
	var res V4AddClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
