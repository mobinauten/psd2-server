// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BulkPaymentsByPaymentProductPostHandlerFunc turns a function with the right signature into a bulk payments by payment product post handler
type BulkPaymentsByPaymentProductPostHandlerFunc func(BulkPaymentsByPaymentProductPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BulkPaymentsByPaymentProductPostHandlerFunc) Handle(params BulkPaymentsByPaymentProductPostParams) middleware.Responder {
	return fn(params)
}

// BulkPaymentsByPaymentProductPostHandler interface for that can handle valid bulk payments by payment product post params
type BulkPaymentsByPaymentProductPostHandler interface {
	Handle(BulkPaymentsByPaymentProductPostParams) middleware.Responder
}

// NewBulkPaymentsByPaymentProductPost creates a new http.Handler for the bulk payments by payment product post operation
func NewBulkPaymentsByPaymentProductPost(ctx *middleware.Context, handler BulkPaymentsByPaymentProductPostHandler) *BulkPaymentsByPaymentProductPost {
	return &BulkPaymentsByPaymentProductPost{Context: ctx, Handler: handler}
}

/*BulkPaymentsByPaymentProductPost swagger:route POST /bulk-payments/{payment-product} payments bulkPaymentsByPaymentProductPost

create a bulk payment

Creates a payment initiation resource addressable under {paymentId} with all data relevant for the corresponding payment product. This is the first step in the API to initiate the related payment

*/
type BulkPaymentsByPaymentProductPost struct {
	Context *middleware.Context
	Handler BulkPaymentsByPaymentProductPostHandler
}

func (o *BulkPaymentsByPaymentProductPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBulkPaymentsByPaymentProductPostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}