package j_account

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

// NewJAccountCheckGroupMembershipParams creates a new JAccountCheckGroupMembershipParams object
// with the default values initialized.
func NewJAccountCheckGroupMembershipParams() *JAccountCheckGroupMembershipParams {
	var ()
	return &JAccountCheckGroupMembershipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountCheckGroupMembershipParamsWithTimeout creates a new JAccountCheckGroupMembershipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountCheckGroupMembershipParamsWithTimeout(timeout time.Duration) *JAccountCheckGroupMembershipParams {
	var ()
	return &JAccountCheckGroupMembershipParams{

		timeout: timeout,
	}
}

// NewJAccountCheckGroupMembershipParamsWithContext creates a new JAccountCheckGroupMembershipParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountCheckGroupMembershipParamsWithContext(ctx context.Context) *JAccountCheckGroupMembershipParams {
	var ()
	return &JAccountCheckGroupMembershipParams{

		Context: ctx,
	}
}

/*JAccountCheckGroupMembershipParams contains all the parameters to send to the API endpoint
for the j account check group membership operation typically these are written to a http.Request
*/
type JAccountCheckGroupMembershipParams struct {

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

// WithTimeout adds the timeout to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) WithTimeout(timeout time.Duration) *JAccountCheckGroupMembershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) WithContext(ctx context.Context) *JAccountCheckGroupMembershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) WithBody(body models.DefaultSelector) *JAccountCheckGroupMembershipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) WithID(id string) *JAccountCheckGroupMembershipParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j account check group membership params
func (o *JAccountCheckGroupMembershipParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountCheckGroupMembershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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