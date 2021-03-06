package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ConfigGet gets server configuration

Get the RackHD server configuration properties.
*/
func (a *Client) ConfigGet(params *ConfigGetParams, authInfo runtime.ClientAuthInfoWriter) (*ConfigGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "configGet",
		Method:             "GET",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConfigGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConfigGetOK), nil

}

/*
ConfigPatch patches server configuration

Modify the RackHD server configuration.
*/
func (a *Client) ConfigPatch(params *ConfigPatchParams, authInfo runtime.ClientAuthInfoWriter) (*ConfigPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "configPatch",
		Method:             "PATCH",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConfigPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConfigPatchOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
