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

// NewJGroupFetchMembersParams creates a new JGroupFetchMembersParams object
// with the default values initialized.
func NewJGroupFetchMembersParams() *JGroupFetchMembersParams {
	var ()
	return &JGroupFetchMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupFetchMembersParamsWithTimeout creates a new JGroupFetchMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupFetchMembersParamsWithTimeout(timeout time.Duration) *JGroupFetchMembersParams {
	var ()
	return &JGroupFetchMembersParams{

		timeout: timeout,
	}
}

// NewJGroupFetchMembersParamsWithContext creates a new JGroupFetchMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupFetchMembersParamsWithContext(ctx context.Context) *JGroupFetchMembersParams {
	var ()
	return &JGroupFetchMembersParams{

		Context: ctx,
	}
}

/*JGroupFetchMembersParams contains all the parameters to send to the API endpoint
for the j group fetch members operation typically these are written to a http.Request
*/
type JGroupFetchMembersParams struct {

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

// WithTimeout adds the timeout to the j group fetch members params
func (o *JGroupFetchMembersParams) WithTimeout(timeout time.Duration) *JGroupFetchMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group fetch members params
func (o *JGroupFetchMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group fetch members params
func (o *JGroupFetchMembersParams) WithContext(ctx context.Context) *JGroupFetchMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group fetch members params
func (o *JGroupFetchMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group fetch members params
func (o *JGroupFetchMembersParams) WithBody(body models.DefaultSelector) *JGroupFetchMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group fetch members params
func (o *JGroupFetchMembersParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group fetch members params
func (o *JGroupFetchMembersParams) WithID(id string) *JGroupFetchMembersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group fetch members params
func (o *JGroupFetchMembersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupFetchMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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