// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openpsd/psd2-server/models"
)

// AccountsGetReader is a Reader for the AccountsGet structure.
type AccountsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAccountsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAccountsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAccountsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAccountsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAccountsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountsGetOK creates a AccountsGetOK with default headers values
func NewAccountsGetOK() *AccountsGetOK {
	return &AccountsGetOK{}
}

/*AccountsGetOK handles this case with default header values.

OK
*/
type AccountsGetOK struct {
	Payload *models.AccountsResponse
}

func (o *AccountsGetOK) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] accountsGetOK  %+v", 200, o.Payload)
}

func (o *AccountsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetBadRequest creates a AccountsGetBadRequest with default headers values
func NewAccountsGetBadRequest() *AccountsGetBadRequest {
	return &AccountsGetBadRequest{}
}

/*AccountsGetBadRequest handles this case with default header values.

Bad Request
*/
type AccountsGetBadRequest struct {
	Payload *models.TppMessage
}

func (o *AccountsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] accountsGetBadRequest  %+v", 400, o.Payload)
}

func (o *AccountsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetUnauthorized creates a AccountsGetUnauthorized with default headers values
func NewAccountsGetUnauthorized() *AccountsGetUnauthorized {
	return &AccountsGetUnauthorized{}
}

/*AccountsGetUnauthorized handles this case with default header values.

Unauthorized
*/
type AccountsGetUnauthorized struct {
	Payload *models.TppMessage
}

func (o *AccountsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] accountsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *AccountsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetForbidden creates a AccountsGetForbidden with default headers values
func NewAccountsGetForbidden() *AccountsGetForbidden {
	return &AccountsGetForbidden{}
}

/*AccountsGetForbidden handles this case with default header values.

Forbidden
*/
type AccountsGetForbidden struct {
	Payload *models.TppMessage
}

func (o *AccountsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] accountsGetForbidden  %+v", 403, o.Payload)
}

func (o *AccountsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetDefault creates a AccountsGetDefault with default headers values
func NewAccountsGetDefault(code int) *AccountsGetDefault {
	return &AccountsGetDefault{
		_statusCode: code,
	}
}

/*AccountsGetDefault handles this case with default header values.

Internal Server Error
*/
type AccountsGetDefault struct {
	_statusCode int

	Payload *models.TppMessage
}

// Code gets the status code for the accounts get default response
func (o *AccountsGetDefault) Code() int {
	return o._statusCode
}

func (o *AccountsGetDefault) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] AccountsGet default  %+v", o._statusCode, o.Payload)
}

func (o *AccountsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
