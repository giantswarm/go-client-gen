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

// V4AddKeyPairRequest v4 add key pair request
// swagger:model v4AddKeyPairRequest
type V4AddKeyPairRequest struct {

	// This will set the certificate subject's `organization` fields.
	// Use a comma seperated list of values.
	//
	CertificateOrganizations string `json:"certificate_organizations,omitempty"`

	// The common name prefix of the certificate subject. This only allows characters that are usable in domain names (`a-z`, `0-9`, and `.-`, where `.-` must not occur at either the start or the end).
	CnPrefix string `json:"cn_prefix,omitempty"`

	// Free text information about the key pair
	// Required: true
	Description *string `json:"description"`

	// Expiration time (from creation) in hours
	TTLHours int32 `json:"ttl_hours,omitempty"`
}

// Validate validates this v4 add key pair request
func (m *V4AddKeyPairRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4AddKeyPairRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4AddKeyPairRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddKeyPairRequest) UnmarshalBinary(b []byte) error {
	var res V4AddKeyPairRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
