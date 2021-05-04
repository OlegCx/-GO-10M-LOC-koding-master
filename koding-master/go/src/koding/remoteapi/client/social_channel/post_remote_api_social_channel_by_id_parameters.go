package social_channel

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

// NewPostRemoteAPISocialChannelByIDParams creates a new PostRemoteAPISocialChannelByIDParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelByIDParams() *PostRemoteAPISocialChannelByIDParams {
	var ()
	return &PostRemoteAPISocialChannelByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelByIDParamsWithTimeout creates a new PostRemoteAPISocialChannelByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelByIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelByIDParams {
	var ()
	return &PostRemoteAPISocialChannelByIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelByIDParamsWithContext creates a new PostRemoteAPISocialChannelByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelByIDParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelByIDParams {
	var ()
	return &PostRemoteAPISocialChannelByIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelByIDParams contains all the parameters to send to the API endpoint
for the post remote API social channel by ID operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelByIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social channel by ID params
func (o *PostRemoteAPISocialChannelByIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel by ID params
func (o *PostRemoteAPISocialChannelByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel by ID params
func (o *PostRemoteAPISocialChannelByIDParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel by ID params
func (o *PostRemoteAPISocialChannelByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel by ID params
func (o *PostRemoteAPISocialChannelByIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPISocialChannelByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel by ID params
func (o *PostRemoteAPISocialChannelByIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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