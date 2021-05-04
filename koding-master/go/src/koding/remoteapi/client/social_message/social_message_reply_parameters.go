package social_message

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

// NewSocialMessageReplyParams creates a new SocialMessageReplyParams object
// with the default values initialized.
func NewSocialMessageReplyParams() *SocialMessageReplyParams {
	var ()
	return &SocialMessageReplyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialMessageReplyParamsWithTimeout creates a new SocialMessageReplyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialMessageReplyParamsWithTimeout(timeout time.Duration) *SocialMessageReplyParams {
	var ()
	return &SocialMessageReplyParams{

		timeout: timeout,
	}
}

// NewSocialMessageReplyParamsWithContext creates a new SocialMessageReplyParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialMessageReplyParamsWithContext(ctx context.Context) *SocialMessageReplyParams {
	var ()
	return &SocialMessageReplyParams{

		Context: ctx,
	}
}

/*SocialMessageReplyParams contains all the parameters to send to the API endpoint
for the social message reply operation typically these are written to a http.Request
*/
type SocialMessageReplyParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social message reply params
func (o *SocialMessageReplyParams) WithTimeout(timeout time.Duration) *SocialMessageReplyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social message reply params
func (o *SocialMessageReplyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social message reply params
func (o *SocialMessageReplyParams) WithContext(ctx context.Context) *SocialMessageReplyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social message reply params
func (o *SocialMessageReplyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social message reply params
func (o *SocialMessageReplyParams) WithBody(body models.DefaultSelector) *SocialMessageReplyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social message reply params
func (o *SocialMessageReplyParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialMessageReplyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
