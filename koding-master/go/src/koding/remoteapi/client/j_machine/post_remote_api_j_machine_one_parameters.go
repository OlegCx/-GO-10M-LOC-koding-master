package j_machine

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

// NewPostRemoteAPIJMachineOneParams creates a new PostRemoteAPIJMachineOneParams object
// with the default values initialized.
func NewPostRemoteAPIJMachineOneParams() *PostRemoteAPIJMachineOneParams {
	var ()
	return &PostRemoteAPIJMachineOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJMachineOneParamsWithTimeout creates a new PostRemoteAPIJMachineOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJMachineOneParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJMachineOneParams {
	var ()
	return &PostRemoteAPIJMachineOneParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJMachineOneParamsWithContext creates a new PostRemoteAPIJMachineOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJMachineOneParamsWithContext(ctx context.Context) *PostRemoteAPIJMachineOneParams {
	var ()
	return &PostRemoteAPIJMachineOneParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJMachineOneParams contains all the parameters to send to the API endpoint
for the post remote API j machine one operation typically these are written to a http.Request
*/
type PostRemoteAPIJMachineOneParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j machine one params
func (o *PostRemoteAPIJMachineOneParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJMachineOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j machine one params
func (o *PostRemoteAPIJMachineOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j machine one params
func (o *PostRemoteAPIJMachineOneParams) WithContext(ctx context.Context) *PostRemoteAPIJMachineOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j machine one params
func (o *PostRemoteAPIJMachineOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j machine one params
func (o *PostRemoteAPIJMachineOneParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJMachineOneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j machine one params
func (o *PostRemoteAPIJMachineOneParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJMachineOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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