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

// NewJGroupAddSubscriptionParams creates a new JGroupAddSubscriptionParams object
// with the default values initialized.
func NewJGroupAddSubscriptionParams() *JGroupAddSubscriptionParams {
	var ()
	return &JGroupAddSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupAddSubscriptionParamsWithTimeout creates a new JGroupAddSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupAddSubscriptionParamsWithTimeout(timeout time.Duration) *JGroupAddSubscriptionParams {
	var ()
	return &JGroupAddSubscriptionParams{

		timeout: timeout,
	}
}

// NewJGroupAddSubscriptionParamsWithContext creates a new JGroupAddSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupAddSubscriptionParamsWithContext(ctx context.Context) *JGroupAddSubscriptionParams {
	var ()
	return &JGroupAddSubscriptionParams{

		Context: ctx,
	}
}

/*JGroupAddSubscriptionParams contains all the parameters to send to the API endpoint
for the j group add subscription operation typically these are written to a http.Request
*/
type JGroupAddSubscriptionParams struct {

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

// WithTimeout adds the timeout to the j group add subscription params
func (o *JGroupAddSubscriptionParams) WithTimeout(timeout time.Duration) *JGroupAddSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group add subscription params
func (o *JGroupAddSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group add subscription params
func (o *JGroupAddSubscriptionParams) WithContext(ctx context.Context) *JGroupAddSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group add subscription params
func (o *JGroupAddSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group add subscription params
func (o *JGroupAddSubscriptionParams) WithBody(body models.DefaultSelector) *JGroupAddSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group add subscription params
func (o *JGroupAddSubscriptionParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group add subscription params
func (o *JGroupAddSubscriptionParams) WithID(id string) *JGroupAddSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group add subscription params
func (o *JGroupAddSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupAddSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
