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

// NewPostRemoteAPIJProposedDomainActivateDomainIDParams creates a new PostRemoteAPIJProposedDomainActivateDomainIDParams object
// with the default values initialized.
func NewPostRemoteAPIJProposedDomainActivateDomainIDParams() *PostRemoteAPIJProposedDomainActivateDomainIDParams {
	var ()
	return &PostRemoteAPIJProposedDomainActivateDomainIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProposedDomainActivateDomainIDParamsWithTimeout creates a new PostRemoteAPIJProposedDomainActivateDomainIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProposedDomainActivateDomainIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProposedDomainActivateDomainIDParams {
	var ()
	return &PostRemoteAPIJProposedDomainActivateDomainIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProposedDomainActivateDomainIDParamsWithContext creates a new PostRemoteAPIJProposedDomainActivateDomainIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProposedDomainActivateDomainIDParamsWithContext(ctx context.Context) *PostRemoteAPIJProposedDomainActivateDomainIDParams {
	var ()
	return &PostRemoteAPIJProposedDomainActivateDomainIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProposedDomainActivateDomainIDParams contains all the parameters to send to the API endpoint
for the post remote API j proposed domain activate domain ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJProposedDomainActivateDomainIDParams struct {

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

// WithTimeout adds the timeout to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProposedDomainActivateDomainIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) WithContext(ctx context.Context) *PostRemoteAPIJProposedDomainActivateDomainIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJProposedDomainActivateDomainIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) WithID(id string) *PostRemoteAPIJProposedDomainActivateDomainIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j proposed domain activate domain ID params
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProposedDomainActivateDomainIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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