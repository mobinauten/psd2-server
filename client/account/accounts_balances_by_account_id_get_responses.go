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

// AccountsBalancesByAccountIDGetReader is a Reader for the AccountsBalancesByAccountIDGet structure.
type AccountsBalancesByAccountIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountsBalancesByAccountIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAccountsBalancesByAccountIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAccountsBalancesByAccountIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAccountsBalancesByAccountIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAccountsBalancesByAccountIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAccountsBalancesByAccountIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountsBalancesByAccountIDGetOK creates a AccountsBalancesByAccountIDGetOK with default headers values
func NewAccountsBalancesByAccountIDGetOK() *AccountsBalancesByAccountIDGetOK {
	return &AccountsBalancesByAccountIDGetOK{}
}

/*AccountsBalancesByAccountIDGetOK handles this case with default header values.

OK
*/
type AccountsBalancesByAccountIDGetOK struct {
	Payload *models.AccountsBalancesResponse
}

func (o *AccountsBalancesByAccountIDGetOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account-id}/balances][%d] accountsBalancesByAccountIdGetOK  %+v", 200, o.Payload)
}

func (o *AccountsBalancesByAccountIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountsBalancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsBalancesByAccountIDGetBadRequest creates a AccountsBalancesByAccountIDGetBadRequest with default headers values
func NewAccountsBalancesByAccountIDGetBadRequest() *AccountsBalancesByAccountIDGetBadRequest {
	return &AccountsBalancesByAccountIDGetBadRequest{}
}

/*AccountsBalancesByAccountIDGetBadRequest handles this case with default header values.

Bad Request
*/
type AccountsBalancesByAccountIDGetBadRequest struct {
	Payload *models.TppMessage
}

func (o *AccountsBalancesByAccountIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{account-id}/balances][%d] accountsBalancesByAccountIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *AccountsBalancesByAccountIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsBalancesByAccountIDGetUnauthorized creates a AccountsBalancesByAccountIDGetUnauthorized with default headers values
func NewAccountsBalancesByAccountIDGetUnauthorized() *AccountsBalancesByAccountIDGetUnauthorized {
	return &AccountsBalancesByAccountIDGetUnauthorized{}
}

/*AccountsBalancesByAccountIDGetUnauthorized handles this case with default header values.

Unauthorized
*/
type AccountsBalancesByAccountIDGetUnauthorized struct {
	Payload *models.TppMessage
}

func (o *AccountsBalancesByAccountIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{account-id}/balances][%d] accountsBalancesByAccountIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *AccountsBalancesByAccountIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsBalancesByAccountIDGetForbidden creates a AccountsBalancesByAccountIDGetForbidden with default headers values
func NewAccountsBalancesByAccountIDGetForbidden() *AccountsBalancesByAccountIDGetForbidden {
	return &AccountsBalancesByAccountIDGetForbidden{}
}

/*AccountsBalancesByAccountIDGetForbidden handles this case with default header values.

Forbidden
*/
type AccountsBalancesByAccountIDGetForbidden struct {
	Payload *models.TppMessage
}

func (o *AccountsBalancesByAccountIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{account-id}/balances][%d] accountsBalancesByAccountIdGetForbidden  %+v", 403, o.Payload)
}

func (o *AccountsBalancesByAccountIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsBalancesByAccountIDGetDefault creates a AccountsBalancesByAccountIDGetDefault with default headers values
func NewAccountsBalancesByAccountIDGetDefault(code int) *AccountsBalancesByAccountIDGetDefault {
	return &AccountsBalancesByAccountIDGetDefault{
		_statusCode: code,
	}
}

/*AccountsBalancesByAccountIDGetDefault handles this case with default header values.

Internal Server Error
*/
type AccountsBalancesByAccountIDGetDefault struct {
	_statusCode int

	Payload *models.TppMessage
}

// Code gets the status code for the accounts balances by account Id get default response
func (o *AccountsBalancesByAccountIDGetDefault) Code() int {
	return o._statusCode
}

func (o *AccountsBalancesByAccountIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account-id}/balances][%d] AccountsBalancesByAccountIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *AccountsBalancesByAccountIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TppMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
