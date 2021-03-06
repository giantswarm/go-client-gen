// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4GetKeyPairsResponseItems v4 get key pairs response items
// swagger:model v4GetKeyPairsResponseItems
type V4GetKeyPairsResponseItems struct {

	// The certificate subject's `organization` fields.
	CertificateOrganizations string `json:"certificate_organizations,omitempty"`

	// The common name of the certificate subject.
	CommonName string `json:"common_name,omitempty"`

	// Date/time of creation
	CreateDate string `json:"create_date,omitempty"`

	// Free text information about the key pair
	Description string `json:"description,omitempty"`

	// Unique identifier of the key pair
	ID string `json:"id,omitempty"`

	// Expiration time (from creation) in hours
	TTLHours int64 `json:"ttl_hours,omitempty"`
}

// Validate validates this v4 get key pairs response items
func (m *V4GetKeyPairsResponseItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4GetKeyPairsResponseItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4GetKeyPairsResponseItems) UnmarshalBinary(b []byte) error {
	var res V4GetKeyPairsResponseItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
