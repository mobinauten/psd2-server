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

	models "github.com/openpsd/psd2-server/models"
)

// NewConsentsByConsentIDPutParams creates a new ConsentsByConsentIDPutParams object
// with the default values initialized.
func NewConsentsByConsentIDPutParams() *ConsentsByConsentIDPutParams {
	var ()
	return &ConsentsByConsentIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsentsByConsentIDPutParamsWithTimeout creates a new ConsentsByConsentIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsentsByConsentIDPutParamsWithTimeout(timeout time.Duration) *ConsentsByConsentIDPutParams {
	var ()
	return &ConsentsByConsentIDPutParams{

		timeout: timeout,
	}
}

// NewConsentsByConsentIDPutParamsWithContext creates a new ConsentsByConsentIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsentsByConsentIDPutParamsWithContext(ctx context.Context) *ConsentsByConsentIDPutParams {
	var ()
	return &ConsentsByConsentIDPutParams{

		Context: ctx,
	}
}

// NewConsentsByConsentIDPutParamsWithHTTPClient creates a new ConsentsByConsentIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsentsByConsentIDPutParamsWithHTTPClient(client *http.Client) *ConsentsByConsentIDPutParams {
	var ()
	return &ConsentsByConsentIDPutParams{
		HTTPClient: client,
	}
}

/*ConsentsByConsentIDPutParams contains all the parameters to send to the API endpoint
for the consents by consent Id put operation typically these are written to a http.Request
*/
type ConsentsByConsentIDPutParams struct {

	/*Digest
	  Is contained if and only if the "Signature" element is contained in the header of the request.

	*/
	Digest *string
	/*PSUCorporateID
	  Might be mandated in the ASPSP's documentation. Only used in a corporate context.

	*/
	PSUCorporateID *string
	/*PSUCorporateIDType
	  Might be mandated in the ASPSPs documentation. Only used in a corporate context.

	*/
	PSUCorporateIDType *string
	/*PSUID
	  Might be mandated in the ASPSP's documentation, if OAuth is not chosen as Pre-Step.

	*/
	PSUID *string
	/*PSUIDType
	  Type of the PSU-ID, needed in scenarios where PSUs have several PSU-IDs as access possibility.

	*/
	PSUIDType *string
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
	/*Body*/
	Body *models.ConsentsRequest144
	/*ConsentID
	  ID of the corresponding consent object as returned by an Account Information Consent Request

	*/
	ConsentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithTimeout(timeout time.Duration) *ConsentsByConsentIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithContext(ctx context.Context) *ConsentsByConsentIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithHTTPClient(client *http.Client) *ConsentsByConsentIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDigest adds the digest to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithDigest(digest *string) *ConsentsByConsentIDPutParams {
	o.SetDigest(digest)
	return o
}

// SetDigest adds the digest to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetDigest(digest *string) {
	o.Digest = digest
}

// WithPSUCorporateID adds the pSUCorporateID to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithPSUCorporateID(pSUCorporateID *string) *ConsentsByConsentIDPutParams {
	o.SetPSUCorporateID(pSUCorporateID)
	return o
}

// SetPSUCorporateID adds the pSUCorporateId to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetPSUCorporateID(pSUCorporateID *string) {
	o.PSUCorporateID = pSUCorporateID
}

// WithPSUCorporateIDType adds the pSUCorporateIDType to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithPSUCorporateIDType(pSUCorporateIDType *string) *ConsentsByConsentIDPutParams {
	o.SetPSUCorporateIDType(pSUCorporateIDType)
	return o
}

// SetPSUCorporateIDType adds the pSUCorporateIdType to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetPSUCorporateIDType(pSUCorporateIDType *string) {
	o.PSUCorporateIDType = pSUCorporateIDType
}

// WithPSUID adds the pSUID to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithPSUID(pSUID *string) *ConsentsByConsentIDPutParams {
	o.SetPSUID(pSUID)
	return o
}

// SetPSUID adds the pSUId to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetPSUID(pSUID *string) {
	o.PSUID = pSUID
}

// WithPSUIDType adds the pSUIDType to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithPSUIDType(pSUIDType *string) *ConsentsByConsentIDPutParams {
	o.SetPSUIDType(pSUIDType)
	return o
}

// SetPSUIDType adds the pSUIdType to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetPSUIDType(pSUIDType *string) {
	o.PSUIDType = pSUIDType
}

// WithSignature adds the signature to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithSignature(signature *string) *ConsentsByConsentIDPutParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithTPPSignatureCertificate adds the tPPSignatureCertificate to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithTPPSignatureCertificate(tPPSignatureCertificate *string) *ConsentsByConsentIDPutParams {
	o.SetTPPSignatureCertificate(tPPSignatureCertificate)
	return o
}

// SetTPPSignatureCertificate adds the tPPSignatureCertificate to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetTPPSignatureCertificate(tPPSignatureCertificate *string) {
	o.TPPSignatureCertificate = tPPSignatureCertificate
}

// WithXRequestID adds the xRequestID to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithXRequestID(xRequestID strfmt.UUID) *ConsentsByConsentIDPutParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetXRequestID(xRequestID strfmt.UUID) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithBody(body *models.ConsentsRequest144) *ConsentsByConsentIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetBody(body *models.ConsentsRequest144) {
	o.Body = body
}

// WithConsentID adds the consentID to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) WithConsentID(consentID string) *ConsentsByConsentIDPutParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the consents by consent Id put params
func (o *ConsentsByConsentIDPutParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsentsByConsentIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PSUCorporateID != nil {

		// header param PSU-Corporate-ID
		if err := r.SetHeaderParam("PSU-Corporate-ID", *o.PSUCorporateID); err != nil {
			return err
		}

	}

	if o.PSUCorporateIDType != nil {

		// header param PSU-Corporate-ID-Type
		if err := r.SetHeaderParam("PSU-Corporate-ID-Type", *o.PSUCorporateIDType); err != nil {
			return err
		}

	}

	if o.PSUID != nil {

		// header param PSU-ID
		if err := r.SetHeaderParam("PSU-ID", *o.PSUID); err != nil {
			return err
		}

	}

	if o.PSUIDType != nil {

		// header param PSU-ID-Type
		if err := r.SetHeaderParam("PSU-ID-Type", *o.PSUIDType); err != nil {
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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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