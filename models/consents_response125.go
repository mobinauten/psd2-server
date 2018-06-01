// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
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

// ConsentsResponse125 Consents Response125
// swagger:model ConsentsResponse125
type ConsentsResponse125 struct {

	// access
	// Required: true
	Access *AccountAccess `json:"access"`

	// consent status
	// Required: true
	ConsentStatus ConsentStatus `json:"consentStatus"`

	// frequency per day
	// Required: true
	FrequencyPerDay *int32 `json:"frequencyPerDay"`

	// last action date
	// Required: true
	// Format: date
	LastActionDate *strfmt.Date `json:"lastActionDate"`

	// "true" if the consent is for recurring access to the account data "false", if the consent is for one access to the account data
	// Required: true
	RecurringIndicator *bool `json:"recurringIndicator"`

	// This data element is containing information about the status of the SCA method applied. This is free text but might be coded in a future version of the specification.
	ScaStatus string `json:"scaStatus,omitempty"`

	// valid until
	// Required: true
	// Format: date
	ValidUntil *strfmt.Date `json:"validUntil"`
}

// Validate validates this consents response125
func (m *ConsentsResponse125) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyPerDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActionDate(formats); err != nil {
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

func (m *ConsentsResponse125) validateAccess(formats strfmt.Registry) error {

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

func (m *ConsentsResponse125) validateConsentStatus(formats strfmt.Registry) error {

	if err := m.ConsentStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("consentStatus")
		}
		return err
	}

	return nil
}

func (m *ConsentsResponse125) validateFrequencyPerDay(formats strfmt.Registry) error {

	if err := validate.Required("frequencyPerDay", "body", m.FrequencyPerDay); err != nil {
		return err
	}

	return nil
}

func (m *ConsentsResponse125) validateLastActionDate(formats strfmt.Registry) error {

	if err := validate.Required("lastActionDate", "body", m.LastActionDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastActionDate", "body", "date", m.LastActionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsentsResponse125) validateRecurringIndicator(formats strfmt.Registry) error {

	if err := validate.Required("recurringIndicator", "body", m.RecurringIndicator); err != nil {
		return err
	}

	return nil
}

func (m *ConsentsResponse125) validateValidUntil(formats strfmt.Registry) error {

	if err := validate.Required("validUntil", "body", m.ValidUntil); err != nil {
		return err
	}

	if err := validate.FormatOf("validUntil", "body", "date", m.ValidUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentsResponse125) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentsResponse125) UnmarshalBinary(b []byte) error {
	var res ConsentsResponse125
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
