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
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAccountsByAccountIDGetParams creates a new AccountsByAccountIDGetParams object
// with the default values initialized.
func NewAccountsByAccountIDGetParams() *AccountsByAccountIDGetParams {
	var ()
	return &AccountsByAccountIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountsByAccountIDGetParamsWithTimeout creates a new AccountsByAccountIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountsByAccountIDGetParamsWithTimeout(timeout time.Duration) *AccountsByAccountIDGetParams {
	var ()
	return &AccountsByAccountIDGetParams{

		timeout: timeout,
	}
}

// NewAccountsByAccountIDGetParamsWithContext creates a new AccountsByAccountIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountsByAccountIDGetParamsWithContext(ctx context.Context) *AccountsByAccountIDGetParams {
	var ()
	return &AccountsByAccountIDGetParams{

		Context: ctx,
	}
}

// NewAccountsByAccountIDGetParamsWithHTTPClient creates a new AccountsByAccountIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountsByAccountIDGetParamsWithHTTPClient(client *http.Client) *AccountsByAccountIDGetParams {
	var ()
	return &AccountsByAccountIDGetParams{
		HTTPClient: client,
	}
}

/*AccountsByAccountIDGetParams contains all the parameters to send to the API endpoint
for the accounts by account Id get operation typically these are written to a http.Request
*/
type AccountsByAccountIDGetParams struct {

	/*Authorization
	  Is contained only, if an OAuth2 based authentication was performed in a pre-step or an OAuth2 based SCA was performed in the related consent authorisation.

	*/
	Authorization *string
	/*ConsentID
	  Shall be contained since "Establish Consent Transaction" was performed via this API before.

	*/
	ConsentID string
	/*Digest
	  Is contained if and only if the "Signature" element is contained in the header of the request.

	*/
	Digest *string
	/*Signature
	  A signature of the request by the TPP on application level. This might be mandated by ASPSP.

	*/
	Signature *string
	/*TPPSignatureCertificate
	  The certificate used for signing the request, in base64 encoding. It shall be contained if a signature is used, see above.

	*/
	TPPSignatureCertificate *string
	/*XRequestID
	  ID of the request, unique to the call, as determined by the initiating party.

	*/
	XRequestID strfmt.UUID
	/*AccountID
	  This identification is denoting the addressed account. The account-id is retrieved by using a "Read Account List" call. The account-id is the "resourceId" attribute of the account structure. Its value is constant at least throughout the lifecycle of a given consent.

	*/
	AccountID string
	/*WithBalance
	  If contained, this function reads the list of accessible payment accounts including the booking balance, if granted by the PSU in the related consent and available by the ASPSP. This parameter might be ignored by the ASPSP.

	*/
	WithBalance *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithTimeout(timeout time.Duration) *AccountsByAccountIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithContext(ctx context.Context) *AccountsByAccountIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithHTTPClient(client *http.Client) *AccountsByAccountIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithAuthorization(authorization *string) *AccountsByAccountIDGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithConsentID adds the consentID to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithConsentID(consentID string) *AccountsByAccountIDGetParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WithDigest adds the digest to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithDigest(digest *string) *AccountsByAccountIDGetParams {
	o.SetDigest(digest)
	return o
}

// SetDigest adds the digest to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetDigest(digest *string) {
	o.Digest = digest
}

// WithSignature adds the signature to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithSignature(signature *string) *AccountsByAccountIDGetParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithTPPSignatureCertificate adds the tPPSignatureCertificate to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithTPPSignatureCertificate(tPPSignatureCertificate *string) *AccountsByAccountIDGetParams {
	o.SetTPPSignatureCertificate(tPPSignatureCertificate)
	return o
}

// SetTPPSignatureCertificate adds the tPPSignatureCertificate to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetTPPSignatureCertificate(tPPSignatureCertificate *string) {
	o.TPPSignatureCertificate = tPPSignatureCertificate
}

// WithXRequestID adds the xRequestID to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithXRequestID(xRequestID strfmt.UUID) *AccountsByAccountIDGetParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetXRequestID(xRequestID strfmt.UUID) {
	o.XRequestID = xRequestID
}

// WithAccountID adds the accountID to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithAccountID(accountID string) *AccountsByAccountIDGetParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithWithBalance adds the withBalance to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) WithWithBalance(withBalance *bool) *AccountsByAccountIDGetParams {
	o.SetWithBalance(withBalance)
	return o
}

// SetWithBalance adds the withBalance to the accounts by account Id get params
func (o *AccountsByAccountIDGetParams) SetWithBalance(withBalance *bool) {
	o.WithBalance = withBalance
}

// WriteToRequest writes these params to a swagger request
func (o *AccountsByAccountIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// header param Consent-ID
	if err := r.SetHeaderParam("Consent-ID", o.ConsentID); err != nil {
		return err
	}

	if o.Digest != nil {

		// header param Digest
		if err := r.SetHeaderParam("Digest", *o.Digest); err != nil {
			return err
		}

	}

	if o.Signature != nil {

		// header param Signature
		if err := r.SetHeaderParam("Signature", *o.Signature); err != nil {
			return err
		}

	}

	if o.TPPSignatureCertificate != nil {

		// header param TPP-Signature-Certificate
		if err := r.SetHeaderParam("TPP-Signature-Certificate", *o.TPPSignatureCertificate); err != nil {
			return err
		}

	}

	// header param X-Request-ID
	if err := r.SetHeaderParam("X-Request-ID", o.XRequestID.String()); err != nil {
		return err
	}

	// path param account-id
	if err := r.SetPathParam("account-id", o.AccountID); err != nil {
		return err
	}

	if o.WithBalance != nil {

		// query param withBalance
		var qrWithBalance bool
		if o.WithBalance != nil {
			qrWithBalance = *o.WithBalance
		}
		qWithBalance := swag.FormatBool(qrWithBalance)
		if qWithBalance != "" {
			if err := r.SetQueryParam("withBalance", qWithBalance); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
