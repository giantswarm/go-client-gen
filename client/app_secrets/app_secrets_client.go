// Code generated by go-swagger; DO NOT EDIT.

package app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new app secrets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for app secrets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateClusterAppSecretV4 creates secret v4

This operation allows you to create a Secret for a specific app. The app does
not have to exist before hand.

If the app does exist, this endpoint will ensure that the App CR gets it's
`spec.user_config.secret` field set correctly.

However, if the app exists and the `spec.user_config.secret` is already set to something,
then this request will fail. You will in that case most likely want to
update the Secret using the `PATCH /v4/clusters/{cluster_id}/apps/{app_name}/secret/`
endpoint.

For apps on v5 clusters, please use the v5 version of this endpoint.

### Example POST request
```json
  {
    "secret": "value"
  }
```

*/
func (a *Client) CreateClusterAppSecretV4(params *CreateClusterAppSecretV4Params, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterAppSecretV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterAppSecretV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createClusterAppSecretV4",
		Method:             "PUT",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterAppSecretV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterAppSecretV4OK), nil

}

/*
CreateClusterAppSecretV5 creates secret v5

This operation allows you to create a Secret for a specific app. The app does
not have to exist before hand.

If the app does exist, this endpoint will ensure that the App CR gets it's
`spec.user_config.secret` field set correctly.

However, if the app exists and the `spec.user_config.secret` is already set to something,
then this request will fail. You will in that case most likely want to
update the Secret using the `PATCH /v5/clusters/{cluster_id}/apps/{app_name}/secret/`
endpoint.

### Example POST request
```json
  {
    "secret": "value"
  }
```

*/
func (a *Client) CreateClusterAppSecretV5(params *CreateClusterAppSecretV5Params, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterAppSecretV5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterAppSecretV5Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createClusterAppSecretV5",
		Method:             "PUT",
		PathPattern:        "/v5/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterAppSecretV5Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterAppSecretV5OK), nil

}

/*
DeleteClusterAppSecretV4 deletes a secret v4

This operation allows a user to delete an app's Secret if it has been named according to the convention of {app-name}-user-secrets and
stored in the cluster ID namespace.

Calling this endpoint will delete the Secret, and also remove the reference to the Secret in the (spec.user_config.secret field) from the app.

For apps on v5 clusters, please use the v5 version of this endpoint.

*/
func (a *Client) DeleteClusterAppSecretV4(params *DeleteClusterAppSecretV4Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterAppSecretV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterAppSecretV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteClusterAppSecretV4",
		Method:             "DELETE",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClusterAppSecretV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterAppSecretV4OK), nil

}

/*
DeleteClusterAppSecretV5 deletes a secret v5

This operation allows a user to delete an app's Secret if it has been named according to the convention of {app-name}-user-secrets and
stored in the cluster ID namespace.

Calling this endpoint will delete the Secret, and also remove the reference to the Secret in the (spec.user_config.secret field) from the app.

*/
func (a *Client) DeleteClusterAppSecretV5(params *DeleteClusterAppSecretV5Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterAppSecretV5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterAppSecretV5Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteClusterAppSecretV5",
		Method:             "DELETE",
		PathPattern:        "/v5/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClusterAppSecretV5Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterAppSecretV5OK), nil

}

/*
GetClusterAppSecretV4 gets secret v4

This operation allows you to fetch the Secret associated
with an app.

For apps on v5 clusters, please use the v5 version of this endpoint.

*/
func (a *Client) GetClusterAppSecretV4(params *GetClusterAppSecretV4Params, authInfo runtime.ClientAuthInfoWriter) (*GetClusterAppSecretV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterAppSecretV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterAppSecretV4",
		Method:             "GET",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterAppSecretV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterAppSecretV4OK), nil

}

/*
GetClusterAppSecretV5 gets secret v5

This operation allows you to fetch the Secret associated
with an app.

*/
func (a *Client) GetClusterAppSecretV5(params *GetClusterAppSecretV5Params, authInfo runtime.ClientAuthInfoWriter) (*GetClusterAppSecretV5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterAppSecretV5Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterAppSecretV5",
		Method:             "GET",
		PathPattern:        "/v5/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterAppSecretV5Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterAppSecretV5OK), nil

}

/*
ModifyClusterAppSecretV4 modifies secret v4

This operation allows you to modify the Secret for a specific app.
It's only possible to modify Secrets that have been named according to the convention of
{app-name}-user-secrets and stored in the cluster ID namespace.

The full values key of the Secret will be replaced by the JSON body
of your request.

For apps on v5 clusters, please use the v5 version of this endpoint.

### Example PATCH request
```json
  {
    "secret": "new-value"
  }
```

If the Secret contained content like:

```json
  {
    "secret": "old-value",
    "secret2": "another-old-value"
  }
```

Then the "secret2" will be removed, and "secret" will be set to "new-value"

*/
func (a *Client) ModifyClusterAppSecretV4(params *ModifyClusterAppSecretV4Params, authInfo runtime.ClientAuthInfoWriter) (*ModifyClusterAppSecretV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterAppSecretV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyClusterAppSecretV4",
		Method:             "PATCH",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyClusterAppSecretV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterAppSecretV4OK), nil

}

/*
ModifyClusterAppSecretV5 modifies secret v5

This operation allows you to modify the Secret for a specific app.
It's only possible to modify Secrets that have been named according to the convention of
{app-name}-user-secrets and stored in the cluster ID namespace.

The full values key of the Secret will be replaced by the JSON body
of your request.

### Example PATCH request
```json
  {
    "secret": "new-value"
  }
```

If the Secret contained content like:

```json
  {
    "secret": "old-value",
    "secret2": "another-old-value"
  }
```

Then the "secret2" will be removed, and "secret" will be set to "new-value"

*/
func (a *Client) ModifyClusterAppSecretV5(params *ModifyClusterAppSecretV5Params, authInfo runtime.ClientAuthInfoWriter) (*ModifyClusterAppSecretV5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterAppSecretV5Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyClusterAppSecretV5",
		Method:             "PATCH",
		PathPattern:        "/v5/clusters/{cluster_id}/apps/{app_name}/secret/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyClusterAppSecretV5Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterAppSecretV5OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
