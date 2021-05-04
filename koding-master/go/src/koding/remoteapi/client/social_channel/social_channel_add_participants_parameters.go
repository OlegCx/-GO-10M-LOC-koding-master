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

// NewSocialChannelAddParticipantsParams creates a new SocialChannelAddParticipantsParams object
// with the default values initialized.
func NewSocialChannelAddParticipantsParams() *SocialChannelAddParticipantsParams {
	var ()
	return &SocialChannelAddParticipantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialChannelAddParticipantsParamsWithTimeout creates a new SocialChannelAddParticipantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialChannelAddParticipantsParamsWithTimeout(timeout time.Duration) *SocialChannelAddParticipantsParams {
	var ()
	return &SocialChannelAddParticipantsParams{

		timeout: timeout,
	}
}

// NewSocialChannelAddParticipantsParamsWithContext creates a new SocialChannelAddParticipantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialChannelAddParticipantsParamsWithContext(ctx context.Context) *SocialChannelAddParticipantsParams {
	var ()
	return &SocialChannelAddParticipantsParams{

		Context: ctx,
	}
}

/*SocialChannelAddParticipantsParams contains all the parameters to send to the API endpoint
for the social channel add participants operation typically these are written to a http.Request
*/
type SocialChannelAddParticipantsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social channel add participants params
func (o *SocialChannelAddParticipantsParams) WithTimeout(timeout time.Duration) *SocialChannelAddParticipantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social channel add participants params
func (o *SocialChannelAddParticipantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social channel add participants params
func (o *SocialChannelAddParticipantsParams) WithContext(ctx context.Context) *SocialChannelAddParticipantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social channel add participants params
func (o *SocialChannelAddParticipantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social channel add participants params
func (o *SocialChannelAddParticipantsParams) WithBody(body models.DefaultSelector) *SocialChannelAddParticipantsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social channel add participants params
func (o *SocialChannelAddParticipantsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialChannelAddParticipantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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