// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V4ReleaseListItem v4 release list item
// swagger:model v4ReleaseListItem
type V4ReleaseListItem struct {

	// If true, the version is available for new clusters and cluster
	// upgrades. Older versions become unavailable and thus have the
	// value `false` here.
	//
	Active bool `json:"active,omitempty"`

	// Structured list of changes in this release, in comparison to the
	// previous version, with respect to the contained components.
	//
	// Required: true
	Changelog []*V4ReleaseListItemChangelogItems `json:"changelog"`

	// List of components and their version contained in the release
	//
	// Required: true
	Components []*V4ReleaseListItemComponentsItems `json:"components"`

	// Date and time of the release creation
	// Required: true
	Timestamp *string `json:"timestamp"`

	// The semantic version number
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this v4 release list item
func (m *V4ReleaseListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangelog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
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

func (m *V4ReleaseListItem) validateChangelog(formats strfmt.Registry) error {

	if err := validate.Required("changelog", "body", m.Changelog); err != nil {
		return err
	}

	for i := 0; i < len(m.Changelog); i++ {
		if swag.IsZero(m.Changelog[i]) { // not required
			continue
		}

		if m.Changelog[i] != nil {
			if err := m.Changelog[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changelog" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V4ReleaseListItem) validateComponents(formats strfmt.Registry) error {

	if err := validate.Required("components", "body", m.Components); err != nil {
		return err
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V4ReleaseListItem) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *V4ReleaseListItem) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4ReleaseListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4ReleaseListItem) UnmarshalBinary(b []byte) error {
	var res V4ReleaseListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
