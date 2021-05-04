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

// NewJMachineShareParams creates a new JMachineShareParams object
// with the default values initialized.
func NewJMachineShareParams() *JMachineShareParams {
	var ()
	return &JMachineShareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJMachineShareParamsWithTimeout creates a new JMachineShareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJMachineShareParamsWithTimeout(timeout time.Duration) *JMachineShareParams {
	var ()
	return &JMachineShareParams{

		timeout: timeout,
	}
}

// NewJMachineShareParamsWithContext creates a new JMachineShareParams object
// with the default values initialized, and the ability to set a context for a request
func NewJMachineShareParamsWithContext(ctx context.Context) *JMachineShareParams {
	var ()
	return &JMachineShareParams{

		Context: ctx,
	}
}

/*JMachineShareParams contains all the parameters to send to the API endpoint
for the j machine share operation typically these are written to a http.Request
*/
type JMachineShareParams struct {

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

// WithTimeout adds the timeout to the j machine share params
func (o *JMachineShareParams) WithTimeout(timeout time.Duration) *JMachineShareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j machine share params
func (o *JMachineShareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j machine share params
func (o *JMachineShareParams) WithContext(ctx context.Context) *JMachineShareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j machine share params
func (o *JMachineShareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j machine share params
func (o *JMachineShareParams) WithBody(body models.DefaultSelector) *JMachineShareParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j machine share params
func (o *JMachineShareParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j machine share params
func (o *JMachineShareParams) WithID(id string) *JMachineShareParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j machine share params
func (o *JMachineShareParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JMachineShareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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