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

// V5AddNodePoolRequestAvailabilityZones Specifies how the nodes of a pool are spread over availability zones.
// The object must contain either the `number` attribute or the `zones`
// attribute, but not both.
// The maximum `number` of availbility zones is the same as that found
// under `general.availability_zones.max` from the `/v4/info/` endpoint.
// When not given, availability zones assignment is handled automatically.
//
// swagger:model v5AddNodePoolRequestAvailabilityZones
type V5AddNodePoolRequestAvailabilityZones struct {

	// Number of zones to use. If given, the zones are picked
	// automatically. _(Maximum limit of 4 supported.)_
	//
	// Maximum: 4
	Number int64 `json:"number,omitempty"`

	// Names of the availability zones to use. _(Must be same region as the cluster.)_
	//
	// Max Items: 4
	// Min Items: 1
	Zones []string `json:"zones"`
}

// Validate validates this v5 add node pool request availability zones
func (m *V5AddNodePoolRequestAvailabilityZones) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5AddNodePoolRequestAvailabilityZones) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MaximumInt("number", "body", int64(m.Number), 4, false); err != nil {
		return err
	}

	return nil
}

func (m *V5AddNodePoolRequestAvailabilityZones) validateZones(formats strfmt.Registry) error {

	if swag.IsZero(m.Zones) { // not required
		return nil
	}

	iZonesSize := int64(len(m.Zones))

	if err := validate.MinItems("zones", "body", iZonesSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("zones", "body", iZonesSize, 4); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5AddNodePoolRequestAvailabilityZones) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5AddNodePoolRequestAvailabilityZones) UnmarshalBinary(b []byte) error {
	var res V5AddNodePoolRequestAvailabilityZones
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
