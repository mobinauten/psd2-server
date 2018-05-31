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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openpsd/psd2-server/models"
)

// ConsentsByConsentIDDeleteReader is a Reader for the ConsentsByConsentIDDelete structure.
type ConsentsByConsentIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsentsByConsentIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewConsentsByConsentIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConsentsByConsentIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConsentsByConsentIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewConsentsByConsentIDDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewConsentsByConsentIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsentsByConsentIDDeleteNoContent creates a ConsentsByConsentIDDeleteNoContent with default headers values
func NewConsentsByConsentIDDeleteNoContent() *ConsentsByConsentIDDeleteNoContent {
	return &ConsentsByConsentIDDeleteNoContent{}
}

/*ConsentsByConsentIDDeleteNoContent handles this case with default header values.

Consent resource was successfully deleted.
*/
type ConsentsByConsentIDDeleteNoContent struct {
}

func (o *ConsentsByConsentIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsByConsentIdDeleteNoContent ", 204)
}

func (o *ConsentsByConsentIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsentsByConsentIDDeleteBadRequest creates a ConsentsByConsentIDDeleteBadRequest with default headers values
func NewConsentsByConsentIDDeleteBadRequest() *ConsentsByConsentIDDeleteBadRequest {
	return &ConsentsByConsentIDDeleteBadRequest{}
}

/*ConsentsByConsentIDDeleteBadRequest handles this case with default header values.

Bad Request
*/
type ConsentsByConsentIDDeleteBadRequest struct {
	Payload *models.TppMessage
}

func (o *ConsentsByConsentIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsByConsentIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ConsentsByConsentIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsByConsentIDDeleteUnauthorized creates a ConsentsByConsentIDDeleteUnauthorized with default headers values
func NewConsentsByConsentIDDeleteUnauthorized() *ConsentsByConsentIDDeleteUnauthorized {
	return &ConsentsByConsentIDDeleteUnauthorized{}
}

/*ConsentsByConsentIDDeleteUnauthorized handles this case with default header values.

Unauthorized
*/
type ConsentsByConsentIDDeleteUnauthorized struct {
	Payload *models.TppMessage
}

func (o *ConsentsByConsentIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsByConsentIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ConsentsByConsentIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsByConsentIDDeleteForbidden creates a ConsentsByConsentIDDeleteForbidden with default headers values
func NewConsentsByConsentIDDeleteForbidden() *ConsentsByConsentIDDeleteForbidden {
	return &ConsentsByConsentIDDeleteForbidden{}
}

/*ConsentsByConsentIDDeleteForbidden handles this case with default header values.

Forbidden
*/
type ConsentsByConsentIDDeleteForbidden struct {
	Payload *models.TppMessage
}

func (o *ConsentsByConsentIDDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsByConsentIdDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ConsentsByConsentIDDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsByConsentIDDeleteDefault creates a ConsentsByConsentIDDeleteDefault with default headers values
func NewConsentsByConsentIDDeleteDefault(code int) *ConsentsByConsentIDDeleteDefault {
	return &ConsentsByConsentIDDeleteDefault{
		_statusCode: code,
	}
}

/*ConsentsByConsentIDDeleteDefault handles this case with default header values.

Internal Server Error
*/
type ConsentsByConsentIDDeleteDefault struct {
	_statusCode int

	Payload *models.TppMessage
}

// Code gets the status code for the consents by consent Id delete default response
func (o *ConsentsByConsentIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ConsentsByConsentIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] ConsentsByConsentIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *ConsentsByConsentIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
