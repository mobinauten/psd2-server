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

// Balance Balance
// swagger:model Balance
type Balance struct {

	// balance amount
	// Required: true
	BalanceAmount *Amount `json:"balanceAmount"`

	// balance type
	// Required: true
	BalanceType BalanceType `json:"balanceType"`

	// last change date time
	// Format: date-time
	LastChangeDateTime strfmt.DateTime `json:"lastChangeDateTime,omitempty"`

	// last committed transaction
	// Max Length: 35
	LastCommittedTransaction string `json:"lastCommittedTransaction,omitempty"`

	// reference date
	// Format: date
	ReferenceDate strfmt.Date `json:"referenceDate,omitempty"`
}

// Validate validates this balance
func (m *Balance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalanceAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastChangeDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastCommittedTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Balance) validateBalanceAmount(formats strfmt.Registry) error {

	if err := validate.Required("balanceAmount", "body", m.BalanceAmount); err != nil {
		return err
	}

	if m.BalanceAmount != nil {
		if err := m.BalanceAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balanceAmount")
			}
			return err
		}
	}

	return nil
}

func (m *Balance) validateBalanceType(formats strfmt.Registry) error {

	if err := m.BalanceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("balanceType")
		}
		return err
	}

	return nil
}

func (m *Balance) validateLastChangeDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastChangeDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastChangeDateTime", "body", "date-time", m.LastChangeDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateLastCommittedTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.LastCommittedTransaction) { // not required
		return nil
	}

	if err := validate.MaxLength("lastCommittedTransaction", "body", string(m.LastCommittedTransaction), 35); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateReferenceDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceDate) { // not required
		return nil
	}

	if err := validate.FormatOf("referenceDate", "body", "date", m.ReferenceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Balance) UnmarshalBinary(b []byte) error {
	var res Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
