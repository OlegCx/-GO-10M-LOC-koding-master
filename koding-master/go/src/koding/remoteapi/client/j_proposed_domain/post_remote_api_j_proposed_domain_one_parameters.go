package j_proposed_domain

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

// NewPostRemoteAPIJProposedDomainOneParams creates a new PostRemoteAPIJProposedDomainOneParams object
// with the default values initialized.
func NewPostRemoteAPIJProposedDomainOneParams() *PostRemoteAPIJProposedDomainOneParams {
	var ()
	return &PostRemoteAPIJProposedDomainOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProposedDomainOneParamsWithTimeout creates a new PostRemoteAPIJProposedDomainOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProposedDomainOneParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProposedDomainOneParams {
	var ()
	return &PostRemoteAPIJProposedDomainOneParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProposedDomainOneParamsWithContext creates a new PostRemoteAPIJProposedDomainOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProposedDomainOneParamsWithContext(ctx context.Context) *PostRemoteAPIJProposedDomainOneParams {
	var ()
	return &PostRemoteAPIJProposedDomainOneParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProposedDomainOneParams contains all the parameters to send to the API endpoint
for the post remote API j proposed domain one operation typically these are written to a http.Request
*/
type PostRemoteAPIJProposedDomainOneParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j proposed domain one params
func (o *PostRemoteAPIJProposedDomainOneParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProposedDomainOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j proposed domain one params
func (o *PostRemoteAPIJProposedDomainOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j proposed domain one params
func (o *PostRemoteAPIJProposedDomainOneParams) WithContext(ctx context.Context) *PostRemoteAPIJProposedDomainOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j proposed domain one params
func (o *PostRemoteAPIJProposedDomainOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j proposed domain one params
func (o *PostRemoteAPIJProposedDomainOneParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJProposedDomainOneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j proposed domain one params
func (o *PostRemoteAPIJProposedDomainOneParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProposedDomainOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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