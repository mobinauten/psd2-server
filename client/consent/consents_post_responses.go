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

// ConsentsPostReader is a Reader for the ConsentsPost structure.
type ConsentsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsentsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewConsentsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConsentsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConsentsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewConsentsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewConsentsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsentsPostCreated creates a ConsentsPostCreated with default headers values
func NewConsentsPostCreated() *ConsentsPostCreated {
	return &ConsentsPostCreated{}
}

/*ConsentsPostCreated handles this case with default header values.

Consent Request was correctly performed.
*/
type ConsentsPostCreated struct {
	Payload *models.ConsentsResponse
}

func (o *ConsentsPostCreated) Error() string {
	return fmt.Sprintf("[POST /consents][%d] consentsPostCreated  %+v", 201, o.Payload)
}

func (o *ConsentsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsPostBadRequest creates a ConsentsPostBadRequest with default headers values
func NewConsentsPostBadRequest() *ConsentsPostBadRequest {
	return &ConsentsPostBadRequest{}
}

/*ConsentsPostBadRequest handles this case with default header values.

Bad Request
*/
type ConsentsPostBadRequest struct {
	Payload *models.TppMessage
}

func (o *ConsentsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /consents][%d] consentsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ConsentsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsPostUnauthorized creates a ConsentsPostUnauthorized with default headers values
func NewConsentsPostUnauthorized() *ConsentsPostUnauthorized {
	return &ConsentsPostUnauthorized{}
}

/*ConsentsPostUnauthorized handles this case with default header values.

Unauthorized
*/
type ConsentsPostUnauthorized struct {
	Payload *models.TppMessage
}

func (o *ConsentsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /consents][%d] consentsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ConsentsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsPostForbidden creates a ConsentsPostForbidden with default headers values
func NewConsentsPostForbidden() *ConsentsPostForbidden {
	return &ConsentsPostForbidden{}
}

/*ConsentsPostForbidden handles this case with default header values.

Forbidden
*/
type ConsentsPostForbidden struct {
	Payload *models.TppMessage
}

func (o *ConsentsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /consents][%d] consentsPostForbidden  %+v", 403, o.Payload)
}

func (o *ConsentsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsPostDefault creates a ConsentsPostDefault with default headers values
func NewConsentsPostDefault(code int) *ConsentsPostDefault {
	return &ConsentsPostDefault{
		_statusCode: code,
	}
}

/*ConsentsPostDefault handles this case with default header values.

Internal Server Error
*/
type ConsentsPostDefault struct {
	_statusCode int

	Payload *models.TppMessage
}

// Code gets the status code for the consents post default response
func (o *ConsentsPostDefault) Code() int {
	return o._statusCode
}

func (o *ConsentsPostDefault) Error() string {
	return fmt.Sprintf("[POST /consents][%d] ConsentsPost default  %+v", o._statusCode, o.Payload)
}

func (o *ConsentsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
