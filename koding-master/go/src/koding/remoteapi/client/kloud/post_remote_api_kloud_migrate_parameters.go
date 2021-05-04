package kloud

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

// NewPostRemoteAPIKloudMigrateParams creates a new PostRemoteAPIKloudMigrateParams object
// with the default values initialized.
func NewPostRemoteAPIKloudMigrateParams() *PostRemoteAPIKloudMigrateParams {
	var ()
	return &PostRemoteAPIKloudMigrateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIKloudMigrateParamsWithTimeout creates a new PostRemoteAPIKloudMigrateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIKloudMigrateParamsWithTimeout(timeout time.Duration) *PostRemoteAPIKloudMigrateParams {
	var ()
	return &PostRemoteAPIKloudMigrateParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIKloudMigrateParamsWithContext creates a new PostRemoteAPIKloudMigrateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIKloudMigrateParamsWithContext(ctx context.Context) *PostRemoteAPIKloudMigrateParams {
	var ()
	return &PostRemoteAPIKloudMigrateParams{

		Context: ctx,
	}
}

/*PostRemoteAPIKloudMigrateParams contains all the parameters to send to the API endpoint
for the post remote API kloud migrate operation typically these are written to a http.Request
*/
type PostRemoteAPIKloudMigrateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API kloud migrate params
func (o *PostRemoteAPIKloudMigrateParams) WithTimeout(timeout time.Duration) *PostRemoteAPIKloudMigrateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API kloud migrate params
func (o *PostRemoteAPIKloudMigrateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API kloud migrate params
func (o *PostRemoteAPIKloudMigrateParams) WithContext(ctx context.Context) *PostRemoteAPIKloudMigrateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API kloud migrate params
func (o *PostRemoteAPIKloudMigrateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API kloud migrate params
func (o *PostRemoteAPIKloudMigrateParams) WithBody(body models.DefaultSelector) *PostRemoteAPIKloudMigrateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API kloud migrate params
func (o *PostRemoteAPIKloudMigrateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIKloudMigrateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
