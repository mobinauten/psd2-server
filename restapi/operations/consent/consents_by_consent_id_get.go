// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package consent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ConsentsByConsentIDGetHandlerFunc turns a function with the right signature into a consents by consent Id get handler
type ConsentsByConsentIDGetHandlerFunc func(ConsentsByConsentIDGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConsentsByConsentIDGetHandlerFunc) Handle(params ConsentsByConsentIDGetParams) middleware.Responder {
	return fn(params)
}

// ConsentsByConsentIDGetHandler interface for that can handle valid consents by consent Id get params
type ConsentsByConsentIDGetHandler interface {
	Handle(ConsentsByConsentIDGetParams) middleware.Responder
}

// NewConsentsByConsentIDGet creates a new http.Handler for the consents by consent Id get operation
func NewConsentsByConsentIDGet(ctx *middleware.Context, handler ConsentsByConsentIDGetHandler) *ConsentsByConsentIDGet {
	return &ConsentsByConsentIDGet{Context: ctx, Handler: handler}
}

/*ConsentsByConsentIDGet swagger:route GET /consents/{consentId} Consent consentsByConsentIdGet

get consent details

Reads the exact definition of the given consent resource {consentId}

*/
type ConsentsByConsentIDGet struct {
	Context *middleware.Context
	Handler ConsentsByConsentIDGetHandler
}

func (o *ConsentsByConsentIDGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConsentsByConsentIDGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
