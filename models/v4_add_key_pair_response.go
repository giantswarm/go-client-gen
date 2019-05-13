// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V4AddKeyPairResponse v4 add key pair response
// swagger:model v4AddKeyPairResponse
type V4AddKeyPairResponse struct {

	// PEM-encoded CA certificate of the cluster
	CertificateAuthorityData string `json:"certificate_authority_data,omitempty"`

	// PEM-encoded certificate
	ClientCertificateData string `json:"client_certificate_data,omitempty"`

	// PEM-encoded RSA private key
	ClientKeyData string `json:"client_key_data,omitempty"`

	// Date/time of creation
	CreateDate string `json:"create_date,omitempty"`

	// Free text information about the key pair
	Description string `json:"description,omitempty"`

	// Unique identifier of the key pair
	ID string `json:"id,omitempty"`

	// Expiration time (from creation) in hours
	TTLHours int64 `json:"ttl_hours,omitempty"`
}

// Validate validates this v4 add key pair response
func (m *V4AddKeyPairResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4AddKeyPairResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddKeyPairResponse) UnmarshalBinary(b []byte) error {
	var res V4AddKeyPairResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
