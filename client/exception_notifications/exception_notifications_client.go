// Code generated by go-swagger; DO NOT EDIT.

package exception_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new exception notifications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for exception notifications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddExceptionNotification notifies giant swarm that an error has occured in one of our clients

This endpoint is used to notify Giant Swarm that an error has occured in one of our clients (like our Web UI or gsctl). It is not intended to be called manually, our clients are configured to report errors when they occur.

Find us in your Slack support channel if you want to contact us about any immediate issues.

*/
func (a *Client) AddExceptionNotification(params *AddExceptionNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*AddExceptionNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddExceptionNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addExceptionNotification",
		Method:             "POST",
		PathPattern:        "/v5/exception-notifications/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddExceptionNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddExceptionNotificationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
