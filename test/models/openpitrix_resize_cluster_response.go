// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixResizeClusterResponse openpitrix resize cluster response
// swagger:model openpitrixResizeClusterResponse
type OpenpitrixResizeClusterResponse struct {

	// id of cluster resized
	ClusterID string `json:"cluster_id,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`
}

// Validate validates this openpitrix resize cluster response
func (m *OpenpitrixResizeClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixResizeClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixResizeClusterResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixResizeClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
