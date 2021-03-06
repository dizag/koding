package j_proxy_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j proxy filter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j proxy filter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIJProxyFilterCreate post remote API j proxy filter create API
*/
func (a *Client) PostRemoteAPIJProxyFilterCreate(params *PostRemoteAPIJProxyFilterCreateParams) (*PostRemoteAPIJProxyFilterCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJProxyFilterCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJProxyFilterCreate",
		Method:             "POST",
		PathPattern:        "/remote.api/JProxyFilter.create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJProxyFilterCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJProxyFilterCreateOK), nil

}

/*
PostRemoteAPIJProxyFilterRemoveID post remote API j proxy filter remove ID API
*/
func (a *Client) PostRemoteAPIJProxyFilterRemoveID(params *PostRemoteAPIJProxyFilterRemoveIDParams) (*PostRemoteAPIJProxyFilterRemoveIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJProxyFilterRemoveIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJProxyFilterRemoveID",
		Method:             "POST",
		PathPattern:        "/remote.api/JProxyFilter.remove/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJProxyFilterRemoveIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJProxyFilterRemoveIDOK), nil

}

/*
PostRemoteAPIJProxyFilterSome post remote API j proxy filter some API
*/
func (a *Client) PostRemoteAPIJProxyFilterSome(params *PostRemoteAPIJProxyFilterSomeParams) (*PostRemoteAPIJProxyFilterSomeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJProxyFilterSomeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJProxyFilterSome",
		Method:             "POST",
		PathPattern:        "/remote.api/JProxyFilter.some",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJProxyFilterSomeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJProxyFilterSomeOK), nil

}

/*
PostRemoteAPIJProxyFilterUpdateID post remote API j proxy filter update ID API
*/
func (a *Client) PostRemoteAPIJProxyFilterUpdateID(params *PostRemoteAPIJProxyFilterUpdateIDParams) (*PostRemoteAPIJProxyFilterUpdateIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJProxyFilterUpdateIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJProxyFilterUpdateID",
		Method:             "POST",
		PathPattern:        "/remote.api/JProxyFilter.update/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJProxyFilterUpdateIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJProxyFilterUpdateIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
