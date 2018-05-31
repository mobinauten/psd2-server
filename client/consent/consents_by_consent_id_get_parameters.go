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
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewConsentsByConsentIDGetParams creates a new ConsentsByConsentIDGetParams object
// with the default values initialized.
func NewConsentsByConsentIDGetParams() *ConsentsByConsentIDGetParams {
	var ()
	return &ConsentsByConsentIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsentsByConsentIDGetParamsWithTimeout creates a new ConsentsByConsentIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsentsByConsentIDGetParamsWithTimeout(timeout time.Duration) *ConsentsByConsentIDGetParams {
	var ()
	return &ConsentsByConsentIDGetParams{

		timeout: timeout,
	}
}

// NewConsentsByConsentIDGetParamsWithContext creates a new ConsentsByConsentIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsentsByConsentIDGetParamsWithContext(ctx context.Context) *ConsentsByConsentIDGetParams {
	var ()
	return &ConsentsByConsentIDGetParams{

		Context: ctx,
	}
}

// NewConsentsByConsentIDGetParamsWithHTTPClient creates a new ConsentsByConsentIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsentsByConsentIDGetParamsWithHTTPClient(client *http.Client) *ConsentsByConsentIDGetParams {
	var ()
	return &ConsentsByConsentIDGetParams{
		HTTPClient: client,
	}
}

/*ConsentsByConsentIDGetParams contains all the parameters to send to the API endpoint
for the consents by consent Id get operation typically these are written to a http.Request
*/
type ConsentsByConsentIDGetParams struct {

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

// WithTimeout adds the timeout to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithTimeout(timeout time.Duration) *ConsentsByConsentIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithContext(ctx context.Context) *ConsentsByConsentIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithHTTPClient(client *http.Client) *ConsentsByConsentIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDigest adds the digest to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithDigest(digest *string) *ConsentsByConsentIDGetParams {
	o.SetDigest(digest)
	return o
}

// SetDigest adds the digest to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetDigest(digest *string) {
	o.Digest = digest
}

// WithSignature adds the signature to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithSignature(signature *string) *ConsentsByConsentIDGetParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithTPPSignatureCertificate adds the tPPSignatureCertificate to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithTPPSignatureCertificate(tPPSignatureCertificate *string) *ConsentsByConsentIDGetParams {
	o.SetTPPSignatureCertificate(tPPSignatureCertificate)
	return o
}

// SetTPPSignatureCertificate adds the tPPSignatureCertificate to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetTPPSignatureCertificate(tPPSignatureCertificate *string) {
	o.TPPSignatureCertificate = tPPSignatureCertificate
}

// WithXRequestID adds the xRequestID to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithXRequestID(xRequestID strfmt.UUID) *ConsentsByConsentIDGetParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetXRequestID(xRequestID strfmt.UUID) {
	o.XRequestID = xRequestID
}

// WithConsentID adds the consentID to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) WithConsentID(consentID string) *ConsentsByConsentIDGetParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the consents by consent Id get params
func (o *ConsentsByConsentIDGetParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsentsByConsentIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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