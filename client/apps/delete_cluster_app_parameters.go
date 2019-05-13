// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteClusterAppParams creates a new DeleteClusterAppParams object
// with the default values initialized.
func NewDeleteClusterAppParams() *DeleteClusterAppParams {
	var ()
	return &DeleteClusterAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterAppParamsWithTimeout creates a new DeleteClusterAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterAppParamsWithTimeout(timeout time.Duration) *DeleteClusterAppParams {
	var ()
	return &DeleteClusterAppParams{

		timeout: timeout,
	}
}

// NewDeleteClusterAppParamsWithContext creates a new DeleteClusterAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterAppParamsWithContext(ctx context.Context) *DeleteClusterAppParams {
	var ()
	return &DeleteClusterAppParams{

		Context: ctx,
	}
}

// NewDeleteClusterAppParamsWithHTTPClient creates a new DeleteClusterAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterAppParamsWithHTTPClient(client *http.Client) *DeleteClusterAppParams {
	var ()
	return &DeleteClusterAppParams{
		HTTPClient: client,
	}
}

/*DeleteClusterAppParams contains all the parameters to send to the API endpoint
for the delete cluster app operation typically these are written to a http.Request
*/
type DeleteClusterAppParams struct {

	/*Authorization
	  As described in the [authentication](#section/Authentication) section


	*/
	Authorization string
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

// WithTimeout adds the timeout to the delete cluster app params
func (o *DeleteClusterAppParams) WithTimeout(timeout time.Duration) *DeleteClusterAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster app params
func (o *DeleteClusterAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster app params
func (o *DeleteClusterAppParams) WithContext(ctx context.Context) *DeleteClusterAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster app params
func (o *DeleteClusterAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster app params
func (o *DeleteClusterAppParams) WithHTTPClient(client *http.Client) *DeleteClusterAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster app params
func (o *DeleteClusterAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete cluster app params
func (o *DeleteClusterAppParams) WithAuthorization(authorization string) *DeleteClusterAppParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete cluster app params
func (o *DeleteClusterAppParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the delete cluster app params
func (o *DeleteClusterAppParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *DeleteClusterAppParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the delete cluster app params
func (o *DeleteClusterAppParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete cluster app params
func (o *DeleteClusterAppParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *DeleteClusterAppParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete cluster app params
func (o *DeleteClusterAppParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the delete cluster app params
func (o *DeleteClusterAppParams) WithXRequestID(xRequestID *string) *DeleteClusterAppParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete cluster app params
func (o *DeleteClusterAppParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAppName adds the appName to the delete cluster app params
func (o *DeleteClusterAppParams) WithAppName(appName string) *DeleteClusterAppParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the delete cluster app params
func (o *DeleteClusterAppParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithClusterID adds the clusterID to the delete cluster app params
func (o *DeleteClusterAppParams) WithClusterID(clusterID string) *DeleteClusterAppParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster app params
func (o *DeleteClusterAppParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

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
