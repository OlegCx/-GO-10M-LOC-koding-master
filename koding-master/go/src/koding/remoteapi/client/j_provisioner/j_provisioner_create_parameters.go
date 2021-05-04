package j_provisioner

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

// NewJProvisionerCreateParams creates a new JProvisionerCreateParams object
// with the default values initialized.
func NewJProvisionerCreateParams() *JProvisionerCreateParams {
	var ()
	return &JProvisionerCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJProvisionerCreateParamsWithTimeout creates a new JProvisionerCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJProvisionerCreateParamsWithTimeout(timeout time.Duration) *JProvisionerCreateParams {
	var ()
	return &JProvisionerCreateParams{

		timeout: timeout,
	}
}

// NewJProvisionerCreateParamsWithContext creates a new JProvisionerCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJProvisionerCreateParamsWithContext(ctx context.Context) *JProvisionerCreateParams {
	var ()
	return &JProvisionerCreateParams{

		Context: ctx,
	}
}

/*JProvisionerCreateParams contains all the parameters to send to the API endpoint
for the j provisioner create operation typically these are written to a http.Request
*/
type JProvisionerCreateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j provisioner create params
func (o *JProvisionerCreateParams) WithTimeout(timeout time.Duration) *JProvisionerCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j provisioner create params
func (o *JProvisionerCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j provisioner create params
func (o *JProvisionerCreateParams) WithContext(ctx context.Context) *JProvisionerCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j provisioner create params
func (o *JProvisionerCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j provisioner create params
func (o *JProvisionerCreateParams) WithBody(body models.DefaultSelector) *JProvisionerCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j provisioner create params
func (o *JProvisionerCreateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JProvisionerCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}