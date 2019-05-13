// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V4GetCredentialsResponseItemsAwsRoles IAM roles to assume by certain entities
// swagger:model v4GetCredentialsResponseItemsAwsRoles
type V4GetCredentialsResponseItemsAwsRoles struct {

	// ARN of the IAM role Giant Swarm support staff will use
	Admin string `json:"admin,omitempty"`

	// ARN of the IAM role assumed by the software operating the clusters
	Awsoperator string `json:"awsoperator,omitempty"`
}

// Validate validates this v4 get credentials response items aws roles
func (m *V4GetCredentialsResponseItemsAwsRoles) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4GetCredentialsResponseItemsAwsRoles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4GetCredentialsResponseItemsAwsRoles) UnmarshalBinary(b []byte) error {
	var res V4GetCredentialsResponseItemsAwsRoles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
