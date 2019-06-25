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

// V4CreateAppRequestSpec v4 create app request spec
// swagger:model v4CreateAppRequestSpec
type V4CreateAppRequestSpec struct {

	// The catalog where the chart for this app can be found
	// Required: true
	Catalog *string `json:"catalog"`

	// Name of the chart that should be used to install this app
	// Required: true
	Name *string `json:"name"`

	// Namespace that this app will be installed to
	// Required: true
	Namespace *string `json:"namespace"`

	// user config
	UserConfig *V4CreateAppRequestSpecUserConfig `json:"user_config,omitempty"`

	// Version of the chart that should be used to install this app
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this v4 create app request spec
func (m *V4CreateAppRequestSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4CreateAppRequestSpec) validateCatalog(formats strfmt.Registry) error {

	if err := validate.Required("catalog", "body", m.Catalog); err != nil {
		return err
	}

	return nil
}

func (m *V4CreateAppRequestSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V4CreateAppRequestSpec) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V4CreateAppRequestSpec) validateUserConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.UserConfig) { // not required
		return nil
	}

	if m.UserConfig != nil {
		if err := m.UserConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_config")
			}
			return err
		}
	}

	return nil
}

func (m *V4CreateAppRequestSpec) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4CreateAppRequestSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4CreateAppRequestSpec) UnmarshalBinary(b []byte) error {
	var res V4CreateAppRequestSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
