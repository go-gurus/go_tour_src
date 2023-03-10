// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Beer beer
//
// swagger:model beer
type Beer struct {

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// origin
	// Required: true
	// Min Length: 1
	Origin *string `json:"origin"`

	// title
	// Required: true
	// Min Length: 1
	Title *string `json:"title"`

	// volume percentage
	// Required: true
	VolumePercentage *float32 `json:"volume-percentage"`
}

// Validate validates this beer
func (m *Beer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumePercentage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Beer) validateOrigin(formats strfmt.Registry) error {

	if err := validate.Required("origin", "body", m.Origin); err != nil {
		return err
	}

	if err := validate.MinLength("origin", "body", *m.Origin, 1); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateVolumePercentage(formats strfmt.Registry) error {

	if err := validate.Required("volume-percentage", "body", m.VolumePercentage); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this beer based on the context it is used
func (m *Beer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Beer) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Beer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Beer) UnmarshalBinary(b []byte) error {
	var res Beer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
