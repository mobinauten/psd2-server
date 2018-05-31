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
)

// AccountsTransactionsResponse81 Accounts Transactions Response81
// swagger:model AccountsTransactionsResponse81
type AccountsTransactionsResponse81 struct {

	// transactions details
	TransactionsDetails *Transaction `json:"transactionsDetails,omitempty"`
}

// Validate validates this accounts transactions response81
func (m *AccountsTransactionsResponse81) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransactionsDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsTransactionsResponse81) validateTransactionsDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionsDetails) { // not required
		return nil
	}

	if m.TransactionsDetails != nil {
		if err := m.TransactionsDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactionsDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountsTransactionsResponse81) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountsTransactionsResponse81) UnmarshalBinary(b []byte) error {
	var res AccountsTransactionsResponse81
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
