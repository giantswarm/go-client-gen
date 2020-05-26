// Code generated by go-swagger; DO NOT EDIT.

package node_pools

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

// NewDeleteNodePoolParams creates a new DeleteNodePoolParams object
// with the default values initialized.
func NewDeleteNodePoolParams() *DeleteNodePoolParams {
	var ()
	return &DeleteNodePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodePoolParamsWithTimeout creates a new DeleteNodePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNodePoolParamsWithTimeout(timeout time.Duration) *DeleteNodePoolParams {
	var ()
	return &DeleteNodePoolParams{

		timeout: timeout,
	}
}

// NewDeleteNodePoolParamsWithContext creates a new DeleteNodePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNodePoolParamsWithContext(ctx context.Context) *DeleteNodePoolParams {
	var ()
	return &DeleteNodePoolParams{

		Context: ctx,
	}
}

// NewDeleteNodePoolParamsWithHTTPClient creates a new DeleteNodePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNodePoolParamsWithHTTPClient(client *http.Client) *DeleteNodePoolParams {
	var ()
	return &DeleteNodePoolParams{
		HTTPClient: client,
	}
}

/*DeleteNodePoolParams contains all the parameters to send to the API endpoint
for the delete node pool operation typically these are written to a http.Request
*/
type DeleteNodePoolParams struct {

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

// WithTimeout adds the timeout to the delete node pool params
func (o *DeleteNodePoolParams) WithTimeout(timeout time.Duration) *DeleteNodePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete node pool params
func (o *DeleteNodePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete node pool params
func (o *DeleteNodePoolParams) WithContext(ctx context.Context) *DeleteNodePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete node pool params
func (o *DeleteNodePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete node pool params
func (o *DeleteNodePoolParams) WithHTTPClient(client *http.Client) *DeleteNodePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete node pool params
func (o *DeleteNodePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the delete node pool params
func (o *DeleteNodePoolParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *DeleteNodePoolParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the delete node pool params
func (o *DeleteNodePoolParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete node pool params
func (o *DeleteNodePoolParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *DeleteNodePoolParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete node pool params
func (o *DeleteNodePoolParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the delete node pool params
func (o *DeleteNodePoolParams) WithXRequestID(xRequestID *string) *DeleteNodePoolParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete node pool params
func (o *DeleteNodePoolParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithClusterID adds the clusterID to the delete node pool params
func (o *DeleteNodePoolParams) WithClusterID(clusterID string) *DeleteNodePoolParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete node pool params
func (o *DeleteNodePoolParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNodepoolID adds the nodepoolID to the delete node pool params
func (o *DeleteNodePoolParams) WithNodepoolID(nodepoolID string) *DeleteNodePoolParams {
	o.SetNodepoolID(nodepoolID)
	return o
}

// SetNodepoolID adds the nodepoolId to the delete node pool params
func (o *DeleteNodePoolParams) SetNodepoolID(nodepoolID string) {
	o.NodepoolID = nodepoolID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nodepool_id
	if err := r.SetPathParam("nodepool_id", o.NodepoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
