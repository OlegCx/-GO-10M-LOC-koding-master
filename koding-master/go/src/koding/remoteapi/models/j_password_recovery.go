package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// JPasswordRecovery j password recovery
// swagger:model JPasswordRecovery
type JPasswordRecovery struct {

	// id
	ID string `json:"_id,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// expires
	Expires bool `json:"expires,omitempty"`

	// expires at
	ExpiresAt strfmt.Date `json:"expiresAt,omitempty"`

	// redeemed at
	RedeemedAt strfmt.Date `json:"redeemedAt,omitempty"`

	// requested at
	RequestedAt strfmt.Date `json:"requestedAt,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this j password recovery
func (m *JPasswordRecovery) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
