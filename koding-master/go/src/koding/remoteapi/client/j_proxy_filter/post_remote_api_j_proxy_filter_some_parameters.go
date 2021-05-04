package j_proxy_filter

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

// NewPostRemoteAPIJProxyFilterSomeParams creates a new PostRemoteAPIJProxyFilterSomeParams object
// with the default values initialized.
func NewPostRemoteAPIJProxyFilterSomeParams() *PostRemoteAPIJProxyFilterSomeParams {
	var ()
	return &PostRemoteAPIJProxyFilterSomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProxyFilterSomeParamsWithTimeout creates a new PostRemoteAPIJProxyFilterSomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProxyFilterSomeParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProxyFilterSomeParams {
	var ()
	return &PostRemoteAPIJProxyFilterSomeParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProxyFilterSomeParamsWithContext creates a new PostRemoteAPIJProxyFilterSomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProxyFilterSomeParamsWithContext(ctx context.Context) *PostRemoteAPIJProxyFilterSomeParams {
	var ()
	return &PostRemoteAPIJProxyFilterSomeParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProxyFilterSomeParams contains all the parameters to send to the API endpoint
for the post remote API j proxy filter some operation typically these are written to a http.Request
*/
type PostRemoteAPIJProxyFilterSomeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j proxy filter some params
func (o *PostRemoteAPIJProxyFilterSomeParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProxyFilterSomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j proxy filter some params
func (o *PostRemoteAPIJProxyFilterSomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j proxy filter some params
func (o *PostRemoteAPIJProxyFilterSomeParams) WithContext(ctx context.Context) *PostRemoteAPIJProxyFilterSomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j proxy filter some params
func (o *PostRemoteAPIJProxyFilterSomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j proxy filter some params
func (o *PostRemoteAPIJProxyFilterSomeParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJProxyFilterSomeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j proxy filter some params
func (o *PostRemoteAPIJProxyFilterSomeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProxyFilterSomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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