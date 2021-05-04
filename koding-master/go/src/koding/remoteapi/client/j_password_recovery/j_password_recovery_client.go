package j_password_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j password recovery API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j password recovery API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
JPasswordRecoveryFetchRegistrationDetails Method JPasswordRecovery.fetchRegistrationDetails
*/
func (a *Client) JPasswordRecoveryFetchRegistrationDetails(params *JPasswordRecoveryFetchRegistrationDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*JPasswordRecoveryFetchRegistrationDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJPasswordRecoveryFetchRegistrationDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JPasswordRecovery.fetchRegistrationDetails",
		Method:             "POST",
		PathPattern:        "/remote.api/JPasswordRecovery.fetchRegistrationDetails",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JPasswordRecoveryFetchRegistrationDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JPasswordRecoveryFetchRegistrationDetailsOK), nil

}

/*
JPasswordRecoveryRecoverPassword j password recovery recover password API
*/
func (a *Client) JPasswordRecoveryRecoverPassword(params *JPasswordRecoveryRecoverPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*JPasswordRecoveryRecoverPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJPasswordRecoveryRecoverPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JPasswordRecovery.recoverPassword",
		Method:             "POST",
		PathPattern:        "/remote.api/JPasswordRecovery.recoverPassword",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JPasswordRecoveryRecoverPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JPasswordRecoveryRecoverPasswordOK), nil

}

/*
JPasswordRecoveryResendVerification j password recovery resend verification API
*/
func (a *Client) JPasswordRecoveryResendVerification(params *JPasswordRecoveryResendVerificationParams, authInfo runtime.ClientAuthInfoWriter) (*JPasswordRecoveryResendVerificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJPasswordRecoveryResendVerificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JPasswordRecovery.resendVerification",
		Method:             "POST",
		PathPattern:        "/remote.api/JPasswordRecovery.resendVerification",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JPasswordRecoveryResendVerificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JPasswordRecoveryResendVerificationOK), nil

}

/*
JPasswordRecoveryResetPassword Method JPasswordRecovery.resetPassword
*/
func (a *Client) JPasswordRecoveryResetPassword(params *JPasswordRecoveryResetPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*JPasswordRecoveryResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJPasswordRecoveryResetPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JPasswordRecovery.resetPassword",
		Method:             "POST",
		PathPattern:        "/remote.api/JPasswordRecovery.resetPassword",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JPasswordRecoveryResetPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JPasswordRecoveryResetPasswordOK), nil

}

/*
JPasswordRecoveryValidate Method JPasswordRecovery.validate
*/
func (a *Client) JPasswordRecoveryValidate(params *JPasswordRecoveryValidateParams, authInfo runtime.ClientAuthInfoWriter) (*JPasswordRecoveryValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJPasswordRecoveryValidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JPasswordRecovery.validate",
		Method:             "POST",
		PathPattern:        "/remote.api/JPasswordRecovery.validate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JPasswordRecoveryValidateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JPasswordRecoveryValidateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}