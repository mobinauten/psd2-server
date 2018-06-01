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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PeriodicPaymentsResponse Periodic Payments Response
// swagger:model PeriodicPaymentsResponse
type PeriodicPaymentsResponse struct {

	// links
	// Required: true
	Links map[string]interface{} `json:"_links"`

	// challenge data
	ChallengeData *Challenge `json:"challengeData,omitempty"`

	// chosen sca method
	ChosenScaMethod *Authentication `json:"chosenScaMethod,omitempty"`

	// resource identification of the generated payment initiation resource.
	// Required: true
	PaymentID *string `json:"paymentId"`

	// Text to be displayed to the PSU
	PsuMessage string `json:"psuMessage,omitempty"`

	// sca methods
	ScaMethods []*Authentication `json:"scaMethods"`

	// tpp messages
	TppMessages []*TppMessage261 `json:"tppMessages"`

	// transaction fee indicator
	TransactionFeeIndicator bool `json:"transactionFeeIndicator,omitempty"`

	// transaction fees
	TransactionFees *Amount `json:"transactionFees,omitempty"`

	// transaction status
	// Required: true
	TransactionStatus TransactionStatus `json:"transactionStatus"`
}

// Validate validates this periodic payments response
func (m *PeriodicPaymentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChallengeData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChosenScaMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScaMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTppMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionFees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeriodicPaymentsResponse) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *PeriodicPaymentsResponse) validateChallengeData(formats strfmt.Registry) error {

	if swag.IsZero(m.ChallengeData) { // not required
		return nil
	}

	if m.ChallengeData != nil {
		if err := m.ChallengeData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("challengeData")
			}
			return err
		}
	}

	return nil
}

func (m *PeriodicPaymentsResponse) validateChosenScaMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.ChosenScaMethod) { // not required
		return nil
	}

	if m.ChosenScaMethod != nil {
		if err := m.ChosenScaMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chosenScaMethod")
			}
			return err
		}
	}

	return nil
}

func (m *PeriodicPaymentsResponse) validatePaymentID(formats strfmt.Registry) error {

	if err := validate.Required("paymentId", "body", m.PaymentID); err != nil {
		return err
	}

	return nil
}

func (m *PeriodicPaymentsResponse) validateScaMethods(formats strfmt.Registry) error {

	if swag.IsZero(m.ScaMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.ScaMethods); i++ {
		if swag.IsZero(m.ScaMethods[i]) { // not required
			continue
		}

		if m.ScaMethods[i] != nil {
			if err := m.ScaMethods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scaMethods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PeriodicPaymentsResponse) validateTppMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.TppMessages) { // not required
		return nil
	}

	for i := 0; i < len(m.TppMessages); i++ {
		if swag.IsZero(m.TppMessages[i]) { // not required
			continue
		}

		if m.TppMessages[i] != nil {
			if err := m.TppMessages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tppMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PeriodicPaymentsResponse) validateTransactionFees(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionFees) { // not required
		return nil
	}

	if m.TransactionFees != nil {
		if err := m.TransactionFees.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactionFees")
			}
			return err
		}
	}

	return nil
}

func (m *PeriodicPaymentsResponse) validateTransactionStatus(formats strfmt.Registry) error {

	if err := m.TransactionStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactionStatus")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeriodicPaymentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeriodicPaymentsResponse) UnmarshalBinary(b []byte) error {
	var res PeriodicPaymentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
