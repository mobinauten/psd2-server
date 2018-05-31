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
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewConsentsByConsentIDGetParams creates a new ConsentsByConsentIDGetParams object
// no default values defined in spec.
func NewConsentsByConsentIDGetParams() ConsentsByConsentIDGetParams {

	return ConsentsByConsentIDGetParams{}
}

// ConsentsByConsentIDGetParams contains all the bound params for the consents by consent Id get operation
// typically these are obtained from a http.Request
//
// swagger:parameters ConsentsByConsentIdGet
type ConsentsByConsentIDGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Is contained if and only if the "Signature" element is contained in the header of the request.
	  In: header
	*/
	Digest *string
	/*A signature of the request by the TPP on application level. This might be mandated by ASPSP.
	  In: header
	*/
	Signature *string
	/*The certificate used for signing the request, in base64 encoding. It shall be contained if a signature is used, see above.
	  In: header
	*/
	TPPSignatureCertificate *string
	/*ID of the request, unique to the call, as determined by the initiating party.
	  Required: true
	  In: header
	*/
	XRequestID strfmt.UUID
	/*ID of the corresponding consent object as returned by an Account Information Consent Request
	  Required: true
	  In: path
	*/
	ConsentID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewConsentsByConsentIDGetParams() beforehand.
func (o *ConsentsByConsentIDGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindDigest(r.Header[http.CanonicalHeaderKey("Digest")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindSignature(r.Header[http.CanonicalHeaderKey("Signature")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindTPPSignatureCertificate(r.Header[http.CanonicalHeaderKey("TPP-Signature-Certificate")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-ID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rConsentID, rhkConsentID, _ := route.Params.GetOK("consentId")
	if err := o.bindConsentID(rConsentID, rhkConsentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConsentsByConsentIDGetParams) bindDigest(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Digest = &raw

	return nil
}

func (o *ConsentsByConsentIDGetParams) bindSignature(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Signature = &raw

	return nil
}

func (o *ConsentsByConsentIDGetParams) bindTPPSignatureCertificate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TPPSignatureCertificate = &raw

	return nil
}

func (o *ConsentsByConsentIDGetParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Request-ID", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Request-ID", "header", raw); err != nil {
		return err
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("X-Request-ID", "header", "strfmt.UUID", raw)
	}
	o.XRequestID = *(value.(*strfmt.UUID))

	if err := o.validateXRequestID(formats); err != nil {
		return err
	}

	return nil
}

func (o *ConsentsByConsentIDGetParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.FormatOf("X-Request-ID", "header", "uuid", o.XRequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ConsentsByConsentIDGetParams) bindConsentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ConsentID = raw

	return nil
}