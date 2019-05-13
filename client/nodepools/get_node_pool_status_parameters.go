// Code generated by go-swagger; DO NOT EDIT.

package nodepools

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

// NewGetNodePoolStatusParams creates a new GetNodePoolStatusParams object
// with the default values initialized.
func NewGetNodePoolStatusParams() *GetNodePoolStatusParams {
	var ()
	return &GetNodePoolStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodePoolStatusParamsWithTimeout creates a new GetNodePoolStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodePoolStatusParamsWithTimeout(timeout time.Duration) *GetNodePoolStatusParams {
	var ()
	return &GetNodePoolStatusParams{

		timeout: timeout,
	}
}

// NewGetNodePoolStatusParamsWithContext creates a new GetNodePoolStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodePoolStatusParamsWithContext(ctx context.Context) *GetNodePoolStatusParams {
	var ()
	return &GetNodePoolStatusParams{

		Context: ctx,
	}
}

// NewGetNodePoolStatusParamsWithHTTPClient creates a new GetNodePoolStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodePoolStatusParamsWithHTTPClient(client *http.Client) *GetNodePoolStatusParams {
	var ()
	return &GetNodePoolStatusParams{
		HTTPClient: client,
	}
}

/*GetNodePoolStatusParams contains all the parameters to send to the API endpoint
for the get node pool status operation typically these are written to a http.Request
*/
type GetNodePoolStatusParams struct {

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
	/*NodepoolID
	  Node Pool ID

	*/
	NodepoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get node pool status params
func (o *GetNodePoolStatusParams) WithTimeout(timeout time.Duration) *GetNodePoolStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node pool status params
func (o *GetNodePoolStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node pool status params
func (o *GetNodePoolStatusParams) WithContext(ctx context.Context) *GetNodePoolStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node pool status params
func (o *GetNodePoolStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node pool status params
func (o *GetNodePoolStatusParams) WithHTTPClient(client *http.Client) *GetNodePoolStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node pool status params
func (o *GetNodePoolStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get node pool status params
func (o *GetNodePoolStatusParams) WithAuthorization(authorization string) *GetNodePoolStatusParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get node pool status params
func (o *GetNodePoolStatusParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the get node pool status params
func (o *GetNodePoolStatusParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *GetNodePoolStatusParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the get node pool status params
func (o *GetNodePoolStatusParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get node pool status params
func (o *GetNodePoolStatusParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *GetNodePoolStatusParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get node pool status params
func (o *GetNodePoolStatusParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the get node pool status params
func (o *GetNodePoolStatusParams) WithXRequestID(xRequestID *string) *GetNodePoolStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get node pool status params
func (o *GetNodePoolStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithClusterID adds the clusterID to the get node pool status params
func (o *GetNodePoolStatusParams) WithClusterID(clusterID string) *GetNodePoolStatusParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get node pool status params
func (o *GetNodePoolStatusParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNodepoolID adds the nodepoolID to the get node pool status params
func (o *GetNodePoolStatusParams) WithNodepoolID(nodepoolID string) *GetNodePoolStatusParams {
	o.SetNodepoolID(nodepoolID)
	return o
}

// SetNodepoolID adds the nodepoolId to the get node pool status params
func (o *GetNodePoolStatusParams) SetNodepoolID(nodepoolID string) {
	o.NodepoolID = nodepoolID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodePoolStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nodepool_id
	if err := r.SetPathParam("nodepool_id", o.NodepoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}