// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewModifyOrganizationParams creates a new ModifyOrganizationParams object
// with the default values initialized.
func NewModifyOrganizationParams() *ModifyOrganizationParams {
	var ()
	return &ModifyOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyOrganizationParamsWithTimeout creates a new ModifyOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyOrganizationParamsWithTimeout(timeout time.Duration) *ModifyOrganizationParams {
	var ()
	return &ModifyOrganizationParams{

		timeout: timeout,
	}
}

// NewModifyOrganizationParamsWithContext creates a new ModifyOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyOrganizationParamsWithContext(ctx context.Context) *ModifyOrganizationParams {
	var ()
	return &ModifyOrganizationParams{

		Context: ctx,
	}
}

// NewModifyOrganizationParamsWithHTTPClient creates a new ModifyOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyOrganizationParamsWithHTTPClient(client *http.Client) *ModifyOrganizationParams {
	var ()
	return &ModifyOrganizationParams{
		HTTPClient: client,
	}
}

/*ModifyOrganizationParams contains all the parameters to send to the API endpoint
for the modify organization operation typically these are written to a http.Request
*/
type ModifyOrganizationParams struct {

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
	/*Body*/
	Body *models.ModifyOrganizationParamsBody
	/*OrganizationID
	  An ID for the organization.
	This ID must be unique and match this regular
	expression: ^[a-z0-9_]{4,30}$


	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify organization params
func (o *ModifyOrganizationParams) WithTimeout(timeout time.Duration) *ModifyOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify organization params
func (o *ModifyOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify organization params
func (o *ModifyOrganizationParams) WithContext(ctx context.Context) *ModifyOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify organization params
func (o *ModifyOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify organization params
func (o *ModifyOrganizationParams) WithHTTPClient(client *http.Client) *ModifyOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify organization params
func (o *ModifyOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the modify organization params
func (o *ModifyOrganizationParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *ModifyOrganizationParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the modify organization params
func (o *ModifyOrganizationParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify organization params
func (o *ModifyOrganizationParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *ModifyOrganizationParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the modify organization params
func (o *ModifyOrganizationParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the modify organization params
func (o *ModifyOrganizationParams) WithXRequestID(xRequestID *string) *ModifyOrganizationParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the modify organization params
func (o *ModifyOrganizationParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the modify organization params
func (o *ModifyOrganizationParams) WithBody(body *models.ModifyOrganizationParamsBody) *ModifyOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify organization params
func (o *ModifyOrganizationParams) SetBody(body *models.ModifyOrganizationParamsBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the modify organization params
func (o *ModifyOrganizationParams) WithOrganizationID(organizationID string) *ModifyOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the modify organization params
func (o *ModifyOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
