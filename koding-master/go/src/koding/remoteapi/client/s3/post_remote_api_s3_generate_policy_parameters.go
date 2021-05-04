package s3

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
)

// NewPostRemoteAPIS3GeneratePolicyParams creates a new PostRemoteAPIS3GeneratePolicyParams object
// with the default values initialized.
func NewPostRemoteAPIS3GeneratePolicyParams() *PostRemoteAPIS3GeneratePolicyParams {

	return &PostRemoteAPIS3GeneratePolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIS3GeneratePolicyParamsWithTimeout creates a new PostRemoteAPIS3GeneratePolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIS3GeneratePolicyParamsWithTimeout(timeout time.Duration) *PostRemoteAPIS3GeneratePolicyParams {

	return &PostRemoteAPIS3GeneratePolicyParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIS3GeneratePolicyParamsWithContext creates a new PostRemoteAPIS3GeneratePolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIS3GeneratePolicyParamsWithContext(ctx context.Context) *PostRemoteAPIS3GeneratePolicyParams {

	return &PostRemoteAPIS3GeneratePolicyParams{

		Context: ctx,
	}
}

/*PostRemoteAPIS3GeneratePolicyParams contains all the parameters to send to the API endpoint
for the post remote API s3 generate policy operation typically these are written to a http.Request
*/
type PostRemoteAPIS3GeneratePolicyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API s3 generate policy params
func (o *PostRemoteAPIS3GeneratePolicyParams) WithTimeout(timeout time.Duration) *PostRemoteAPIS3GeneratePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API s3 generate policy params
func (o *PostRemoteAPIS3GeneratePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API s3 generate policy params
func (o *PostRemoteAPIS3GeneratePolicyParams) WithContext(ctx context.Context) *PostRemoteAPIS3GeneratePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API s3 generate policy params
func (o *PostRemoteAPIS3GeneratePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIS3GeneratePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}