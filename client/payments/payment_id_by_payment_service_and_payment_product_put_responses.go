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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openpsd/psd2-server/models"
)

// PaymentIDByPaymentServiceAndPaymentProductPutReader is a Reader for the PaymentIDByPaymentServiceAndPaymentProductPut structure.
type PaymentIDByPaymentServiceAndPaymentProductPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentIDByPaymentServiceAndPaymentProductPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPaymentIDByPaymentServiceAndPaymentProductPutCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPaymentIDByPaymentServiceAndPaymentProductPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPaymentIDByPaymentServiceAndPaymentProductPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPaymentIDByPaymentServiceAndPaymentProductPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPaymentIDByPaymentServiceAndPaymentProductPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutCreated creates a PaymentIDByPaymentServiceAndPaymentProductPutCreated with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductPutCreated() *PaymentIDByPaymentServiceAndPaymentProductPutCreated {
	return &PaymentIDByPaymentServiceAndPaymentProductPutCreated{}
}

/*PaymentIDByPaymentServiceAndPaymentProductPutCreated handles this case with default header values.

OK result with Decoupled Approach
*/
type PaymentIDByPaymentServiceAndPaymentProductPutCreated struct {
	Payload *models.PaymentidResponse
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutCreated) Error() string {
	return fmt.Sprintf("[PUT /{payment-service}/{payment-product}/{paymentId}][%d] paymentIdByPaymentServiceAndPaymentProductPutCreated  %+v", 201, o.Payload)
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentidResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutBadRequest creates a PaymentIDByPaymentServiceAndPaymentProductPutBadRequest with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductPutBadRequest() *PaymentIDByPaymentServiceAndPaymentProductPutBadRequest {
	return &PaymentIDByPaymentServiceAndPaymentProductPutBadRequest{}
}

/*PaymentIDByPaymentServiceAndPaymentProductPutBadRequest handles this case with default header values.

Bad Request
*/
type PaymentIDByPaymentServiceAndPaymentProductPutBadRequest struct {
	Payload *models.TppMessage
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{payment-service}/{payment-product}/{paymentId}][%d] paymentIdByPaymentServiceAndPaymentProductPutBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutUnauthorized creates a PaymentIDByPaymentServiceAndPaymentProductPutUnauthorized with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductPutUnauthorized() *PaymentIDByPaymentServiceAndPaymentProductPutUnauthorized {
	return &PaymentIDByPaymentServiceAndPaymentProductPutUnauthorized{}
}

/*PaymentIDByPaymentServiceAndPaymentProductPutUnauthorized handles this case with default header values.

Unauthorized
*/
type PaymentIDByPaymentServiceAndPaymentProductPutUnauthorized struct {
	Payload *models.TppMessage
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /{payment-service}/{payment-product}/{paymentId}][%d] paymentIdByPaymentServiceAndPaymentProductPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutForbidden creates a PaymentIDByPaymentServiceAndPaymentProductPutForbidden with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductPutForbidden() *PaymentIDByPaymentServiceAndPaymentProductPutForbidden {
	return &PaymentIDByPaymentServiceAndPaymentProductPutForbidden{}
}

/*PaymentIDByPaymentServiceAndPaymentProductPutForbidden handles this case with default header values.

Forbidden
*/
type PaymentIDByPaymentServiceAndPaymentProductPutForbidden struct {
	Payload *models.TppMessage
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /{payment-service}/{payment-product}/{paymentId}][%d] paymentIdByPaymentServiceAndPaymentProductPutForbidden  %+v", 403, o.Payload)
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutDefault creates a PaymentIDByPaymentServiceAndPaymentProductPutDefault with default headers values
func NewPaymentIDByPaymentServiceAndPaymentProductPutDefault(code int) *PaymentIDByPaymentServiceAndPaymentProductPutDefault {
	return &PaymentIDByPaymentServiceAndPaymentProductPutDefault{
		_statusCode: code,
	}
}

/*PaymentIDByPaymentServiceAndPaymentProductPutDefault handles this case with default header values.

Internal Server Error
*/
type PaymentIDByPaymentServiceAndPaymentProductPutDefault struct {
	_statusCode int

	Payload *models.TppMessage
}

// Code gets the status code for the payment Id by payment service and payment product put default response
func (o *PaymentIDByPaymentServiceAndPaymentProductPutDefault) Code() int {
	return o._statusCode
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutDefault) Error() string {
	return fmt.Sprintf("[PUT /{payment-service}/{payment-product}/{paymentId}][%d] PaymentIdByPaymentServiceAndPaymentProductPut default  %+v", o._statusCode, o.Payload)
}

func (o *PaymentIDByPaymentServiceAndPaymentProductPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
