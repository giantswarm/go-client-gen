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

	models "github.com/giantswarm/gsclientgen/models"
)

// NewCreateClusterAppParams creates a new CreateClusterAppParams object
// with the default values initialized.
func NewCreateClusterAppParams() *CreateClusterAppParams {
	var ()
	return &CreateClusterAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterAppParamsWithTimeout creates a new CreateClusterAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterAppParamsWithTimeout(timeout time.Duration) *CreateClusterAppParams {
	var ()
	return &CreateClusterAppParams{

		timeout: timeout,
	}
}

// NewCreateClusterAppParamsWithContext creates a new CreateClusterAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterAppParamsWithContext(ctx context.Context) *CreateClusterAppParams {
	var ()
	return &CreateClusterAppParams{

		Context: ctx,
	}
}

// NewCreateClusterAppParamsWithHTTPClient creates a new CreateClusterAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterAppParamsWithHTTPClient(client *http.Client) *CreateClusterAppParams {
	var ()
	return &CreateClusterAppParams{
		HTTPClient: client,
	}
}

/*CreateClusterAppParams contains all the parameters to send to the API endpoint
for the create cluster app operation typically these are written to a http.Request
*/
type CreateClusterAppParams struct {

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
	/*Body*/
	Body *models.V4CreateAppRequest
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster app params
func (o *CreateClusterAppParams) WithTimeout(timeout time.Duration) *CreateClusterAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster app params
func (o *CreateClusterAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster app params
func (o *CreateClusterAppParams) WithContext(ctx context.Context) *CreateClusterAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster app params
func (o *CreateClusterAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster app params
func (o *CreateClusterAppParams) WithHTTPClient(client *http.Client) *CreateClusterAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster app params
func (o *CreateClusterAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create cluster app params
func (o *CreateClusterAppParams) WithAuthorization(authorization string) *CreateClusterAppParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create cluster app params
func (o *CreateClusterAppParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the create cluster app params
func (o *CreateClusterAppParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *CreateClusterAppParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the create cluster app params
func (o *CreateClusterAppParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the create cluster app params
func (o *CreateClusterAppParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *CreateClusterAppParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the create cluster app params
func (o *CreateClusterAppParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the create cluster app params
func (o *CreateClusterAppParams) WithXRequestID(xRequestID *string) *CreateClusterAppParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create cluster app params
func (o *CreateClusterAppParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAppName adds the appName to the create cluster app params
func (o *CreateClusterAppParams) WithAppName(appName string) *CreateClusterAppParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the create cluster app params
func (o *CreateClusterAppParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the create cluster app params
func (o *CreateClusterAppParams) WithBody(body *models.V4CreateAppRequest) *CreateClusterAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster app params
func (o *CreateClusterAppParams) SetBody(body *models.V4CreateAppRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create cluster app params
func (o *CreateClusterAppParams) WithClusterID(clusterID string) *CreateClusterAppParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create cluster app params
func (o *CreateClusterAppParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
