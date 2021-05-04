package j_machine

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

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJMachineReviveUsersIDParams creates a new PostRemoteAPIJMachineReviveUsersIDParams object
// with the default values initialized.
func NewPostRemoteAPIJMachineReviveUsersIDParams() *PostRemoteAPIJMachineReviveUsersIDParams {
	var ()
	return &PostRemoteAPIJMachineReviveUsersIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJMachineReviveUsersIDParamsWithTimeout creates a new PostRemoteAPIJMachineReviveUsersIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJMachineReviveUsersIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJMachineReviveUsersIDParams {
	var ()
	return &PostRemoteAPIJMachineReviveUsersIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJMachineReviveUsersIDParamsWithContext creates a new PostRemoteAPIJMachineReviveUsersIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJMachineReviveUsersIDParamsWithContext(ctx context.Context) *PostRemoteAPIJMachineReviveUsersIDParams {
	var ()
	return &PostRemoteAPIJMachineReviveUsersIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJMachineReviveUsersIDParams contains all the parameters to send to the API endpoint
for the post remote API j machine revive users ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJMachineReviveUsersIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJMachineReviveUsersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) WithContext(ctx context.Context) *PostRemoteAPIJMachineReviveUsersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJMachineReviveUsersIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) WithID(id string) *PostRemoteAPIJMachineReviveUsersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j machine revive users ID params
func (o *PostRemoteAPIJMachineReviveUsersIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJMachineReviveUsersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
