package j_stack_template

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

// NewJStackTemplateUpdateParams creates a new JStackTemplateUpdateParams object
// with the default values initialized.
func NewJStackTemplateUpdateParams() *JStackTemplateUpdateParams {
	var ()
	return &JStackTemplateUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJStackTemplateUpdateParamsWithTimeout creates a new JStackTemplateUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJStackTemplateUpdateParamsWithTimeout(timeout time.Duration) *JStackTemplateUpdateParams {
	var ()
	return &JStackTemplateUpdateParams{

		timeout: timeout,
	}
}

// NewJStackTemplateUpdateParamsWithContext creates a new JStackTemplateUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJStackTemplateUpdateParamsWithContext(ctx context.Context) *JStackTemplateUpdateParams {
	var ()
	return &JStackTemplateUpdateParams{

		Context: ctx,
	}
}

/*JStackTemplateUpdateParams contains all the parameters to send to the API endpoint
for the j stack template update operation typically these are written to a http.Request
*/
type JStackTemplateUpdateParams struct {

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

// WithTimeout adds the timeout to the j stack template update params
func (o *JStackTemplateUpdateParams) WithTimeout(timeout time.Duration) *JStackTemplateUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j stack template update params
func (o *JStackTemplateUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j stack template update params
func (o *JStackTemplateUpdateParams) WithContext(ctx context.Context) *JStackTemplateUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j stack template update params
func (o *JStackTemplateUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j stack template update params
func (o *JStackTemplateUpdateParams) WithBody(body models.DefaultSelector) *JStackTemplateUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j stack template update params
func (o *JStackTemplateUpdateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j stack template update params
func (o *JStackTemplateUpdateParams) WithID(id string) *JStackTemplateUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j stack template update params
func (o *JStackTemplateUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JStackTemplateUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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