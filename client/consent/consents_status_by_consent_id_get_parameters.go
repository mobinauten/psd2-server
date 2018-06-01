// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package consent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewConsentsStatusByConsentIDGetParams creates a new ConsentsStatusByConsentIDGetParams object
// with the default values initialized.
func NewConsentsStatusByConsentIDGetParams() *ConsentsStatusByConsentIDGetParams {
	var ()
	return &ConsentsStatusByConsentIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsentsStatusByConsentIDGetParamsWithTimeout creates a new ConsentsStatusByConsentIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsentsStatusByConsentIDGetParamsWithTimeout(timeout time.Duration) *ConsentsStatusByConsentIDGetParams {
	var ()
	return &ConsentsStatusByConsentIDGetParams{

		timeout: timeout,
	}
}

// NewConsentsStatusByConsentIDGetParamsWithContext creates a new ConsentsStatusByConsentIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsentsStatusByConsentIDGetParamsWithContext(ctx context.Context) *ConsentsStatusByConsentIDGetParams {
	var ()
	return &ConsentsStatusByConsentIDGetParams{

		Context: ctx,
	}
}

// NewConsentsStatusByConsentIDGetParamsWithHTTPClient creates a new ConsentsStatusByConsentIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsentsStatusByConsentIDGetParamsWithHTTPClient(client *http.Client) *ConsentsStatusByConsentIDGetParams {
	var ()
	return &ConsentsStatusByConsentIDGetParams{
		HTTPClient: client,
	}
}

/*ConsentsStatusByConsentIDGetParams contains all the parameters to send to the API endpoint
for the consents status by consent Id get operation typically these are written to a http.Request
*/
type ConsentsStatusByConsentIDGetParams struct {

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
	/*ConsentID
	  ID of the corresponding consent object as returned by an Account Information Consent Request

	*/
	ConsentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithTimeout(timeout time.Duration) *ConsentsStatusByConsentIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithContext(ctx context.Context) *ConsentsStatusByConsentIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithHTTPClient(client *http.Client) *ConsentsStatusByConsentIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDigest adds the digest to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithDigest(digest *string) *ConsentsStatusByConsentIDGetParams {
	o.SetDigest(digest)
	return o
}

// SetDigest adds the digest to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetDigest(digest *string) {
	o.Digest = digest
}

// WithSignature adds the signature to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithSignature(signature *string) *ConsentsStatusByConsentIDGetParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithTPPSignatureCertificate adds the tPPSignatureCertificate to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithTPPSignatureCertificate(tPPSignatureCertificate *string) *ConsentsStatusByConsentIDGetParams {
	o.SetTPPSignatureCertificate(tPPSignatureCertificate)
	return o
}

// SetTPPSignatureCertificate adds the tPPSignatureCertificate to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetTPPSignatureCertificate(tPPSignatureCertificate *string) {
	o.TPPSignatureCertificate = tPPSignatureCertificate
}

// WithXRequestID adds the xRequestID to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithXRequestID(xRequestID strfmt.UUID) *ConsentsStatusByConsentIDGetParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetXRequestID(xRequestID strfmt.UUID) {
	o.XRequestID = xRequestID
}

// WithConsentID adds the consentID to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) WithConsentID(consentID string) *ConsentsStatusByConsentIDGetParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the consents status by consent Id get params
func (o *ConsentsStatusByConsentIDGetParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsentsStatusByConsentIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param consentId
	if err := r.SetPathParam("consentId", o.ConsentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
