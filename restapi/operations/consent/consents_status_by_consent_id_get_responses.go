// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package consent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/openpsd/psd2-server/models"
)

// ConsentsStatusByConsentIDGetOKCode is the HTTP code returned for type ConsentsStatusByConsentIDGetOK
const ConsentsStatusByConsentIDGetOKCode int = 200

/*ConsentsStatusByConsentIDGetOK Transaction status

swagger:response consentsStatusByConsentIdGetOK
*/
type ConsentsStatusByConsentIDGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.ConsentsStatusResponse `json:"body,omitempty"`
}

// NewConsentsStatusByConsentIDGetOK creates ConsentsStatusByConsentIDGetOK with default headers values
func NewConsentsStatusByConsentIDGetOK() *ConsentsStatusByConsentIDGetOK {

	return &ConsentsStatusByConsentIDGetOK{}
}

// WithPayload adds the payload to the consents status by consent Id get o k response
func (o *ConsentsStatusByConsentIDGetOK) WithPayload(payload *models.ConsentsStatusResponse) *ConsentsStatusByConsentIDGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the consents status by consent Id get o k response
func (o *ConsentsStatusByConsentIDGetOK) SetPayload(payload *models.ConsentsStatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConsentsStatusByConsentIDGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConsentsStatusByConsentIDGetBadRequestCode is the HTTP code returned for type ConsentsStatusByConsentIDGetBadRequest
const ConsentsStatusByConsentIDGetBadRequestCode int = 400

/*ConsentsStatusByConsentIDGetBadRequest Bad Request

swagger:response consentsStatusByConsentIdGetBadRequest
*/
type ConsentsStatusByConsentIDGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewConsentsStatusByConsentIDGetBadRequest creates ConsentsStatusByConsentIDGetBadRequest with default headers values
func NewConsentsStatusByConsentIDGetBadRequest() *ConsentsStatusByConsentIDGetBadRequest {

	return &ConsentsStatusByConsentIDGetBadRequest{}
}

// WithPayload adds the payload to the consents status by consent Id get bad request response
func (o *ConsentsStatusByConsentIDGetBadRequest) WithPayload(payload *models.TppMessage) *ConsentsStatusByConsentIDGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the consents status by consent Id get bad request response
func (o *ConsentsStatusByConsentIDGetBadRequest) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConsentsStatusByConsentIDGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConsentsStatusByConsentIDGetUnauthorizedCode is the HTTP code returned for type ConsentsStatusByConsentIDGetUnauthorized
const ConsentsStatusByConsentIDGetUnauthorizedCode int = 401

/*ConsentsStatusByConsentIDGetUnauthorized Unauthorized

swagger:response consentsStatusByConsentIdGetUnauthorized
*/
type ConsentsStatusByConsentIDGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewConsentsStatusByConsentIDGetUnauthorized creates ConsentsStatusByConsentIDGetUnauthorized with default headers values
func NewConsentsStatusByConsentIDGetUnauthorized() *ConsentsStatusByConsentIDGetUnauthorized {

	return &ConsentsStatusByConsentIDGetUnauthorized{}
}

// WithPayload adds the payload to the consents status by consent Id get unauthorized response
func (o *ConsentsStatusByConsentIDGetUnauthorized) WithPayload(payload *models.TppMessage) *ConsentsStatusByConsentIDGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the consents status by consent Id get unauthorized response
func (o *ConsentsStatusByConsentIDGetUnauthorized) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConsentsStatusByConsentIDGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConsentsStatusByConsentIDGetForbiddenCode is the HTTP code returned for type ConsentsStatusByConsentIDGetForbidden
const ConsentsStatusByConsentIDGetForbiddenCode int = 403

/*ConsentsStatusByConsentIDGetForbidden Forbidden

swagger:response consentsStatusByConsentIdGetForbidden
*/
type ConsentsStatusByConsentIDGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewConsentsStatusByConsentIDGetForbidden creates ConsentsStatusByConsentIDGetForbidden with default headers values
func NewConsentsStatusByConsentIDGetForbidden() *ConsentsStatusByConsentIDGetForbidden {

	return &ConsentsStatusByConsentIDGetForbidden{}
}

// WithPayload adds the payload to the consents status by consent Id get forbidden response
func (o *ConsentsStatusByConsentIDGetForbidden) WithPayload(payload *models.TppMessage) *ConsentsStatusByConsentIDGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the consents status by consent Id get forbidden response
func (o *ConsentsStatusByConsentIDGetForbidden) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConsentsStatusByConsentIDGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ConsentsStatusByConsentIDGetDefault Internal Server Error

swagger:response consentsStatusByConsentIdGetDefault
*/
type ConsentsStatusByConsentIDGetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.TppMessage `json:"body,omitempty"`
}

// NewConsentsStatusByConsentIDGetDefault creates ConsentsStatusByConsentIDGetDefault with default headers values
func NewConsentsStatusByConsentIDGetDefault(code int) *ConsentsStatusByConsentIDGetDefault {
	if code <= 0 {
		code = 500
	}

	return &ConsentsStatusByConsentIDGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the consents status by consent Id get default response
func (o *ConsentsStatusByConsentIDGetDefault) WithStatusCode(code int) *ConsentsStatusByConsentIDGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the consents status by consent Id get default response
func (o *ConsentsStatusByConsentIDGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the consents status by consent Id get default response
func (o *ConsentsStatusByConsentIDGetDefault) WithPayload(payload *models.TppMessage) *ConsentsStatusByConsentIDGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the consents status by consent Id get default response
func (o *ConsentsStatusByConsentIDGetDefault) SetPayload(payload *models.TppMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConsentsStatusByConsentIDGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
