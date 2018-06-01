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

// BulkPaymentsByPaymentProductPostOKCode is the HTTP code returned for type BulkPaymentsByPaymentProductPostOK
const BulkPaymentsByPaymentProductPostOKCode int = 200

/*BulkPaymentsByPaymentProductPostOK OK

swagger:response bulkPaymentsByPaymentProductPostOK
*/
type BulkPaymentsByPaymentProductPostOK struct {

	/*
	  In: Body
	*/
	Payload *models.BulkPaymentsResponse `json:"body,omitempty"`
}

// NewBulkPaymentsByPaymentProductPostOK creates BulkPaymentsByPaymentProductPostOK with default headers values
func NewBulkPaymentsByPaymentProductPostOK() *BulkPaymentsByPaymentProductPostOK {

	return &BulkPaymentsByPaymentProductPostOK{}
}

// WithPayload adds the payload to the bulk payments by payment product post o k response
func (o *BulkPaymentsByPaymentProductPostOK) WithPayload(payload *models.BulkPaymentsResponse) *BulkPaymentsByPaymentProductPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bulk payments by payment product post o k response
func (o *BulkPaymentsByPaymentProductPostOK) SetPayload(payload *models.BulkPaymentsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BulkPaymentsByPaymentProductPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BulkPaymentsByPaymentProductPostBadRequestCode is the HTTP code returned for type BulkPaymentsByPaymentProductPostBadRequest
const BulkPaymentsByPaymentProductPostBadRequestCode int = 400

/*BulkPaymentsByPaymentProductPostBadRequest Bad Request

swagger:response bulkPaymentsByPaymentProductPostBadRequest
*/
type BulkPaymentsByPaymentProductPostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewBulkPaymentsByPaymentProductPostBadRequest creates BulkPaymentsByPaymentProductPostBadRequest with default headers values
func NewBulkPaymentsByPaymentProductPostBadRequest() *BulkPaymentsByPaymentProductPostBadRequest {

	return &BulkPaymentsByPaymentProductPostBadRequest{}
}

// WithPayload adds the payload to the bulk payments by payment product post bad request response
func (o *BulkPaymentsByPaymentProductPostBadRequest) WithPayload(payload *models.TppMessage) *BulkPaymentsByPaymentProductPostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bulk payments by payment product post bad request response
func (o *BulkPaymentsByPaymentProductPostBadRequest) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BulkPaymentsByPaymentProductPostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BulkPaymentsByPaymentProductPostUnauthorizedCode is the HTTP code returned for type BulkPaymentsByPaymentProductPostUnauthorized
const BulkPaymentsByPaymentProductPostUnauthorizedCode int = 401

/*BulkPaymentsByPaymentProductPostUnauthorized Unauthorized

swagger:response bulkPaymentsByPaymentProductPostUnauthorized
*/
type BulkPaymentsByPaymentProductPostUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewBulkPaymentsByPaymentProductPostUnauthorized creates BulkPaymentsByPaymentProductPostUnauthorized with default headers values
func NewBulkPaymentsByPaymentProductPostUnauthorized() *BulkPaymentsByPaymentProductPostUnauthorized {

	return &BulkPaymentsByPaymentProductPostUnauthorized{}
}

// WithPayload adds the payload to the bulk payments by payment product post unauthorized response
func (o *BulkPaymentsByPaymentProductPostUnauthorized) WithPayload(payload *models.TppMessage) *BulkPaymentsByPaymentProductPostUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bulk payments by payment product post unauthorized response
func (o *BulkPaymentsByPaymentProductPostUnauthorized) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BulkPaymentsByPaymentProductPostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BulkPaymentsByPaymentProductPostForbiddenCode is the HTTP code returned for type BulkPaymentsByPaymentProductPostForbidden
const BulkPaymentsByPaymentProductPostForbiddenCode int = 403

/*BulkPaymentsByPaymentProductPostForbidden Forbidden

swagger:response bulkPaymentsByPaymentProductPostForbidden
*/
type BulkPaymentsByPaymentProductPostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewBulkPaymentsByPaymentProductPostForbidden creates BulkPaymentsByPaymentProductPostForbidden with default headers values
func NewBulkPaymentsByPaymentProductPostForbidden() *BulkPaymentsByPaymentProductPostForbidden {

	return &BulkPaymentsByPaymentProductPostForbidden{}
}

// WithPayload adds the payload to the bulk payments by payment product post forbidden response
func (o *BulkPaymentsByPaymentProductPostForbidden) WithPayload(payload *models.TppMessage) *BulkPaymentsByPaymentProductPostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bulk payments by payment product post forbidden response
func (o *BulkPaymentsByPaymentProductPostForbidden) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BulkPaymentsByPaymentProductPostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BulkPaymentsByPaymentProductPostDefault Internal Server Error

swagger:response bulkPaymentsByPaymentProductPostDefault
*/
type BulkPaymentsByPaymentProductPostDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewBulkPaymentsByPaymentProductPostDefault creates BulkPaymentsByPaymentProductPostDefault with default headers values
func NewBulkPaymentsByPaymentProductPostDefault(code int) *BulkPaymentsByPaymentProductPostDefault {
	if code <= 0 {
		code = 500
	}

	return &BulkPaymentsByPaymentProductPostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the bulk payments by payment product post default response
func (o *BulkPaymentsByPaymentProductPostDefault) WithStatusCode(code int) *BulkPaymentsByPaymentProductPostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the bulk payments by payment product post default response
func (o *BulkPaymentsByPaymentProductPostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the bulk payments by payment product post default response
func (o *BulkPaymentsByPaymentProductPostDefault) WithPayload(payload *models.TppMessage) *BulkPaymentsByPaymentProductPostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bulk payments by payment product post default response
func (o *BulkPaymentsByPaymentProductPostDefault) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BulkPaymentsByPaymentProductPostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
