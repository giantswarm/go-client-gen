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

// V5ClusterDetailsResponse v5 cluster details response
// swagger:model v5ClusterDetailsResponse
type V5ClusterDetailsResponse struct {

	// URI of the Kubernetes API endpoint
	APIEndpoint string `json:"api_endpoint,omitempty"`

	// List of conditions the cluster has gone through
	Conditions []*V5ClusterDetailsResponseConditionsItems `json:"conditions"`

	// Date/time of cluster creation
	CreateDate string `json:"create_date,omitempty"`

	// ID of the credentials used to operate the cluster. See
	// [Set credentials](#operation/addCredentials) for details.
	//
	CredentialID string `json:"credential_id,omitempty"`

	// Date/time when cluster has been deleted
	// Format: date-time
	DeleteDate *strfmt.DateTime `json:"delete_date,omitempty"`

	// Unique cluster identifier
	ID string `json:"id,omitempty"`

	// Labels object
	//
	// Object containing keys with string values representing the labels attached to the cluster
	Labels map[string]string `json:"labels,omitempty"`

	// master
	Master *V5ClusterDetailsResponseMaster `json:"master,omitempty"`

	// Cluster name
	Name string `json:"name,omitempty"`

	// Name of the organization owning the cluster
	//
	Owner string `json:"owner,omitempty"`

	// The [release](https://docs.giantswarm.io/api/#tag/releases) version
	// used by the cluster
	//
	ReleaseVersion string `json:"release_version,omitempty"`

	// versions
	Versions []*V5ClusterDetailsResponseVersionsItems `json:"versions"`
}

// Validate validates this v5 cluster details response
func (m *V5ClusterDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5ClusterDetailsResponse) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V5ClusterDetailsResponse) validateDeleteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeleteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("delete_date", "body", "date-time", m.DeleteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V5ClusterDetailsResponse) validateMaster(formats strfmt.Registry) error {

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

func (m *V5ClusterDetailsResponse) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5ClusterDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5ClusterDetailsResponse) UnmarshalBinary(b []byte) error {
	var res V5ClusterDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
