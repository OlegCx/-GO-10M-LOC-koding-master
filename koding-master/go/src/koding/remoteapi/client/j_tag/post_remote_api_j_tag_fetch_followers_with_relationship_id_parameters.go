package j_tag

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

// NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDParams creates a new PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams object
// with the default values initialized.
func NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDParams() *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams {
	var ()
	return &PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDParamsWithTimeout creates a new PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams {
	var ()
	return &PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDParamsWithContext creates a new PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDParamsWithContext(ctx context.Context) *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams {
	var ()
	return &PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams contains all the parameters to send to the API endpoint
for the post remote API j tag fetch followers with relationship ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams struct {

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

// WithTimeout adds the timeout to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) WithContext(ctx context.Context) *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) WithID(id string) *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j tag fetch followers with relationship ID params
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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