// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewGetClustersParams creates a new GetClustersParams object
// with the default values initialized.
func NewGetClustersParams() *GetClustersParams {
	var ()
	return &GetClustersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersParamsWithTimeout creates a new GetClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClustersParamsWithTimeout(timeout time.Duration) *GetClustersParams {
	var ()
	return &GetClustersParams{

		timeout: timeout,
	}
}

// NewGetClustersParamsWithContext creates a new GetClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClustersParamsWithContext(ctx context.Context) *GetClustersParams {
	var ()
	return &GetClustersParams{

		Context: ctx,
	}
}

// NewGetClustersParamsWithHTTPClient creates a new GetClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClustersParamsWithHTTPClient(client *http.Client) *GetClustersParams {
	var ()
	return &GetClustersParams{
		HTTPClient: client,
	}
}

/*GetClustersParams contains all the parameters to send to the API endpoint
for the get clusters operation typically these are written to a http.Request
*/
type GetClustersParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) WithTimeout(timeout time.Duration) *GetClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters params
func (o *GetClustersParams) WithContext(ctx context.Context) *GetClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters params
func (o *GetClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) WithHTTPClient(client *http.Client) *GetClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get clusters params
func (o *GetClustersParams) WithAuthorization(authorization string) *GetClustersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get clusters params
func (o *GetClustersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the get clusters params
func (o *GetClustersParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *GetClustersParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the get clusters params
func (o *GetClustersParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get clusters params
func (o *GetClustersParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *GetClustersParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get clusters params
func (o *GetClustersParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the get clusters params
func (o *GetClustersParams) WithXRequestID(xRequestID *string) *GetClustersParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get clusters params
func (o *GetClustersParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
