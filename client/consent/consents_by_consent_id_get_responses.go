// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package consent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openpsd/psd2-server/models"
)

// ConsentsByConsentIDGetReader is a Reader for the ConsentsByConsentIDGet structure.
type ConsentsByConsentIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsentsByConsentIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConsentsByConsentIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConsentsByConsentIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConsentsByConsentIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewConsentsByConsentIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewConsentsByConsentIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsentsByConsentIDGetOK creates a ConsentsByConsentIDGetOK with default headers values
func NewConsentsByConsentIDGetOK() *ConsentsByConsentIDGetOK {
	return &ConsentsByConsentIDGetOK{}
}

/*ConsentsByConsentIDGetOK handles this case with default header values.

OK
*/
type ConsentsByConsentIDGetOK struct {
	Payload *models.ConsentsResponse125
}

func (o *ConsentsByConsentIDGetOK) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsByConsentIdGetOK  %+v", 200, o.Payload)
}

func (o *ConsentsByConsentIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentsResponse125)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsByConsentIDGetBadRequest creates a ConsentsByConsentIDGetBadRequest with default headers values
func NewConsentsByConsentIDGetBadRequest() *ConsentsByConsentIDGetBadRequest {
	return &ConsentsByConsentIDGetBadRequest{}
}

/*ConsentsByConsentIDGetBadRequest handles this case with default header values.

Bad Request
*/
type ConsentsByConsentIDGetBadRequest struct {
	Payload *models.TppMessage
}

func (o *ConsentsByConsentIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsByConsentIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *ConsentsByConsentIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsByConsentIDGetUnauthorized creates a ConsentsByConsentIDGetUnauthorized with default headers values
func NewConsentsByConsentIDGetUnauthorized() *ConsentsByConsentIDGetUnauthorized {
	return &ConsentsByConsentIDGetUnauthorized{}
}

/*ConsentsByConsentIDGetUnauthorized handles this case with default header values.

Unauthorized
*/
type ConsentsByConsentIDGetUnauthorized struct {
	Payload *models.TppMessage
}

func (o *ConsentsByConsentIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsByConsentIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ConsentsByConsentIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsByConsentIDGetForbidden creates a ConsentsByConsentIDGetForbidden with default headers values
func NewConsentsByConsentIDGetForbidden() *ConsentsByConsentIDGetForbidden {
	return &ConsentsByConsentIDGetForbidden{}
}

/*ConsentsByConsentIDGetForbidden handles this case with default header values.

Forbidden
*/
type ConsentsByConsentIDGetForbidden struct {
	Payload *models.TppMessage
}

func (o *ConsentsByConsentIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsByConsentIdGetForbidden  %+v", 403, o.Payload)
}

func (o *ConsentsByConsentIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsByConsentIDGetDefault creates a ConsentsByConsentIDGetDefault with default headers values
func NewConsentsByConsentIDGetDefault(code int) *ConsentsByConsentIDGetDefault {
	return &ConsentsByConsentIDGetDefault{
		_statusCode: code,
	}
}

/*ConsentsByConsentIDGetDefault handles this case with default header values.

Internal Server Error
*/
type ConsentsByConsentIDGetDefault struct {
	_statusCode int

	Payload *models.TppMessage
}

// Code gets the status code for the consents by consent Id get default response
func (o *ConsentsByConsentIDGetDefault) Code() int {
	return o._statusCode
}

func (o *ConsentsByConsentIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] ConsentsByConsentIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *ConsentsByConsentIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
