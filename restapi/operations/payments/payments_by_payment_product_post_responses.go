// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/openpsd/psd2-server/models"
)

// PaymentsByPaymentProductPostOKCode is the HTTP code returned for type PaymentsByPaymentProductPostOK
const PaymentsByPaymentProductPostOKCode int = 200

/*PaymentsByPaymentProductPostOK OK

swagger:response paymentsByPaymentProductPostOK
*/
type PaymentsByPaymentProductPostOK struct {

	/*
	  In: Body
	*/
	Payload *models.PaymentsResponse `json:"body,omitempty"`
}

// NewPaymentsByPaymentProductPostOK creates PaymentsByPaymentProductPostOK with default headers values
func NewPaymentsByPaymentProductPostOK() *PaymentsByPaymentProductPostOK {

	return &PaymentsByPaymentProductPostOK{}
}

// WithPayload adds the payload to the payments by payment product post o k response
func (o *PaymentsByPaymentProductPostOK) WithPayload(payload *models.PaymentsResponse) *PaymentsByPaymentProductPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payments by payment product post o k response
func (o *PaymentsByPaymentProductPostOK) SetPayload(payload *models.PaymentsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentsByPaymentProductPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PaymentsByPaymentProductPostBadRequestCode is the HTTP code returned for type PaymentsByPaymentProductPostBadRequest
const PaymentsByPaymentProductPostBadRequestCode int = 400

/*PaymentsByPaymentProductPostBadRequest Bad Request

swagger:response paymentsByPaymentProductPostBadRequest
*/
type PaymentsByPaymentProductPostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentsByPaymentProductPostBadRequest creates PaymentsByPaymentProductPostBadRequest with default headers values
func NewPaymentsByPaymentProductPostBadRequest() *PaymentsByPaymentProductPostBadRequest {

	return &PaymentsByPaymentProductPostBadRequest{}
}

// WithPayload adds the payload to the payments by payment product post bad request response
func (o *PaymentsByPaymentProductPostBadRequest) WithPayload(payload *models.TppMessage) *PaymentsByPaymentProductPostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payments by payment product post bad request response
func (o *PaymentsByPaymentProductPostBadRequest) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentsByPaymentProductPostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PaymentsByPaymentProductPostUnauthorizedCode is the HTTP code returned for type PaymentsByPaymentProductPostUnauthorized
const PaymentsByPaymentProductPostUnauthorizedCode int = 401

/*PaymentsByPaymentProductPostUnauthorized Unauthorized

swagger:response paymentsByPaymentProductPostUnauthorized
*/
type PaymentsByPaymentProductPostUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentsByPaymentProductPostUnauthorized creates PaymentsByPaymentProductPostUnauthorized with default headers values
func NewPaymentsByPaymentProductPostUnauthorized() *PaymentsByPaymentProductPostUnauthorized {

	return &PaymentsByPaymentProductPostUnauthorized{}
}

// WithPayload adds the payload to the payments by payment product post unauthorized response
func (o *PaymentsByPaymentProductPostUnauthorized) WithPayload(payload *models.TppMessage) *PaymentsByPaymentProductPostUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payments by payment product post unauthorized response
func (o *PaymentsByPaymentProductPostUnauthorized) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentsByPaymentProductPostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PaymentsByPaymentProductPostForbiddenCode is the HTTP code returned for type PaymentsByPaymentProductPostForbidden
const PaymentsByPaymentProductPostForbiddenCode int = 403

/*PaymentsByPaymentProductPostForbidden Forbidden

swagger:response paymentsByPaymentProductPostForbidden
*/
type PaymentsByPaymentProductPostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentsByPaymentProductPostForbidden creates PaymentsByPaymentProductPostForbidden with default headers values
func NewPaymentsByPaymentProductPostForbidden() *PaymentsByPaymentProductPostForbidden {

	return &PaymentsByPaymentProductPostForbidden{}
}

// WithPayload adds the payload to the payments by payment product post forbidden response
func (o *PaymentsByPaymentProductPostForbidden) WithPayload(payload *models.TppMessage) *PaymentsByPaymentProductPostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payments by payment product post forbidden response
func (o *PaymentsByPaymentProductPostForbidden) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentsByPaymentProductPostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PaymentsByPaymentProductPostDefault Internal Server Error

swagger:response paymentsByPaymentProductPostDefault
*/
type PaymentsByPaymentProductPostDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentsByPaymentProductPostDefault creates PaymentsByPaymentProductPostDefault with default headers values
func NewPaymentsByPaymentProductPostDefault(code int) *PaymentsByPaymentProductPostDefault {
	if code <= 0 {
		code = 500
	}

	return &PaymentsByPaymentProductPostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the payments by payment product post default response
func (o *PaymentsByPaymentProductPostDefault) WithStatusCode(code int) *PaymentsByPaymentProductPostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the payments by payment product post default response
func (o *PaymentsByPaymentProductPostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the payments by payment product post default response
func (o *PaymentsByPaymentProductPostDefault) WithPayload(payload *models.TppMessage) *PaymentsByPaymentProductPostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payments by payment product post default response
func (o *PaymentsByPaymentProductPostDefault) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentsByPaymentProductPostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
