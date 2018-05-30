// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/giantswarm/gsclientgen/models"
)

// NewModifyClusterParams creates a new ModifyClusterParams object
// with the default values initialized.
func NewModifyClusterParams() *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyClusterParamsWithTimeout creates a new ModifyClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyClusterParamsWithTimeout(timeout time.Duration) *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{

		timeout: timeout,
	}
}

// NewModifyClusterParamsWithContext creates a new ModifyClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyClusterParamsWithContext(ctx context.Context) *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{

		Context: ctx,
	}
}

// NewModifyClusterParamsWithHTTPClient creates a new ModifyClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyClusterParamsWithHTTPClient(client *http.Client) *ModifyClusterParams {
	var ()
	return &ModifyClusterParams{
		HTTPClient: client,
	}
}

/*ModifyClusterParams contains all the parameters to send to the API endpoint
for the modify cluster operation typically these are written to a http.Request
*/
type ModifyClusterParams struct {

	/*Authorization
	  Header to pass an authorization token. The value has to be in the form
	`giantswarm <token>`.


	*/
	Authorization string
	/*XGiantSwarmActivity
	  Name of an activity to track, like "list-clusters"

	*/
	XGiantSwarmActivity *string
	/*XGiantSwarmCmdLine
	  If activity has been issued by a CLI, this header can contain the
	command line


	*/
	XGiantSwarmCmdLine *string
	/*XRequestID
	  A randomly generated key that can be used to track a request throughout
	services of Giant Swarm


	*/
	XRequestID *string
	/*Body
	  Modified cluster definition (JSON merge-patch)

	*/
	Body *models.V4ModifyClusterRequest
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify cluster params
func (o *ModifyClusterParams) WithTimeout(timeout time.Duration) *ModifyClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify cluster params
func (o *ModifyClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify cluster params
func (o *ModifyClusterParams) WithContext(ctx context.Context) *ModifyClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify cluster params
func (o *ModifyClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify cluster params
func (o *ModifyClusterParams) WithHTTPClient(client *http.Client) *ModifyClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify cluster params
func (o *ModifyClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the modify cluster params
func (o *ModifyClusterParams) WithAuthorization(authorization string) *ModifyClusterParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the modify cluster params
func (o *ModifyClusterParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the modify cluster params
func (o *ModifyClusterParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *ModifyClusterParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the modify cluster params
func (o *ModifyClusterParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify cluster params
func (o *ModifyClusterParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *ModifyClusterParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify cluster params
func (o *ModifyClusterParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the modify cluster params
func (o *ModifyClusterParams) WithXRequestID(xRequestID *string) *ModifyClusterParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the modify cluster params
func (o *ModifyClusterParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the modify cluster params
func (o *ModifyClusterParams) WithBody(body *models.V4ModifyClusterRequest) *ModifyClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify cluster params
func (o *ModifyClusterParams) SetBody(body *models.V4ModifyClusterRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the modify cluster params
func (o *ModifyClusterParams) WithClusterID(clusterID string) *ModifyClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the modify cluster params
func (o *ModifyClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
