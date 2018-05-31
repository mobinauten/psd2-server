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
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PaymentService payment-service
// swagger:model PaymentService
type PaymentService string

const (

	// PaymentServicePayments captures enum value "payments"
	PaymentServicePayments PaymentService = "payments"

	// PaymentServiceBulkPayments captures enum value "bulk-payments"
	PaymentServiceBulkPayments PaymentService = "bulk-payments"

	// PaymentServicePeriodicPayments captures enum value "periodic-payments"
	PaymentServicePeriodicPayments PaymentService = "periodic-payments"
)

// for schema
var paymentServiceEnum []interface{}

func init() {
	var res []PaymentService
	if err := json.Unmarshal([]byte(`["payments","bulk-payments","periodic-payments"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentServiceEnum = append(paymentServiceEnum, v)
	}
}

func (m PaymentService) validatePaymentServiceEnum(path, location string, value PaymentService) error {
	if err := validate.Enum(path, location, value, paymentServiceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment service
func (m PaymentService) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentServiceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}