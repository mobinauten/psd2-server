// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
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

// PaymentIDByPaymentServiceAndPaymentProductGetOKCode is the HTTP code returned for type PaymentIDByPaymentServiceAndPaymentProductGetOK
const PaymentIDByPaymentServiceAndPaymentProductGetOKCode int = 200

/*PaymentIDByPaymentServiceAndPaymentProductGetOK OK

swagger:response paymentIdByPaymentServiceAndPaymentProductGetOK
*/
type PaymentIDByPaymentServiceAndPaymentProductGetOK struct {

	/*
	  In: Body
	*/
	Payload PaymentIDByPaymentServiceAndPaymentProductGetOKBody `json:"body,omitempty"`
}

// NewPaymentIDByPaymentServiceAndPaymentProductGetOK creates PaymentIDByPaymentServiceAndPaymentProductGetOK with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductGetOK() *PaymentIDByPaymentServiceAndPaymentProductGetOK {

	return &PaymentIDByPaymentServiceAndPaymentProductGetOK{}
}

// WithPayload adds the payload to the payment Id by payment service and payment product get o k response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetOK) WithPayload(payload PaymentIDByPaymentServiceAndPaymentProductGetOKBody) *PaymentIDByPaymentServiceAndPaymentProductGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payment Id by payment service and payment product get o k response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetOK) SetPayload(payload PaymentIDByPaymentServiceAndPaymentProductGetOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentIDByPaymentServiceAndPaymentProductGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PaymentIDByPaymentServiceAndPaymentProductGetBadRequestCode is the HTTP code returned for type PaymentIDByPaymentServiceAndPaymentProductGetBadRequest
const PaymentIDByPaymentServiceAndPaymentProductGetBadRequestCode int = 400

/*PaymentIDByPaymentServiceAndPaymentProductGetBadRequest Bad Request

swagger:response paymentIdByPaymentServiceAndPaymentProductGetBadRequest
*/
type PaymentIDByPaymentServiceAndPaymentProductGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentIDByPaymentServiceAndPaymentProductGetBadRequest creates PaymentIDByPaymentServiceAndPaymentProductGetBadRequest with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductGetBadRequest() *PaymentIDByPaymentServiceAndPaymentProductGetBadRequest {

	return &PaymentIDByPaymentServiceAndPaymentProductGetBadRequest{}
}

// WithPayload adds the payload to the payment Id by payment service and payment product get bad request response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetBadRequest) WithPayload(payload *models.TppMessage) *PaymentIDByPaymentServiceAndPaymentProductGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payment Id by payment service and payment product get bad request response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetBadRequest) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentIDByPaymentServiceAndPaymentProductGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PaymentIDByPaymentServiceAndPaymentProductGetUnauthorizedCode is the HTTP code returned for type PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized
const PaymentIDByPaymentServiceAndPaymentProductGetUnauthorizedCode int = 401

/*PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized Unauthorized

swagger:response paymentIdByPaymentServiceAndPaymentProductGetUnauthorized
*/
type PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentIDByPaymentServiceAndPaymentProductGetUnauthorized creates PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductGetUnauthorized() *PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized {

	return &PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized{}
}

// WithPayload adds the payload to the payment Id by payment service and payment product get unauthorized response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized) WithPayload(payload *models.TppMessage) *PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payment Id by payment service and payment product get unauthorized response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentIDByPaymentServiceAndPaymentProductGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PaymentIDByPaymentServiceAndPaymentProductGetForbiddenCode is the HTTP code returned for type PaymentIDByPaymentServiceAndPaymentProductGetForbidden
const PaymentIDByPaymentServiceAndPaymentProductGetForbiddenCode int = 403

/*PaymentIDByPaymentServiceAndPaymentProductGetForbidden Forbidden

swagger:response paymentIdByPaymentServiceAndPaymentProductGetForbidden
*/
type PaymentIDByPaymentServiceAndPaymentProductGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentIDByPaymentServiceAndPaymentProductGetForbidden creates PaymentIDByPaymentServiceAndPaymentProductGetForbidden with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductGetForbidden() *PaymentIDByPaymentServiceAndPaymentProductGetForbidden {

	return &PaymentIDByPaymentServiceAndPaymentProductGetForbidden{}
}

// WithPayload adds the payload to the payment Id by payment service and payment product get forbidden response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetForbidden) WithPayload(payload *models.TppMessage) *PaymentIDByPaymentServiceAndPaymentProductGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payment Id by payment service and payment product get forbidden response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetForbidden) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentIDByPaymentServiceAndPaymentProductGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PaymentIDByPaymentServiceAndPaymentProductGetDefault Internal Server Error

swagger:response paymentIdByPaymentServiceAndPaymentProductGetDefault
*/
type PaymentIDByPaymentServiceAndPaymentProductGetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewPaymentIDByPaymentServiceAndPaymentProductGetDefault creates PaymentIDByPaymentServiceAndPaymentProductGetDefault with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductGetDefault(code int) *PaymentIDByPaymentServiceAndPaymentProductGetDefault {
	if code <= 0 {
		code = 500
	}

	return &PaymentIDByPaymentServiceAndPaymentProductGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the payment Id by payment service and payment product get default response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetDefault) WithStatusCode(code int) *PaymentIDByPaymentServiceAndPaymentProductGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the payment Id by payment service and payment product get default response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the payment Id by payment service and payment product get default response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetDefault) WithPayload(payload *models.TppMessage) *PaymentIDByPaymentServiceAndPaymentProductGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the payment Id by payment service and payment product get default response
func (o *PaymentIDByPaymentServiceAndPaymentProductGetDefault) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PaymentIDByPaymentServiceAndPaymentProductGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
