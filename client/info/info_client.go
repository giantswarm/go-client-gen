// Code generated by go-swagger; DO NOT EDIT.

package info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new info API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for info API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetInfo gets information on the installation

Returns a set of details on the installation. The output varies based
on the provider used in the installation.

This information is useful for example when creating new cluster, to
prevent creating clusters with more worker nodes than possible.

### Example for an AWS-based installation

```json
{
  "general": {
    "installation_name": "shire",
    "provider": "aws",
    "datacenter": "eu-central-1"
  },
  "workers": {
    "count_per_cluster": {
      "max": null,
      "default": 3
    },
    "instance_type": {
      "options": [
        "m3.medium", "m3.large", "m3.xlarge"
      ],
      "default": "m3.large"
    }
  }
}
```

### Example for a KVM-based installation

```json
{
  "general": {
    "installation_name": "isengard",
    "provider": "kvm",
    "datacenter": "string"
  },
  "workers": {
    "count_per_cluster": {
      "max": 8,
      "default": 3
    },
  }
}
```

*/
func (a *Client) GetInfo(params *GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInfo",
		Method:             "GET",
		PathPattern:        "/v4/info/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInfoOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
