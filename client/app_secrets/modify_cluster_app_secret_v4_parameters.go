// Code generated by go-swagger; DO NOT EDIT.

package app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// NewModifyClusterAppSecretV4Params creates a new ModifyClusterAppSecretV4Params object
// with the default values initialized.
func NewModifyClusterAppSecretV4Params() *ModifyClusterAppSecretV4Params {
	var ()
	return &ModifyClusterAppSecretV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyClusterAppSecretV4ParamsWithTimeout creates a new ModifyClusterAppSecretV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyClusterAppSecretV4ParamsWithTimeout(timeout time.Duration) *ModifyClusterAppSecretV4Params {
	var ()
	return &ModifyClusterAppSecretV4Params{

		timeout: timeout,
	}
}

// NewModifyClusterAppSecretV4ParamsWithContext creates a new ModifyClusterAppSecretV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewModifyClusterAppSecretV4ParamsWithContext(ctx context.Context) *ModifyClusterAppSecretV4Params {
	var ()
	return &ModifyClusterAppSecretV4Params{

		Context: ctx,
	}
}

// NewModifyClusterAppSecretV4ParamsWithHTTPClient creates a new ModifyClusterAppSecretV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyClusterAppSecretV4ParamsWithHTTPClient(client *http.Client) *ModifyClusterAppSecretV4Params {
	var ()
	return &ModifyClusterAppSecretV4Params{
		HTTPClient: client,
	}
}

/*ModifyClusterAppSecretV4Params contains all the parameters to send to the API endpoint
for the modify cluster app secret v4 operation typically these are written to a http.Request
*/
type ModifyClusterAppSecretV4Params struct {

	/*XGiantSwarmActivity
	  Name of an activity to track, like "list-clusters". This allows to
	analyze several API requests sent in context and gives an idea on
	the purpose.


	*/
	XGiantSwarmActivity *string
	/*XGiantSwarmCmdLine
	  If activity has been issued by a CLI, this header can contain the
	command line


	*/
	XGiantSwarmCmdLine *string
	/*XRequestID
	  A randomly generated key that can be used to track a request throughout
	services of Giant Swarm.


	*/
	XRequestID *string
	/*AppName
	  App Name

	*/
	AppName string
	/*Body*/
	Body models.V4CreateClusterAppSecretRequest
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithTimeout(timeout time.Duration) *ModifyClusterAppSecretV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithContext(ctx context.Context) *ModifyClusterAppSecretV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithHTTPClient(client *http.Client) *ModifyClusterAppSecretV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *ModifyClusterAppSecretV4Params {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *ModifyClusterAppSecretV4Params {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithXRequestID(xRequestID *string) *ModifyClusterAppSecretV4Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAppName adds the appName to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithAppName(appName string) *ModifyClusterAppSecretV4Params {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithBody(body models.V4CreateClusterAppSecretRequest) *ModifyClusterAppSecretV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetBody(body models.V4CreateClusterAppSecretRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) WithClusterID(clusterID string) *ModifyClusterAppSecretV4Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the modify cluster app secret v4 params
func (o *ModifyClusterAppSecretV4Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyClusterAppSecretV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XGiantSwarmActivity != nil {

		// header param X-Giant-Swarm-Activity
		if err := r.SetHeaderParam("X-Giant-Swarm-Activity", *o.XGiantSwarmActivity); err != nil {
			return err
		}

	}

	if o.XGiantSwarmCmdLine != nil {

		// header param X-Giant-Swarm-CmdLine
		if err := r.SetHeaderParam("X-Giant-Swarm-CmdLine", *o.XGiantSwarmCmdLine); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-ID
		if err := r.SetHeaderParam("X-Request-ID", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
