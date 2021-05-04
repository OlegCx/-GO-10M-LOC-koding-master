package j_group

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

// NewJGroupCreateParams creates a new JGroupCreateParams object
// with the default values initialized.
func NewJGroupCreateParams() *JGroupCreateParams {
	var ()
	return &JGroupCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupCreateParamsWithTimeout creates a new JGroupCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupCreateParamsWithTimeout(timeout time.Duration) *JGroupCreateParams {
	var ()
	return &JGroupCreateParams{

		timeout: timeout,
	}
}

// NewJGroupCreateParamsWithContext creates a new JGroupCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupCreateParamsWithContext(ctx context.Context) *JGroupCreateParams {
	var ()
	return &JGroupCreateParams{

		Context: ctx,
	}
}

/*JGroupCreateParams contains all the parameters to send to the API endpoint
for the j group create operation typically these are written to a http.Request
*/
type JGroupCreateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j group create params
func (o *JGroupCreateParams) WithTimeout(timeout time.Duration) *JGroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group create params
func (o *JGroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group create params
func (o *JGroupCreateParams) WithContext(ctx context.Context) *JGroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group create params
func (o *JGroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group create params
func (o *JGroupCreateParams) WithBody(body models.DefaultSelector) *JGroupCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group create params
func (o *JGroupCreateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
