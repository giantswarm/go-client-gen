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

// V4AddKeyPairBody v4 add key pair body
// swagger:model V4AddKeyPairBody
type V4AddKeyPairBody struct {

	// certificate organizations
	CertificateOrganizations string `json:"certificate_organizations,omitempty"`

	// cn prefix
	CnPrefix string `json:"cn_prefix,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// ttl hours
	TTLHours int32 `json:"ttl_hours,omitempty"`
}

// Validate validates this v4 add key pair body
func (m *V4AddKeyPairBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4AddKeyPairBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4AddKeyPairBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddKeyPairBody) UnmarshalBinary(b []byte) error {
	var res V4AddKeyPairBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
