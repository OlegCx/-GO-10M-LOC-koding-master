package compute_provider

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

// NewComputeProviderCreateParams creates a new ComputeProviderCreateParams object
// with the default values initialized.
func NewComputeProviderCreateParams() *ComputeProviderCreateParams {
	var ()
	return &ComputeProviderCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewComputeProviderCreateParamsWithTimeout creates a new ComputeProviderCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewComputeProviderCreateParamsWithTimeout(timeout time.Duration) *ComputeProviderCreateParams {
	var ()
	return &ComputeProviderCreateParams{

		timeout: timeout,
	}
}

// NewComputeProviderCreateParamsWithContext creates a new ComputeProviderCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewComputeProviderCreateParamsWithContext(ctx context.Context) *ComputeProviderCreateParams {
	var ()
	return &ComputeProviderCreateParams{

		Context: ctx,
	}
}

/*ComputeProviderCreateParams contains all the parameters to send to the API endpoint
for the compute provider create operation typically these are written to a http.Request
*/
type ComputeProviderCreateParams struct {

	/*Body
	  body of the request

	*/
	Body ComputeProviderCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the compute provider create params
func (o *ComputeProviderCreateParams) WithTimeout(timeout time.Duration) *ComputeProviderCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compute provider create params
func (o *ComputeProviderCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the compute provider create params
func (o *ComputeProviderCreateParams) WithContext(ctx context.Context) *ComputeProviderCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compute provider create params
func (o *ComputeProviderCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the compute provider create params
func (o *ComputeProviderCreateParams) WithBody(body ComputeProviderCreateBody) *ComputeProviderCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the compute provider create params
func (o *ComputeProviderCreateParams) SetBody(body ComputeProviderCreateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ComputeProviderCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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