package j_reward_campaign

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

// NewJRewardCampaignOneParams creates a new JRewardCampaignOneParams object
// with the default values initialized.
func NewJRewardCampaignOneParams() *JRewardCampaignOneParams {
	var ()
	return &JRewardCampaignOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJRewardCampaignOneParamsWithTimeout creates a new JRewardCampaignOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJRewardCampaignOneParamsWithTimeout(timeout time.Duration) *JRewardCampaignOneParams {
	var ()
	return &JRewardCampaignOneParams{

		timeout: timeout,
	}
}

// NewJRewardCampaignOneParamsWithContext creates a new JRewardCampaignOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewJRewardCampaignOneParamsWithContext(ctx context.Context) *JRewardCampaignOneParams {
	var ()
	return &JRewardCampaignOneParams{

		Context: ctx,
	}
}

/*JRewardCampaignOneParams contains all the parameters to send to the API endpoint
for the j reward campaign one operation typically these are written to a http.Request
*/
type JRewardCampaignOneParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j reward campaign one params
func (o *JRewardCampaignOneParams) WithTimeout(timeout time.Duration) *JRewardCampaignOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j reward campaign one params
func (o *JRewardCampaignOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j reward campaign one params
func (o *JRewardCampaignOneParams) WithContext(ctx context.Context) *JRewardCampaignOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j reward campaign one params
func (o *JRewardCampaignOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j reward campaign one params
func (o *JRewardCampaignOneParams) WithBody(body models.DefaultSelector) *JRewardCampaignOneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j reward campaign one params
func (o *JRewardCampaignOneParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JRewardCampaignOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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