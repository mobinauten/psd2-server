// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AccountsTransactionsByAccountIDGetHandlerFunc turns a function with the right signature into a accounts transactions by account Id get handler
type AccountsTransactionsByAccountIDGetHandlerFunc func(AccountsTransactionsByAccountIDGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AccountsTransactionsByAccountIDGetHandlerFunc) Handle(params AccountsTransactionsByAccountIDGetParams) middleware.Responder {
	return fn(params)
}

// AccountsTransactionsByAccountIDGetHandler interface for that can handle valid accounts transactions by account Id get params
type AccountsTransactionsByAccountIDGetHandler interface {
	Handle(AccountsTransactionsByAccountIDGetParams) middleware.Responder
}

// NewAccountsTransactionsByAccountIDGet creates a new http.Handler for the accounts transactions by account Id get operation
func NewAccountsTransactionsByAccountIDGet(ctx *middleware.Context, handler AccountsTransactionsByAccountIDGetHandler) *AccountsTransactionsByAccountIDGet {
	return &AccountsTransactionsByAccountIDGet{Context: ctx, Handler: handler}
}

/*AccountsTransactionsByAccountIDGet swagger:route GET /accounts/{account-id}/transactions account accountsTransactionsByAccountIdGet

transactions for account

Reads a transaction list For a given account, additional parameters are e.g. the attributes “date_from” and “date_to”. If the attribute “withbalance” is used, the ASPSP will add balances to the transaction list. The latter might be provided by the ASPSP anyhow, if transaction lists without balances are not supported.

*/
type AccountsTransactionsByAccountIDGet struct {
	Context *middleware.Context
	Handler AccountsTransactionsByAccountIDGetHandler
}

func (o *AccountsTransactionsByAccountIDGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAccountsTransactionsByAccountIDGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
