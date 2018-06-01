// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AccountsBalancesByAccountIDGet balances for account

Reads the balance list for a given account.
*/
func (a *Client) AccountsBalancesByAccountIDGet(params *AccountsBalancesByAccountIDGetParams) (*AccountsBalancesByAccountIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsBalancesByAccountIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AccountsBalancesByAccountIdGet",
		Method:             "GET",
		PathPattern:        "/accounts/{account-id}/balances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsBalancesByAccountIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AccountsBalancesByAccountIDGetOK), nil

}

/*
AccountsByAccountIDGet accounts details

Reads a list of bank accounts, with balances where required. It is assumed that a consent of the PSU to this access is already given and stored on the ASPSP system. The addressed list of accounts depends then on the PSU ID and the stored consent addressed by consentId, respectively the OAuth2 access token.
*/
func (a *Client) AccountsByAccountIDGet(params *AccountsByAccountIDGetParams) (*AccountsByAccountIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsByAccountIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AccountsByAccountIdGet",
		Method:             "GET",
		PathPattern:        "/accounts/{account-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsByAccountIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AccountsByAccountIDGetOK), nil

}

/*
AccountsGet lists of accounts

Reads a list of bank accounts, with balances where required. It is assumed that a consent of the PSU to this access is already given and stored on the ASPSP system. The addressed list of accounts depends then on the PSU ID and the stored consent addressed by consentId, respectively the OAuth2 access token.
*/
func (a *Client) AccountsGet(params *AccountsGetParams) (*AccountsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AccountsGet",
		Method:             "GET",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AccountsGetOK), nil

}

/*
AccountsTransactionsByAccountIDAndResourceIDGet transactions details for transaction of an account

Reads transaction details of an addressed transaction.
*/
func (a *Client) AccountsTransactionsByAccountIDAndResourceIDGet(params *AccountsTransactionsByAccountIDAndResourceIDGetParams) (*AccountsTransactionsByAccountIDAndResourceIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsTransactionsByAccountIDAndResourceIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AccountsTransactionsByAccountIdAndResourceIdGet",
		Method:             "GET",
		PathPattern:        "/accounts/{account-id}/transactions/{resourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsTransactionsByAccountIDAndResourceIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AccountsTransactionsByAccountIDAndResourceIDGetOK), nil

}

/*
AccountsTransactionsByAccountIDGet transactions for account

Reads a transaction list For a given account, additional parameters are e.g. the attributes “date_from” and “date_to”. If the attribute “withbalance” is used, the ASPSP will add balances to the transaction list. The latter might be provided by the ASPSP anyhow, if transaction lists without balances are not supported.
*/
func (a *Client) AccountsTransactionsByAccountIDGet(params *AccountsTransactionsByAccountIDGetParams) (*AccountsTransactionsByAccountIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsTransactionsByAccountIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AccountsTransactionsByAccountIdGet",
		Method:             "GET",
		PathPattern:        "/accounts/{account-id}/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsTransactionsByAccountIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AccountsTransactionsByAccountIDGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
