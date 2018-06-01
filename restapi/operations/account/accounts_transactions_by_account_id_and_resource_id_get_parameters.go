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
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAccountsTransactionsByAccountIDAndResourceIDGetParams creates a new AccountsTransactionsByAccountIDAndResourceIDGetParams object
// no default values defined in spec.
func NewAccountsTransactionsByAccountIDAndResourceIDGetParams() AccountsTransactionsByAccountIDAndResourceIDGetParams {

	return AccountsTransactionsByAccountIDAndResourceIDGetParams{}
}

// AccountsTransactionsByAccountIDAndResourceIDGetParams contains all the bound params for the accounts transactions by account Id and resource Id get operation
// typically these are obtained from a http.Request
//
// swagger:parameters AccountsTransactionsByAccountIdAndResourceIdGet
type AccountsTransactionsByAccountIDAndResourceIDGetParams struct {

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
	/*This identification is denoting the addressed account. The account-id is retrieved by using a "Read Account List" call. The account-id is the "resourceId" attribute of the account structure. Its value is constant at least throughout the lifecycle of a given consent.
	  Required: true
	  In: path
	*/
	AccountID string
	/*This identification is given by the attribute resourceId of the corresponding entry of a transaction list.
	  Required: true
	  Max Length: 35
	  In: path
	*/
	ResourceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAccountsTransactionsByAccountIDAndResourceIDGetParams() beforehand.
func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindConsentID(r.Header[http.CanonicalHeaderKey("Consent-ID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

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

	rAccountID, rhkAccountID, _ := route.Params.GetOK("account-id")
	if err := o.bindAccountID(rAccountID, rhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	rResourceID, rhkResourceID, _ := route.Params.GetOK("resourceId")
	if err := o.bindResourceID(rResourceID, rhkResourceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindConsentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindDigest(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindSignature(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindTPPSignatureCertificate(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.FormatOf("X-Request-ID", "header", "uuid", o.XRequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AccountID = raw

	return nil
}

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) bindResourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ResourceID = raw

	if err := o.validateResourceID(formats); err != nil {
		return err
	}

	return nil
}

func (o *AccountsTransactionsByAccountIDAndResourceIDGetParams) validateResourceID(formats strfmt.Registry) error {

	if err := validate.MaxLength("resourceId", "path", o.ResourceID, 35); err != nil {
		return err
	}

	return nil
}
