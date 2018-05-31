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
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openpsd/psd2-server/models"
)

// NewPaymentIDByPaymentServiceAndPaymentProductPutParams creates a new PaymentIDByPaymentServiceAndPaymentProductPutParams object
// with the default values initialized.
func NewPaymentIDByPaymentServiceAndPaymentProductPutParams() *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	var ()
	return &PaymentIDByPaymentServiceAndPaymentProductPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutParamsWithTimeout creates a new PaymentIDByPaymentServiceAndPaymentProductPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPaymentIDByPaymentServiceAndPaymentProductPutParamsWithTimeout(timeout time.Duration) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	var ()
	return &PaymentIDByPaymentServiceAndPaymentProductPutParams{

		timeout: timeout,
	}
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutParamsWithContext creates a new PaymentIDByPaymentServiceAndPaymentProductPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPaymentIDByPaymentServiceAndPaymentProductPutParamsWithContext(ctx context.Context) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	var ()
	return &PaymentIDByPaymentServiceAndPaymentProductPutParams{

		Context: ctx,
	}
}

// NewPaymentIDByPaymentServiceAndPaymentProductPutParamsWithHTTPClient creates a new PaymentIDByPaymentServiceAndPaymentProductPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPaymentIDByPaymentServiceAndPaymentProductPutParamsWithHTTPClient(client *http.Client) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	var ()
	return &PaymentIDByPaymentServiceAndPaymentProductPutParams{
		HTTPClient: client,
	}
}

/*PaymentIDByPaymentServiceAndPaymentProductPutParams contains all the parameters to send to the API endpoint
for the payment Id by payment service and payment product put operation typically these are written to a http.Request
*/
type PaymentIDByPaymentServiceAndPaymentProductPutParams struct {

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
	Body *models.PaymentidRequest
	/*PaymentProduct
	  The addressed payment product, e.g. SCT. The default list of products supported in this standard is, pain.001-sepa-credit-transfers, pain.001-instant-sepa-credit-transfers, pain.001-target-2-payments or pain.001-cross-border-credit-transfers. Further products might be published by the ASPSP within its XS2A documentation.

	*/
	PaymentProduct string
	/*PaymentService
	  The payment service to use.

	*/
	PaymentService string
	/*PaymentID
	  payment Id

	*/
	PaymentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithTimeout(timeout time.Duration) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithContext(ctx context.Context) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithHTTPClient(client *http.Client) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDigest adds the digest to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithDigest(digest *string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetDigest(digest)
	return o
}

// SetDigest adds the digest to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetDigest(digest *string) {
	o.Digest = digest
}

// WithPSUCorporateID adds the pSUCorporateID to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithPSUCorporateID(pSUCorporateID *string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetPSUCorporateID(pSUCorporateID)
	return o
}

// SetPSUCorporateID adds the pSUCorporateId to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetPSUCorporateID(pSUCorporateID *string) {
	o.PSUCorporateID = pSUCorporateID
}

// WithPSUCorporateIDType adds the pSUCorporateIDType to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithPSUCorporateIDType(pSUCorporateIDType *string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetPSUCorporateIDType(pSUCorporateIDType)
	return o
}

// SetPSUCorporateIDType adds the pSUCorporateIdType to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetPSUCorporateIDType(pSUCorporateIDType *string) {
	o.PSUCorporateIDType = pSUCorporateIDType
}

// WithPSUID adds the pSUID to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithPSUID(pSUID *string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetPSUID(pSUID)
	return o
}

// SetPSUID adds the pSUId to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetPSUID(pSUID *string) {
	o.PSUID = pSUID
}

// WithPSUIDType adds the pSUIDType to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithPSUIDType(pSUIDType *string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetPSUIDType(pSUIDType)
	return o
}

// SetPSUIDType adds the pSUIdType to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetPSUIDType(pSUIDType *string) {
	o.PSUIDType = pSUIDType
}

// WithSignature adds the signature to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithSignature(signature *string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithTPPSignatureCertificate adds the tPPSignatureCertificate to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithTPPSignatureCertificate(tPPSignatureCertificate *string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetTPPSignatureCertificate(tPPSignatureCertificate)
	return o
}

// SetTPPSignatureCertificate adds the tPPSignatureCertificate to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetTPPSignatureCertificate(tPPSignatureCertificate *string) {
	o.TPPSignatureCertificate = tPPSignatureCertificate
}

// WithXRequestID adds the xRequestID to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithXRequestID(xRequestID strfmt.UUID) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetXRequestID(xRequestID strfmt.UUID) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithBody(body *models.PaymentidRequest) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetBody(body *models.PaymentidRequest) {
	o.Body = body
}

// WithPaymentProduct adds the paymentProduct to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithPaymentProduct(paymentProduct string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetPaymentProduct(paymentProduct)
	return o
}

// SetPaymentProduct adds the paymentProduct to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetPaymentProduct(paymentProduct string) {
	o.PaymentProduct = paymentProduct
}

// WithPaymentService adds the paymentService to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithPaymentService(paymentService string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetPaymentService(paymentService)
	return o
}

// SetPaymentService adds the paymentService to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetPaymentService(paymentService string) {
	o.PaymentService = paymentService
}

// WithPaymentID adds the paymentID to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WithPaymentID(paymentID string) *PaymentIDByPaymentServiceAndPaymentProductPutParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the payment Id by payment service and payment product put params
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) SetPaymentID(paymentID string) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentIDByPaymentServiceAndPaymentProductPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param payment-product
	if err := r.SetPathParam("payment-product", o.PaymentProduct); err != nil {
		return err
	}

	// path param payment-service
	if err := r.SetPathParam("payment-service", o.PaymentService); err != nil {
		return err
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}