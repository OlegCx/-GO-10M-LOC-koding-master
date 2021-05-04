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

// NewJAccountFetchMyPermissionsParams creates a new JAccountFetchMyPermissionsParams object
// with the default values initialized.
func NewJAccountFetchMyPermissionsParams() *JAccountFetchMyPermissionsParams {
	var ()
	return &JAccountFetchMyPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountFetchMyPermissionsParamsWithTimeout creates a new JAccountFetchMyPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountFetchMyPermissionsParamsWithTimeout(timeout time.Duration) *JAccountFetchMyPermissionsParams {
	var ()
	return &JAccountFetchMyPermissionsParams{

		timeout: timeout,
	}
}

// NewJAccountFetchMyPermissionsParamsWithContext creates a new JAccountFetchMyPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountFetchMyPermissionsParamsWithContext(ctx context.Context) *JAccountFetchMyPermissionsParams {
	var ()
	return &JAccountFetchMyPermissionsParams{

		Context: ctx,
	}
}

/*JAccountFetchMyPermissionsParams contains all the parameters to send to the API endpoint
for the j account fetch my permissions operation typically these are written to a http.Request
*/
type JAccountFetchMyPermissionsParams struct {

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

// WithTimeout adds the timeout to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) WithTimeout(timeout time.Duration) *JAccountFetchMyPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) WithContext(ctx context.Context) *JAccountFetchMyPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) WithBody(body models.DefaultSelector) *JAccountFetchMyPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) WithID(id string) *JAccountFetchMyPermissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j account fetch my permissions params
func (o *JAccountFetchMyPermissionsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountFetchMyPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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