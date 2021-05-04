package j_invitation

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

// NewJInvitationSomeParams creates a new JInvitationSomeParams object
// with the default values initialized.
func NewJInvitationSomeParams() *JInvitationSomeParams {
	var ()
	return &JInvitationSomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJInvitationSomeParamsWithTimeout creates a new JInvitationSomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJInvitationSomeParamsWithTimeout(timeout time.Duration) *JInvitationSomeParams {
	var ()
	return &JInvitationSomeParams{

		timeout: timeout,
	}
}

// NewJInvitationSomeParamsWithContext creates a new JInvitationSomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewJInvitationSomeParamsWithContext(ctx context.Context) *JInvitationSomeParams {
	var ()
	return &JInvitationSomeParams{

		Context: ctx,
	}
}

/*JInvitationSomeParams contains all the parameters to send to the API endpoint
for the j invitation some operation typically these are written to a http.Request
*/
type JInvitationSomeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j invitation some params
func (o *JInvitationSomeParams) WithTimeout(timeout time.Duration) *JInvitationSomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j invitation some params
func (o *JInvitationSomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j invitation some params
func (o *JInvitationSomeParams) WithContext(ctx context.Context) *JInvitationSomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j invitation some params
func (o *JInvitationSomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j invitation some params
func (o *JInvitationSomeParams) WithBody(body models.DefaultSelector) *JInvitationSomeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j invitation some params
func (o *JInvitationSomeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JInvitationSomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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