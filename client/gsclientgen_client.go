// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/giantswarm/gsclientgen/client/apps"
	"github.com/giantswarm/gsclientgen/client/auth_tokens"
	"github.com/giantswarm/gsclientgen/client/clusters"
	"github.com/giantswarm/gsclientgen/client/info"
	"github.com/giantswarm/gsclientgen/client/key_pairs"
	"github.com/giantswarm/gsclientgen/client/organizations"
	"github.com/giantswarm/gsclientgen/client/releases"
	"github.com/giantswarm/gsclientgen/client/users"
)

// Default gsclientgen HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new gsclientgen HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Gsclientgen {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new gsclientgen HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Gsclientgen {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new gsclientgen client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Gsclientgen {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Gsclientgen)
	cli.Transport = transport

	cli.Apps = apps.New(transport, formats)

	cli.AuthTokens = auth_tokens.New(transport, formats)

	cli.Clusters = clusters.New(transport, formats)

	cli.Info = info.New(transport, formats)

	cli.KeyPairs = key_pairs.New(transport, formats)

	cli.Organizations = organizations.New(transport, formats)

	cli.Releases = releases.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Gsclientgen is a client for gsclientgen
type Gsclientgen struct {
	Apps *apps.Client

	AuthTokens *auth_tokens.Client

	Clusters *clusters.Client

	Info *info.Client

	KeyPairs *key_pairs.Client

	Organizations *organizations.Client

	Releases *releases.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Gsclientgen) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Apps.SetTransport(transport)

	c.AuthTokens.SetTransport(transport)

	c.Clusters.SetTransport(transport)

	c.Info.SetTransport(transport)

	c.KeyPairs.SetTransport(transport)

	c.Organizations.SetTransport(transport)

	c.Releases.SetTransport(transport)

	c.Users.SetTransport(transport)

}
