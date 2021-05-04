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

// NewJGroupToggleFeatureParams creates a new JGroupToggleFeatureParams object
// with the default values initialized.
func NewJGroupToggleFeatureParams() *JGroupToggleFeatureParams {
	var ()
	return &JGroupToggleFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupToggleFeatureParamsWithTimeout creates a new JGroupToggleFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupToggleFeatureParamsWithTimeout(timeout time.Duration) *JGroupToggleFeatureParams {
	var ()
	return &JGroupToggleFeatureParams{

		timeout: timeout,
	}
}

// NewJGroupToggleFeatureParamsWithContext creates a new JGroupToggleFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupToggleFeatureParamsWithContext(ctx context.Context) *JGroupToggleFeatureParams {
	var ()
	return &JGroupToggleFeatureParams{

		Context: ctx,
	}
}

/*JGroupToggleFeatureParams contains all the parameters to send to the API endpoint
for the j group toggle feature operation typically these are written to a http.Request
*/
type JGroupToggleFeatureParams struct {

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

// WithTimeout adds the timeout to the j group toggle feature params
func (o *JGroupToggleFeatureParams) WithTimeout(timeout time.Duration) *JGroupToggleFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group toggle feature params
func (o *JGroupToggleFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group toggle feature params
func (o *JGroupToggleFeatureParams) WithContext(ctx context.Context) *JGroupToggleFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group toggle feature params
func (o *JGroupToggleFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group toggle feature params
func (o *JGroupToggleFeatureParams) WithBody(body models.DefaultSelector) *JGroupToggleFeatureParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group toggle feature params
func (o *JGroupToggleFeatureParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group toggle feature params
func (o *JGroupToggleFeatureParams) WithID(id string) *JGroupToggleFeatureParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group toggle feature params
func (o *JGroupToggleFeatureParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupToggleFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
