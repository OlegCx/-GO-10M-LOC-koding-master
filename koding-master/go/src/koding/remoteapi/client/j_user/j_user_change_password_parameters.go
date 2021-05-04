package j_user

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

// NewJUserChangePasswordParams creates a new JUserChangePasswordParams object
// with the default values initialized.
func NewJUserChangePasswordParams() *JUserChangePasswordParams {
	var ()
	return &JUserChangePasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJUserChangePasswordParamsWithTimeout creates a new JUserChangePasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJUserChangePasswordParamsWithTimeout(timeout time.Duration) *JUserChangePasswordParams {
	var ()
	return &JUserChangePasswordParams{

		timeout: timeout,
	}
}

// NewJUserChangePasswordParamsWithContext creates a new JUserChangePasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewJUserChangePasswordParamsWithContext(ctx context.Context) *JUserChangePasswordParams {
	var ()
	return &JUserChangePasswordParams{

		Context: ctx,
	}
}

/*JUserChangePasswordParams contains all the parameters to send to the API endpoint
for the j user change password operation typically these are written to a http.Request
*/
type JUserChangePasswordParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j user change password params
func (o *JUserChangePasswordParams) WithTimeout(timeout time.Duration) *JUserChangePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j user change password params
func (o *JUserChangePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j user change password params
func (o *JUserChangePasswordParams) WithContext(ctx context.Context) *JUserChangePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j user change password params
func (o *JUserChangePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j user change password params
func (o *JUserChangePasswordParams) WithBody(body models.DefaultSelector) *JUserChangePasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j user change password params
func (o *JUserChangePasswordParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JUserChangePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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