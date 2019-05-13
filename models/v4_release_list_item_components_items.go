// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V4ReleaseListItemComponentsItems v4 release list item components items
// swagger:model v4ReleaseListItemComponentsItems
type V4ReleaseListItemComponentsItems struct {

	// Name of the component
	// Required: true
	Name *string `json:"name"`

	// Version number of the component
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this v4 release list item components items
func (m *V4ReleaseListItemComponentsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
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

func (m *V4ReleaseListItemComponentsItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V4ReleaseListItemComponentsItems) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4ReleaseListItemComponentsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4ReleaseListItemComponentsItems) UnmarshalBinary(b []byte) error {
	var res V4ReleaseListItemComponentsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
