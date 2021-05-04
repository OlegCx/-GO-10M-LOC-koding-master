package j_group

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

// NewPostRemoteAPIJGroupToggleFeatureIDParams creates a new PostRemoteAPIJGroupToggleFeatureIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupToggleFeatureIDParams() *PostRemoteAPIJGroupToggleFeatureIDParams {
	var ()
	return &PostRemoteAPIJGroupToggleFeatureIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupToggleFeatureIDParamsWithTimeout creates a new PostRemoteAPIJGroupToggleFeatureIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupToggleFeatureIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupToggleFeatureIDParams {
	var ()
	return &PostRemoteAPIJGroupToggleFeatureIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupToggleFeatureIDParamsWithContext creates a new PostRemoteAPIJGroupToggleFeatureIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupToggleFeatureIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupToggleFeatureIDParams {
	var ()
	return &PostRemoteAPIJGroupToggleFeatureIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupToggleFeatureIDParams contains all the parameters to send to the API endpoint
for the post remote API j group toggle feature ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupToggleFeatureIDParams struct {

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

// WithTimeout adds the timeout to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupToggleFeatureIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupToggleFeatureIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJGroupToggleFeatureIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) WithID(id string) *PostRemoteAPIJGroupToggleFeatureIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group toggle feature ID params
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupToggleFeatureIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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