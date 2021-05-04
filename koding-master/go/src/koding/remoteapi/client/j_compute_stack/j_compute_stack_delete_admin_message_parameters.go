package j_compute_stack

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

// NewJComputeStackDeleteAdminMessageParams creates a new JComputeStackDeleteAdminMessageParams object
// with the default values initialized.
func NewJComputeStackDeleteAdminMessageParams() *JComputeStackDeleteAdminMessageParams {
	var ()
	return &JComputeStackDeleteAdminMessageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJComputeStackDeleteAdminMessageParamsWithTimeout creates a new JComputeStackDeleteAdminMessageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJComputeStackDeleteAdminMessageParamsWithTimeout(timeout time.Duration) *JComputeStackDeleteAdminMessageParams {
	var ()
	return &JComputeStackDeleteAdminMessageParams{

		timeout: timeout,
	}
}

// NewJComputeStackDeleteAdminMessageParamsWithContext creates a new JComputeStackDeleteAdminMessageParams object
// with the default values initialized, and the ability to set a context for a request
func NewJComputeStackDeleteAdminMessageParamsWithContext(ctx context.Context) *JComputeStackDeleteAdminMessageParams {
	var ()
	return &JComputeStackDeleteAdminMessageParams{

		Context: ctx,
	}
}

/*JComputeStackDeleteAdminMessageParams contains all the parameters to send to the API endpoint
for the j compute stack delete admin message operation typically these are written to a http.Request
*/
type JComputeStackDeleteAdminMessageParams struct {

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

// WithTimeout adds the timeout to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) WithTimeout(timeout time.Duration) *JComputeStackDeleteAdminMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) WithContext(ctx context.Context) *JComputeStackDeleteAdminMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) WithBody(body models.DefaultSelector) *JComputeStackDeleteAdminMessageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) WithID(id string) *JComputeStackDeleteAdminMessageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j compute stack delete admin message params
func (o *JComputeStackDeleteAdminMessageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JComputeStackDeleteAdminMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
