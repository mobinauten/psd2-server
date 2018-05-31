// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPaymentIDByPaymentServiceAndPaymentProductGetParams creates a new PaymentIDByPaymentServiceAndPaymentProductGetParams object
// no default values defined in spec.
func NewPaymentIDByPaymentServiceAndPaymentProductGetParams() PaymentIDByPaymentServiceAndPaymentProductGetParams {

	return PaymentIDByPaymentServiceAndPaymentProductGetParams{}
}

// PaymentIDByPaymentServiceAndPaymentProductGetParams contains all the bound params for the payment Id by payment service and payment product get operation
// typically these are obtained from a http.Request
//
// swagger:parameters PaymentIdByPaymentServiceAndPaymentProductGet
type PaymentIDByPaymentServiceAndPaymentProductGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Is contained only, if an OAuth2 based authentication was performed in a pre-step or an OAuth2 based SCA was performed in the related consent authorisation.
	  In: header
	*/
	Authorization *string
	/*Shall be contained since "Establish Consent Transaction" was performed via this API before.
	  Required: true
	  In: header
	*/
	ConsentID string
	/*Content-Type
	  Required: true
	  In: header
	*/
	ContentType string
	/*Is contained if and only if the "Signature" element is contained in the header of the request.
	  In: header
	*/
	Digest *string
	/*Might be mandated in the ASPSP's documentation. Only used in a corporate context.
	  In: header
	*/
	PSUCorporateID *string
	/*Might be mandated in the ASPSPs documentation. Only used in a corporate context.
	  In: header
	*/
	PSUCorporateIDType *string
	/*Might be mandated in the ASPSP's documentation, if OAuth is not chosen as Pre-Step.
	  In: header
	*/
	PSUID *string
	/*Type of the PSU-ID, needed in scenarios where PSUs have several PSU-IDs as access possibility.
	  In: header
	*/
	PSUIDType *string
	/*The forwarded IP Address header field consists of the corresponding HTTP request IP Address field between PSU and TPP.
	  Required: true
	  In: header
	*/
	PSUipAddress string
	/*A signature of the request by the TPP on application level. This might be mandated by ASPSP.
	  In: header
	*/
	Signature *string
	/*
	  In: header
	*/
	TPPNokRedirectURI *string
	/*If it equals "true", the TPP prefers a redirect over an embedded SCA approach. If it equals "false", the TPP prefers not to be redirected for SCA. The ASPSP will then choose between the Embedded or the Decoupled SCA approach, depending on the choice of the SCA procedure by the TPP/PSU. If the parameter is not used, the ASPSP will choose the SCA approach to be applied depending on the SCA method chosen by the TPP/PSU.
	  In: header
	*/
	TPPRedirectPreferred *bool
	/*URI of the TPP, where the transaction flow shall be redirected to after a Redirect. Mandatory for the SCA OAuth Approach.
	  In: header
	*/
	TPPRedirectURI *string
	/*The certificate used for signing the request, in base64 encoding. It shall be contained if a signature is used, see above.
	  In: header
	*/
	TPPSignatureCertificate *string
	/*ID of the request, unique to the call, as determined by the initiating party.
	  Required: true
	  In: header
	*/
	XRequestID strfmt.UUID
	/*The addressed payment product, e.g. SCT. The default list of products supported in this standard is, pain.001-sepa-credit-transfers, pain.001-instant-sepa-credit-transfers, pain.001-target-2-payments or pain.001-cross-border-credit-transfers. Further products might be published by the ASPSP within its XS2A documentation.
	  Required: true
	  In: path
	*/
	PaymentProduct string
	/*The payment service to use.
	  Required: true
	  In: path
	*/
	PaymentService string
	/*payment Id
	  Required: true
	  In: path
	*/
	PaymentID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPaymentIDByPaymentServiceAndPaymentProductGetParams() beforehand.
func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindConsentID(r.Header[http.CanonicalHeaderKey("Consent-ID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindContentType(r.Header[http.CanonicalHeaderKey("Content-Type")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindDigest(r.Header[http.CanonicalHeaderKey("Digest")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindPSUCorporateID(r.Header[http.CanonicalHeaderKey("PSU-Corporate-ID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindPSUCorporateIDType(r.Header[http.CanonicalHeaderKey("PSU-Corporate-ID-Type")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindPSUID(r.Header[http.CanonicalHeaderKey("PSU-ID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindPSUIDType(r.Header[http.CanonicalHeaderKey("PSU-ID-Type")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindPSUipAddress(r.Header[http.CanonicalHeaderKey("PSU-IP-Address")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindSignature(r.Header[http.CanonicalHeaderKey("Signature")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindTPPNokRedirectURI(r.Header[http.CanonicalHeaderKey("TPP-Nok-Redirect-URI")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindTPPRedirectPreferred(r.Header[http.CanonicalHeaderKey("TPP-Redirect-Preferred")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindTPPRedirectURI(r.Header[http.CanonicalHeaderKey("TPP-Redirect-URI")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindTPPSignatureCertificate(r.Header[http.CanonicalHeaderKey("TPP-Signature-Certificate")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-ID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rPaymentProduct, rhkPaymentProduct, _ := route.Params.GetOK("payment-product")
	if err := o.bindPaymentProduct(rPaymentProduct, rhkPaymentProduct, route.Formats); err != nil {
		res = append(res, err)
	}

	rPaymentService, rhkPaymentService, _ := route.Params.GetOK("payment-service")
	if err := o.bindPaymentService(rPaymentService, rhkPaymentService, route.Formats); err != nil {
		res = append(res, err)
	}

	rPaymentID, rhkPaymentID, _ := route.Params.GetOK("paymentId")
	if err := o.bindPaymentID(rPaymentID, rhkPaymentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Authorization = &raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindConsentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Consent-ID", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Consent-ID", "header", raw); err != nil {
		return err
	}

	o.ConsentID = raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindContentType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Content-Type", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Content-Type", "header", raw); err != nil {
		return err
	}

	o.ContentType = raw

	if err := o.validateContentType(formats); err != nil {
		return err
	}

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) validateContentType(formats strfmt.Registry) error {

	if err := validate.Enum("Content-Type", "header", o.ContentType, []interface{}{"application/json", "application/xml", "application/text"}); err != nil {
		return err
	}

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindDigest(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPSUCorporateID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.PSUCorporateID = &raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPSUCorporateIDType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.PSUCorporateIDType = &raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPSUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.PSUID = &raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPSUIDType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.PSUIDType = &raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPSUipAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("PSU-IP-Address", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("PSU-IP-Address", "header", raw); err != nil {
		return err
	}

	o.PSUIPAddress = raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindSignature(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindTPPNokRedirectURI(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TPPNokRedirectURI = &raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindTPPRedirectPreferred(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("TPP-Redirect-Preferred", "header", "bool", raw)
	}
	o.TPPRedirectPreferred = &value

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindTPPRedirectURI(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TPPRedirectURI = &raw

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindTPPSignatureCertificate(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.FormatOf("X-Request-ID", "header", "uuid", o.XRequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPaymentProduct(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PaymentProduct = raw

	if err := o.validatePaymentProduct(formats); err != nil {
		return err
	}

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) validatePaymentProduct(formats strfmt.Registry) error {

	if err := validate.Enum("payment-product", "path", o.PaymentProduct, []interface{}{"sepa-credit-transfers", "instant-sepa-credit-transfers", "target-2-payments", "cross-border-credit-transfers"}); err != nil {
		return err
	}

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPaymentService(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PaymentService = raw

	if err := o.validatePaymentService(formats); err != nil {
		return err
	}

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) validatePaymentService(formats strfmt.Registry) error {

	if err := validate.Enum("payment-service", "path", o.PaymentService, []interface{}{"payments", "bulk-payments", "periodic-payments"}); err != nil {
		return err
	}

	return nil
}

func (o *PaymentIDByPaymentServiceAndPaymentProductGetParams) bindPaymentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PaymentID = raw

	return nil
}
