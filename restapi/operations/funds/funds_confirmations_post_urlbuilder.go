// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package funds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// FundsConfirmationsPostURL generates an URL for the funds confirmations post operation
type FundsConfirmationsPostURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FundsConfirmationsPostURL) WithBasePath(bp string) *FundsConfirmationsPostURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FundsConfirmationsPostURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FundsConfirmationsPostURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/funds-confirmations"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/psd2/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FundsConfirmationsPostURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FundsConfirmationsPostURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FundsConfirmationsPostURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FundsConfirmationsPostURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FundsConfirmationsPostURL")
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
func (o *FundsConfirmationsPostURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
