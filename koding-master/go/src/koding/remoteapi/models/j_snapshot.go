package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// JSnapshot j snapshot
// swagger:model JSnapshot
type JSnapshot struct {

	// id
	ID string `json:"_id,omitempty"`

	// created at
	CreatedAt strfmt.Date `json:"createdAt,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// machine Id
	MachineID string `json:"machineId,omitempty"`

	// origin Id
	OriginID string `json:"originId,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// snapshot Id
	SnapshotID string `json:"snapshotId,omitempty"`

	// storage size
	StorageSize string `json:"storageSize,omitempty"`
}

// Validate validates this j snapshot
func (m *JSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}