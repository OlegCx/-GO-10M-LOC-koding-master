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

// NewPostRemoteAPIJAccountSomeParams creates a new PostRemoteAPIJAccountSomeParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountSomeParams() *PostRemoteAPIJAccountSomeParams {
	var ()
	return &PostRemoteAPIJAccountSomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountSomeParamsWithTimeout creates a new PostRemoteAPIJAccountSomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountSomeParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountSomeParams {
	var ()
	return &PostRemoteAPIJAccountSomeParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountSomeParamsWithContext creates a new PostRemoteAPIJAccountSomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountSomeParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountSomeParams {
	var ()
	return &PostRemoteAPIJAccountSomeParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountSomeParams contains all the parameters to send to the API endpoint
for the post remote API j account some operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountSomeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j account some params
func (o *PostRemoteAPIJAccountSomeParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountSomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account some params
func (o *PostRemoteAPIJAccountSomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account some params
func (o *PostRemoteAPIJAccountSomeParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountSomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account some params
func (o *PostRemoteAPIJAccountSomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j account some params
func (o *PostRemoteAPIJAccountSomeParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJAccountSomeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j account some params
func (o *PostRemoteAPIJAccountSomeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountSomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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