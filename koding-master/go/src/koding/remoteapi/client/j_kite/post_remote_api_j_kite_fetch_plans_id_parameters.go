package j_kite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJKiteFetchPlansIDParams creates a new PostRemoteAPIJKiteFetchPlansIDParams object
// with the default values initialized.
func NewPostRemoteAPIJKiteFetchPlansIDParams() *PostRemoteAPIJKiteFetchPlansIDParams {
	var ()
	return &PostRemoteAPIJKiteFetchPlansIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJKiteFetchPlansIDParamsWithTimeout creates a new PostRemoteAPIJKiteFetchPlansIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJKiteFetchPlansIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJKiteFetchPlansIDParams {
	var ()
	return &PostRemoteAPIJKiteFetchPlansIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJKiteFetchPlansIDParamsWithContext creates a new PostRemoteAPIJKiteFetchPlansIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJKiteFetchPlansIDParamsWithContext(ctx context.Context) *PostRemoteAPIJKiteFetchPlansIDParams {
	var ()
	return &PostRemoteAPIJKiteFetchPlansIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJKiteFetchPlansIDParams contains all the parameters to send to the API endpoint
for the post remote API j kite fetch plans ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJKiteFetchPlansIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j kite fetch plans ID params
func (o *PostRemoteAPIJKiteFetchPlansIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJKiteFetchPlansIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j kite fetch plans ID params
func (o *PostRemoteAPIJKiteFetchPlansIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j kite fetch plans ID params
func (o *PostRemoteAPIJKiteFetchPlansIDParams) WithContext(ctx context.Context) *PostRemoteAPIJKiteFetchPlansIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j kite fetch plans ID params
func (o *PostRemoteAPIJKiteFetchPlansIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j kite fetch plans ID params
func (o *PostRemoteAPIJKiteFetchPlansIDParams) WithID(id string) *PostRemoteAPIJKiteFetchPlansIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j kite fetch plans ID params
func (o *PostRemoteAPIJKiteFetchPlansIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJKiteFetchPlansIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
