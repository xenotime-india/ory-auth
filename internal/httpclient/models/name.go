// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Name Name Name Name represents an X.509 distinguished name. This only includes the common
// elements of a DN. When parsing, all elements are stored in Names and
// non-standard elements can be extracted from there. When marshaling, elements
// in ExtraNames are appended and override other values with the same OID.
// swagger:model Name
type Name struct {

	// country
	Country []string `json:"Country"`

	// extra names
	ExtraNames []*AttributeTypeAndValue `json:"ExtraNames"`

	// locality
	Locality []string `json:"Locality"`

	// names
	Names []*AttributeTypeAndValue `json:"Names"`

	// serial number
	SerialNumber string `json:"SerialNumber,omitempty"`

	// street address
	StreetAddress []string `json:"StreetAddress"`
}

// Validate validates this name
func (m *Name) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Name) validateExtraNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtraNames) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtraNames); i++ {
		if swag.IsZero(m.ExtraNames[i]) { // not required
			continue
		}

		if m.ExtraNames[i] != nil {
			if err := m.ExtraNames[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExtraNames" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Name) validateNames(formats strfmt.Registry) error {

	if swag.IsZero(m.Names) { // not required
		return nil
	}

	for i := 0; i < len(m.Names); i++ {
		if swag.IsZero(m.Names[i]) { // not required
			continue
		}

		if m.Names[i] != nil {
			if err := m.Names[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Names" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Name) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Name) UnmarshalBinary(b []byte) error {
	var res Name
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
