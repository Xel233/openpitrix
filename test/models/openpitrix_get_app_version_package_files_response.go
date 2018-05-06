// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetAppVersionPackageFilesResponse openpitrix get app version package files response
// swagger:model openpitrixGetAppVersionPackageFilesResponse
type OpenpitrixGetAppVersionPackageFilesResponse struct {

	// files
	Files map[string]strfmt.Base64 `json:"files,omitempty"`
}

// Validate validates this openpitrix get app version package files response
func (m *OpenpitrixGetAppVersionPackageFilesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetAppVersionPackageFilesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetAppVersionPackageFilesResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetAppVersionPackageFilesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
