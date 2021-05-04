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

// NewJGroupFetchMyRolesParams creates a new JGroupFetchMyRolesParams object
// with the default values initialized.
func NewJGroupFetchMyRolesParams() *JGroupFetchMyRolesParams {
	var ()
	return &JGroupFetchMyRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupFetchMyRolesParamsWithTimeout creates a new JGroupFetchMyRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupFetchMyRolesParamsWithTimeout(timeout time.Duration) *JGroupFetchMyRolesParams {
	var ()
	return &JGroupFetchMyRolesParams{

		timeout: timeout,
	}
}

// NewJGroupFetchMyRolesParamsWithContext creates a new JGroupFetchMyRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupFetchMyRolesParamsWithContext(ctx context.Context) *JGroupFetchMyRolesParams {
	var ()
	return &JGroupFetchMyRolesParams{

		Context: ctx,
	}
}

/*JGroupFetchMyRolesParams contains all the parameters to send to the API endpoint
for the j group fetch my roles operation typically these are written to a http.Request
*/
type JGroupFetchMyRolesParams struct {

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

// WithTimeout adds the timeout to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) WithTimeout(timeout time.Duration) *JGroupFetchMyRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) WithContext(ctx context.Context) *JGroupFetchMyRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) WithBody(body models.DefaultSelector) *JGroupFetchMyRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) WithID(id string) *JGroupFetchMyRolesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group fetch my roles params
func (o *JGroupFetchMyRolesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupFetchMyRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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