// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewDeleteUserParams creates a new DeleteUserParams object
// with the default values initialized.
func NewDeleteUserParams() *DeleteUserParams {
	var ()
	return &DeleteUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserParamsWithTimeout creates a new DeleteUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserParamsWithTimeout(timeout time.Duration) *DeleteUserParams {
	var ()
	return &DeleteUserParams{

		timeout: timeout,
	}
}

// NewDeleteUserParamsWithContext creates a new DeleteUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserParamsWithContext(ctx context.Context) *DeleteUserParams {
	var ()
	return &DeleteUserParams{

		Context: ctx,
	}
}

// NewDeleteUserParamsWithHTTPClient creates a new DeleteUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserParamsWithHTTPClient(client *http.Client) *DeleteUserParams {
	var ()
	return &DeleteUserParams{
		HTTPClient: client,
	}
}

/*DeleteUserParams contains all the parameters to send to the API endpoint
for the delete user operation typically these are written to a http.Request
*/
type DeleteUserParams struct {

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
	/*Email
	  The user's email address

	*/
	Email string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user params
func (o *DeleteUserParams) WithTimeout(timeout time.Duration) *DeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user params
func (o *DeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user params
func (o *DeleteUserParams) WithContext(ctx context.Context) *DeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user params
func (o *DeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user params
func (o *DeleteUserParams) WithHTTPClient(client *http.Client) *DeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user params
func (o *DeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete user params
func (o *DeleteUserParams) WithAuthorization(authorization string) *DeleteUserParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete user params
func (o *DeleteUserParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the delete user params
func (o *DeleteUserParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *DeleteUserParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the delete user params
func (o *DeleteUserParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete user params
func (o *DeleteUserParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *DeleteUserParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the delete user params
func (o *DeleteUserParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the delete user params
func (o *DeleteUserParams) WithXRequestID(xRequestID *string) *DeleteUserParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete user params
func (o *DeleteUserParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEmail adds the email to the delete user params
func (o *DeleteUserParams) WithEmail(email string) *DeleteUserParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the delete user params
func (o *DeleteUserParams) SetEmail(email string) {
	o.Email = email
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
