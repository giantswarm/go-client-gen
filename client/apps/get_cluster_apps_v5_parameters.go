// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewGetClusterAppsV5Params creates a new GetClusterAppsV5Params object
// with the default values initialized.
func NewGetClusterAppsV5Params() *GetClusterAppsV5Params {
	var ()
	return &GetClusterAppsV5Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterAppsV5ParamsWithTimeout creates a new GetClusterAppsV5Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterAppsV5ParamsWithTimeout(timeout time.Duration) *GetClusterAppsV5Params {
	var ()
	return &GetClusterAppsV5Params{

		timeout: timeout,
	}
}

// NewGetClusterAppsV5ParamsWithContext creates a new GetClusterAppsV5Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterAppsV5ParamsWithContext(ctx context.Context) *GetClusterAppsV5Params {
	var ()
	return &GetClusterAppsV5Params{

		Context: ctx,
	}
}

// NewGetClusterAppsV5ParamsWithHTTPClient creates a new GetClusterAppsV5Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterAppsV5ParamsWithHTTPClient(client *http.Client) *GetClusterAppsV5Params {
	var ()
	return &GetClusterAppsV5Params{
		HTTPClient: client,
	}
}

/*GetClusterAppsV5Params contains all the parameters to send to the API endpoint
for the get cluster apps v5 operation typically these are written to a http.Request
*/
type GetClusterAppsV5Params struct {

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
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) WithTimeout(timeout time.Duration) *GetClusterAppsV5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) WithContext(ctx context.Context) *GetClusterAppsV5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) WithHTTPClient(client *http.Client) *GetClusterAppsV5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *GetClusterAppsV5Params {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *GetClusterAppsV5Params {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) WithXRequestID(xRequestID *string) *GetClusterAppsV5Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithClusterID adds the clusterID to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) WithClusterID(clusterID string) *GetClusterAppsV5Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster apps v5 params
func (o *GetClusterAppsV5Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterAppsV5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
