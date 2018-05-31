// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Accept Accept
// swagger:model Accept
type Accept string

const (

	// AcceptXML captures enum value "xml"
	AcceptXML Accept = "xml"

	// AcceptJSON captures enum value "JSON"
	AcceptJSON Accept = "JSON"

	// AcceptText captures enum value "text"
	AcceptText Accept = "text"
)

// for schema
var acceptEnum []interface{}

func init() {
	var res []Accept
	if err := json.Unmarshal([]byte(`["xml","JSON","text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		acceptEnum = append(acceptEnum, v)
	}
}

func (m Accept) validateAcceptEnum(path, location string, value Accept) error {
	if err := validate.Enum(path, location, value, acceptEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this accept
func (m Accept) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAcceptEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
