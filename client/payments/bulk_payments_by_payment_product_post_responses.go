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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openpsd/psd2-server/models"
)

// BulkPaymentsByPaymentProductPostReader is a Reader for the BulkPaymentsByPaymentProductPost structure.
type BulkPaymentsByPaymentProductPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkPaymentsByPaymentProductPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBulkPaymentsByPaymentProductPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBulkPaymentsByPaymentProductPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBulkPaymentsByPaymentProductPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewBulkPaymentsByPaymentProductPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewBulkPaymentsByPaymentProductPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkPaymentsByPaymentProductPostOK creates a BulkPaymentsByPaymentProductPostOK with default headers values
func NewBulkPaymentsByPaymentProductPostOK() *BulkPaymentsByPaymentProductPostOK {
	return &BulkPaymentsByPaymentProductPostOK{}
}

/*BulkPaymentsByPaymentProductPostOK handles this case with default header values.

OK
*/
type BulkPaymentsByPaymentProductPostOK struct {
	Payload *models.BulkPaymentsResponse
}

func (o *BulkPaymentsByPaymentProductPostOK) Error() string {
	return fmt.Sprintf("[POST /bulk-payments/{payment-product}][%d] bulkPaymentsByPaymentProductPostOK  %+v", 200, o.Payload)
}

func (o *BulkPaymentsByPaymentProductPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkPaymentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkPaymentsByPaymentProductPostBadRequest creates a BulkPaymentsByPaymentProductPostBadRequest with default headers values
func NewBulkPaymentsByPaymentProductPostBadRequest() *BulkPaymentsByPaymentProductPostBadRequest {
	return &BulkPaymentsByPaymentProductPostBadRequest{}
}

/*BulkPaymentsByPaymentProductPostBadRequest handles this case with default header values.

Bad Request
*/
type BulkPaymentsByPaymentProductPostBadRequest struct {
	Payload *models.TppMessage
}

func (o *BulkPaymentsByPaymentProductPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /bulk-payments/{payment-product}][%d] bulkPaymentsByPaymentProductPostBadRequest  %+v", 400, o.Payload)
}

func (o *BulkPaymentsByPaymentProductPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkPaymentsByPaymentProductPostUnauthorized creates a BulkPaymentsByPaymentProductPostUnauthorized with default headers values
func NewBulkPaymentsByPaymentProductPostUnauthorized() *BulkPaymentsByPaymentProductPostUnauthorized {
	return &BulkPaymentsByPaymentProductPostUnauthorized{}
}

/*BulkPaymentsByPaymentProductPostUnauthorized handles this case with default header values.

Unauthorized
*/
type BulkPaymentsByPaymentProductPostUnauthorized struct {
	Payload *models.TppMessage
}

func (o *BulkPaymentsByPaymentProductPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bulk-payments/{payment-product}][%d] bulkPaymentsByPaymentProductPostUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkPaymentsByPaymentProductPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkPaymentsByPaymentProductPostForbidden creates a BulkPaymentsByPaymentProductPostForbidden with default headers values
func NewBulkPaymentsByPaymentProductPostForbidden() *BulkPaymentsByPaymentProductPostForbidden {
	return &BulkPaymentsByPaymentProductPostForbidden{}
}

/*BulkPaymentsByPaymentProductPostForbidden handles this case with default header values.

Forbidden
*/
type BulkPaymentsByPaymentProductPostForbidden struct {
	Payload *models.TppMessage
}

func (o *BulkPaymentsByPaymentProductPostForbidden) Error() string {
	return fmt.Sprintf("[POST /bulk-payments/{payment-product}][%d] bulkPaymentsByPaymentProductPostForbidden  %+v", 403, o.Payload)
}

func (o *BulkPaymentsByPaymentProductPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkPaymentsByPaymentProductPostDefault creates a BulkPaymentsByPaymentProductPostDefault with default headers values
func NewBulkPaymentsByPaymentProductPostDefault(code int) *BulkPaymentsByPaymentProductPostDefault {
	return &BulkPaymentsByPaymentProductPostDefault{
		_statusCode: code,
	}
}

/*BulkPaymentsByPaymentProductPostDefault handles this case with default header values.

Internal Server Error
*/
type BulkPaymentsByPaymentProductPostDefault struct {
	_statusCode int

	Payload *models.TppMessage
}

// Code gets the status code for the bulk payments by payment product post default response
func (o *BulkPaymentsByPaymentProductPostDefault) Code() int {
	return o._statusCode
}

func (o *BulkPaymentsByPaymentProductPostDefault) Error() string {
	return fmt.Sprintf("[POST /bulk-payments/{payment-product}][%d] BulkPaymentsByPaymentProductPost default  %+v", o._statusCode, o.Payload)
}

func (o *BulkPaymentsByPaymentProductPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
