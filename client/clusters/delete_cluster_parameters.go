// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewDeleteClusterParams creates a new DeleteClusterParams object
// with the default values initialized.
func NewDeleteClusterParams() *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterParamsWithTimeout creates a new DeleteClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterParamsWithTimeout(timeout time.Duration) *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{

		timeout: timeout,
	}
}

// NewDeleteClusterParamsWithContext creates a new DeleteClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterParamsWithContext(ctx context.Context) *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{

		Context: ctx,
	}
}

// NewDeleteClusterParamsWithHTTPClient creates a new DeleteClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterParamsWithHTTPClient(client *http.Client) *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{
		HTTPClient: client,
	}
}

/*DeleteClusterParams contains all the parameters to send to the API endpoint
for the delete cluster operation typically these are written to a http.Request
*/
type DeleteClusterParams struct {

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
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) WithTimeout(timeout time.Duration) *DeleteClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster params
func (o *DeleteClusterParams) WithContext(ctx context.Context) *DeleteClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster params
func (o *DeleteClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) WithHTTPClient(client *http.Client) *DeleteClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete cluster params
func (o *DeleteClusterParams) WithAuthorization(authorization string) *DeleteClusterParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete cluster params
func (o *DeleteClusterParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the delete cluster params
func (o *DeleteClusterParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *DeleteClusterParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the delete cluster params
func (o *DeleteClusterParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete cluster params
func (o *DeleteClusterParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *DeleteClusterParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete cluster params
func (o *DeleteClusterParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the delete cluster params
func (o *DeleteClusterParams) WithXRequestID(xRequestID *string) *DeleteClusterParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete cluster params
func (o *DeleteClusterParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithClusterID adds the clusterID to the delete cluster params
func (o *DeleteClusterParams) WithClusterID(clusterID string) *DeleteClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster params
func (o *DeleteClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
