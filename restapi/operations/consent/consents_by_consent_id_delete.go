// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package consent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ConsentsByConsentIDDeleteHandlerFunc turns a function with the right signature into a consents by consent Id delete handler
type ConsentsByConsentIDDeleteHandlerFunc func(ConsentsByConsentIDDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConsentsByConsentIDDeleteHandlerFunc) Handle(params ConsentsByConsentIDDeleteParams) middleware.Responder {
	return fn(params)
}

// ConsentsByConsentIDDeleteHandler interface for that can handle valid consents by consent Id delete params
type ConsentsByConsentIDDeleteHandler interface {
	Handle(ConsentsByConsentIDDeleteParams) middleware.Responder
}

// NewConsentsByConsentIDDelete creates a new http.Handler for the consents by consent Id delete operation
func NewConsentsByConsentIDDelete(ctx *middleware.Context, handler ConsentsByConsentIDDeleteHandler) *ConsentsByConsentIDDelete {
	return &ConsentsByConsentIDDelete{Context: ctx, Handler: handler}
}

/*ConsentsByConsentIDDelete swagger:route DELETE /consents/{consentId} Consent consentsByConsentIdDelete

delete consent

Deletes a created consent with ID.

*/
type ConsentsByConsentIDDelete struct {
	Context *middleware.Context
	Handler ConsentsByConsentIDDeleteHandler
}

func (o *ConsentsByConsentIDDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConsentsByConsentIDDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
