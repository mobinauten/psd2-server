// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * 2018 - OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PaymentIDByPaymentServiceAndPaymentProductPutURL generates an URL for the payment Id by payment service and payment product put operation
type PaymentIDByPaymentServiceAndPaymentProductPutURL struct {
	PaymentProduct string
	PaymentService string
	PaymentID      string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PaymentIDByPaymentServiceAndPaymentProductPutURL) WithBasePath(bp string) *PaymentIDByPaymentServiceAndPaymentProductPutURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PaymentIDByPaymentServiceAndPaymentProductPutURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PaymentIDByPaymentServiceAndPaymentProductPutURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/{payment-service}/{payment-product}/{paymentId}"

	paymentProduct := o.PaymentProduct
	if paymentProduct != "" {
		_path = strings.Replace(_path, "{payment-product}", paymentProduct, -1)
	} else {
		return nil, errors.New("PaymentProduct is required on PaymentIDByPaymentServiceAndPaymentProductPutURL")
	}

	paymentService := o.PaymentService
	if paymentService != "" {
		_path = strings.Replace(_path, "{payment-service}", paymentService, -1)
	} else {
		return nil, errors.New("PaymentService is required on PaymentIDByPaymentServiceAndPaymentProductPutURL")
	}

	paymentID := o.PaymentID
	if paymentID != "" {
		_path = strings.Replace(_path, "{paymentId}", paymentID, -1)
	} else {
		return nil, errors.New("PaymentID is required on PaymentIDByPaymentServiceAndPaymentProductPutURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/psd2/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PaymentIDByPaymentServiceAndPaymentProductPutURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PaymentIDByPaymentServiceAndPaymentProductPutURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PaymentIDByPaymentServiceAndPaymentProductPutURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PaymentIDByPaymentServiceAndPaymentProductPutURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PaymentIDByPaymentServiceAndPaymentProductPutURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PaymentIDByPaymentServiceAndPaymentProductPutURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
