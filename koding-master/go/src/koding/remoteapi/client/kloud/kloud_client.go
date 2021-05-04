package kloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new kloud API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kloud API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
KloudAddAdmin kloud add admin API
*/
func (a *Client) KloudAddAdmin(params *KloudAddAdminParams, authInfo runtime.ClientAuthInfoWriter) (*KloudAddAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudAddAdminParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.addAdmin",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.addAdmin",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudAddAdminReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudAddAdminOK), nil

}

/*
KloudBootstrap kloud bootstrap API
*/
func (a *Client) KloudBootstrap(params *KloudBootstrapParams, authInfo runtime.ClientAuthInfoWriter) (*KloudBootstrapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudBootstrapParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.bootstrap",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.bootstrap",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudBootstrapReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudBootstrapOK), nil

}

/*
KloudBuild kloud build API
*/
func (a *Client) KloudBuild(params *KloudBuildParams, authInfo runtime.ClientAuthInfoWriter) (*KloudBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.build",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.build",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudBuildOK), nil

}

/*
KloudBuildStack kloud build stack API
*/
func (a *Client) KloudBuildStack(params *KloudBuildStackParams, authInfo runtime.ClientAuthInfoWriter) (*KloudBuildStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudBuildStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.buildStack",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.buildStack",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudBuildStackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudBuildStackOK), nil

}

/*
KloudCheckCredential kloud check credential API
*/
func (a *Client) KloudCheckCredential(params *KloudCheckCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*KloudCheckCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudCheckCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.checkCredential",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.checkCredential",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudCheckCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudCheckCredentialOK), nil

}

/*
KloudCheckTemplate kloud check template API
*/
func (a *Client) KloudCheckTemplate(params *KloudCheckTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*KloudCheckTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudCheckTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.checkTemplate",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.checkTemplate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudCheckTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudCheckTemplateOK), nil

}

/*
KloudDestroy kloud destroy API
*/
func (a *Client) KloudDestroy(params *KloudDestroyParams, authInfo runtime.ClientAuthInfoWriter) (*KloudDestroyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudDestroyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.destroy",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.destroy",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudDestroyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudDestroyOK), nil

}

/*
KloudDestroyStack kloud destroy stack API
*/
func (a *Client) KloudDestroyStack(params *KloudDestroyStackParams, authInfo runtime.ClientAuthInfoWriter) (*KloudDestroyStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudDestroyStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.destroyStack",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.destroyStack",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudDestroyStackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudDestroyStackOK), nil

}

/*
KloudEvent kloud event API
*/
func (a *Client) KloudEvent(params *KloudEventParams, authInfo runtime.ClientAuthInfoWriter) (*KloudEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.event",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.event",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudEventOK), nil

}

/*
KloudInfo kloud info API
*/
func (a *Client) KloudInfo(params *KloudInfoParams, authInfo runtime.ClientAuthInfoWriter) (*KloudInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.info",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.info",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudInfoOK), nil

}

/*
KloudMigrate kloud migrate API
*/
func (a *Client) KloudMigrate(params *KloudMigrateParams, authInfo runtime.ClientAuthInfoWriter) (*KloudMigrateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudMigrateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.migrate",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.migrate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudMigrateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudMigrateOK), nil

}

/*
KloudPing kloud ping API
*/
func (a *Client) KloudPing(params *KloudPingParams, authInfo runtime.ClientAuthInfoWriter) (*KloudPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.ping",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.ping",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudPingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudPingOK), nil

}

/*
KloudRemoveAdmin kloud remove admin API
*/
func (a *Client) KloudRemoveAdmin(params *KloudRemoveAdminParams, authInfo runtime.ClientAuthInfoWriter) (*KloudRemoveAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudRemoveAdminParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.removeAdmin",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.removeAdmin",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudRemoveAdminReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudRemoveAdminOK), nil

}

/*
KloudRestart kloud restart API
*/
func (a *Client) KloudRestart(params *KloudRestartParams, authInfo runtime.ClientAuthInfoWriter) (*KloudRestartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudRestartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.restart",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.restart",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudRestartReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudRestartOK), nil

}

/*
KloudStart kloud start API
*/
func (a *Client) KloudStart(params *KloudStartParams, authInfo runtime.ClientAuthInfoWriter) (*KloudStartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudStartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.start",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.start",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudStartReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudStartOK), nil

}

/*
KloudStop kloud stop API
*/
func (a *Client) KloudStop(params *KloudStopParams, authInfo runtime.ClientAuthInfoWriter) (*KloudStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKloudStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Kloud.stop",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.stop",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KloudStopReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KloudStopOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}