package kloud

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

// NewKloudBootstrapParams creates a new KloudBootstrapParams object
// with the default values initialized.
func NewKloudBootstrapParams() *KloudBootstrapParams {
	var ()
	return &KloudBootstrapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKloudBootstrapParamsWithTimeout creates a new KloudBootstrapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKloudBootstrapParamsWithTimeout(timeout time.Duration) *KloudBootstrapParams {
	var ()
	return &KloudBootstrapParams{

		timeout: timeout,
	}
}

// NewKloudBootstrapParamsWithContext creates a new KloudBootstrapParams object
// with the default values initialized, and the ability to set a context for a request
func NewKloudBootstrapParamsWithContext(ctx context.Context) *KloudBootstrapParams {
	var ()
	return &KloudBootstrapParams{

		Context: ctx,
	}
}

/*KloudBootstrapParams contains all the parameters to send to the API endpoint
for the kloud bootstrap operation typically these are written to a http.Request
*/
type KloudBootstrapParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the kloud bootstrap params
func (o *KloudBootstrapParams) WithTimeout(timeout time.Duration) *KloudBootstrapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kloud bootstrap params
func (o *KloudBootstrapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kloud bootstrap params
func (o *KloudBootstrapParams) WithContext(ctx context.Context) *KloudBootstrapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kloud bootstrap params
func (o *KloudBootstrapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the kloud bootstrap params
func (o *KloudBootstrapParams) WithBody(body models.DefaultSelector) *KloudBootstrapParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kloud bootstrap params
func (o *KloudBootstrapParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KloudBootstrapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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