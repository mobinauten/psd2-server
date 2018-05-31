// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsentsRequest consentsRequest
// swagger:model ConsentsRequest
type ConsentsRequest struct {

	// access
	// Required: true
	Access *AccountAccess `json:"access"`

	// If "true" indicates that a payment initiation service will be addressed in the same "session", cp. Section 8.
	// Required: true
	CombinedServiceIndicator *bool `json:"combinedServiceIndicator"`

	// This field indicates the requested maximum frequency for an access per day. For a one-off access, this attribute is set to "1".
	// Required: true
	FrequencyPerDay *int32 `json:"frequencyPerDay"`

	// "true" if the consent is for recurring access to the account data "false", if the consent is for one access to the account data
	// Required: true
	RecurringIndicator *bool `json:"recurringIndicator"`

	// This parameter is requesting a valid until date for the requested consent. The content is the local ASPSP date in ISODate Format, e.g. 2017-10-30. If a maximal available date is requested, a date in far future is to be used: "9999-12-31". The consent object to be retrieved by the GET Consent Request will contain the adjusted date.
	// Required: true
	// Format: date
	ValidUntil *strfmt.Date `json:"validUntil"`
}

// Validate validates this consents request
func (m *ConsentsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCombinedServiceIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyPerDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurringIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentsRequest) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	if m.Access != nil {
		if err := m.Access.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentsRequest) validateCombinedServiceIndicator(formats strfmt.Registry) error {

	if err := validate.Required("combinedServiceIndicator", "body", m.CombinedServiceIndicator); err != nil {
		return err
	}

	return nil
}

func (m *ConsentsRequest) validateFrequencyPerDay(formats strfmt.Registry) error {

	if err := validate.Required("frequencyPerDay", "body", m.FrequencyPerDay); err != nil {
		return err
	}

	return nil
}

func (m *ConsentsRequest) validateRecurringIndicator(formats strfmt.Registry) error {

	if err := validate.Required("recurringIndicator", "body", m.RecurringIndicator); err != nil {
		return err
	}

	return nil
}

func (m *ConsentsRequest) validateValidUntil(formats strfmt.Registry) error {

	if err := validate.Required("validUntil", "body", m.ValidUntil); err != nil {
		return err
	}

	if err := validate.FormatOf("validUntil", "body", "date", m.ValidUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentsRequest) UnmarshalBinary(b []byte) error {
	var res ConsentsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
