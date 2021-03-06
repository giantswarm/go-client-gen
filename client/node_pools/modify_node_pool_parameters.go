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

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// NewModifyNodePoolParams creates a new ModifyNodePoolParams object
// with the default values initialized.
func NewModifyNodePoolParams() *ModifyNodePoolParams {
	var ()
	return &ModifyNodePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyNodePoolParamsWithTimeout creates a new ModifyNodePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyNodePoolParamsWithTimeout(timeout time.Duration) *ModifyNodePoolParams {
	var ()
	return &ModifyNodePoolParams{

		timeout: timeout,
	}
}

// NewModifyNodePoolParamsWithContext creates a new ModifyNodePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyNodePoolParamsWithContext(ctx context.Context) *ModifyNodePoolParams {
	var ()
	return &ModifyNodePoolParams{

		Context: ctx,
	}
}

// NewModifyNodePoolParamsWithHTTPClient creates a new ModifyNodePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyNodePoolParamsWithHTTPClient(client *http.Client) *ModifyNodePoolParams {
	var ()
	return &ModifyNodePoolParams{
		HTTPClient: client,
	}
}

/*ModifyNodePoolParams contains all the parameters to send to the API endpoint
for the modify node pool operation typically these are written to a http.Request
*/
type ModifyNodePoolParams struct {

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
	/*Body
	  Merge-patch body

	*/
	Body *models.V5ModifyNodePoolRequest
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

// WithTimeout adds the timeout to the modify node pool params
func (o *ModifyNodePoolParams) WithTimeout(timeout time.Duration) *ModifyNodePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify node pool params
func (o *ModifyNodePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify node pool params
func (o *ModifyNodePoolParams) WithContext(ctx context.Context) *ModifyNodePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify node pool params
func (o *ModifyNodePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify node pool params
func (o *ModifyNodePoolParams) WithHTTPClient(client *http.Client) *ModifyNodePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify node pool params
func (o *ModifyNodePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the modify node pool params
func (o *ModifyNodePoolParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *ModifyNodePoolParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the modify node pool params
func (o *ModifyNodePoolParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify node pool params
func (o *ModifyNodePoolParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *ModifyNodePoolParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify node pool params
func (o *ModifyNodePoolParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the modify node pool params
func (o *ModifyNodePoolParams) WithXRequestID(xRequestID *string) *ModifyNodePoolParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the modify node pool params
func (o *ModifyNodePoolParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the modify node pool params
func (o *ModifyNodePoolParams) WithBody(body *models.V5ModifyNodePoolRequest) *ModifyNodePoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify node pool params
func (o *ModifyNodePoolParams) SetBody(body *models.V5ModifyNodePoolRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the modify node pool params
func (o *ModifyNodePoolParams) WithClusterID(clusterID string) *ModifyNodePoolParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the modify node pool params
func (o *ModifyNodePoolParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNodepoolID adds the nodepoolID to the modify node pool params
func (o *ModifyNodePoolParams) WithNodepoolID(nodepoolID string) *ModifyNodePoolParams {
	o.SetNodepoolID(nodepoolID)
	return o
}

// SetNodepoolID adds the nodepoolId to the modify node pool params
func (o *ModifyNodePoolParams) SetNodepoolID(nodepoolID string) {
	o.NodepoolID = nodepoolID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyNodePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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
