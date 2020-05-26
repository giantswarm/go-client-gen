// Code generated by go-swagger; DO NOT EDIT.

package app_configs

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
)

// NewDeleteClusterAppConfigV4Params creates a new DeleteClusterAppConfigV4Params object
// with the default values initialized.
func NewDeleteClusterAppConfigV4Params() *DeleteClusterAppConfigV4Params {
	var ()
	return &DeleteClusterAppConfigV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterAppConfigV4ParamsWithTimeout creates a new DeleteClusterAppConfigV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterAppConfigV4ParamsWithTimeout(timeout time.Duration) *DeleteClusterAppConfigV4Params {
	var ()
	return &DeleteClusterAppConfigV4Params{

		timeout: timeout,
	}
}

// NewDeleteClusterAppConfigV4ParamsWithContext creates a new DeleteClusterAppConfigV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterAppConfigV4ParamsWithContext(ctx context.Context) *DeleteClusterAppConfigV4Params {
	var ()
	return &DeleteClusterAppConfigV4Params{

		Context: ctx,
	}
}

// NewDeleteClusterAppConfigV4ParamsWithHTTPClient creates a new DeleteClusterAppConfigV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterAppConfigV4ParamsWithHTTPClient(client *http.Client) *DeleteClusterAppConfigV4Params {
	var ()
	return &DeleteClusterAppConfigV4Params{
		HTTPClient: client,
	}
}

/*DeleteClusterAppConfigV4Params contains all the parameters to send to the API endpoint
for the delete cluster app config v4 operation typically these are written to a http.Request
*/
type DeleteClusterAppConfigV4Params struct {

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
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithTimeout(timeout time.Duration) *DeleteClusterAppConfigV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithContext(ctx context.Context) *DeleteClusterAppConfigV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithHTTPClient(client *http.Client) *DeleteClusterAppConfigV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *DeleteClusterAppConfigV4Params {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *DeleteClusterAppConfigV4Params {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithXRequestID(xRequestID *string) *DeleteClusterAppConfigV4Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAppName adds the appName to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithAppName(appName string) *DeleteClusterAppConfigV4Params {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetAppName(appName string) {
	o.AppName = appName
}

// WithClusterID adds the clusterID to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) WithClusterID(clusterID string) *DeleteClusterAppConfigV4Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster app config v4 params
func (o *DeleteClusterAppConfigV4Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterAppConfigV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
