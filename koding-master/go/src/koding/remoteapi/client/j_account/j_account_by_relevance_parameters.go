package j_account

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

// NewJAccountByRelevanceParams creates a new JAccountByRelevanceParams object
// with the default values initialized.
func NewJAccountByRelevanceParams() *JAccountByRelevanceParams {
	var ()
	return &JAccountByRelevanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountByRelevanceParamsWithTimeout creates a new JAccountByRelevanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountByRelevanceParamsWithTimeout(timeout time.Duration) *JAccountByRelevanceParams {
	var ()
	return &JAccountByRelevanceParams{

		timeout: timeout,
	}
}

// NewJAccountByRelevanceParamsWithContext creates a new JAccountByRelevanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountByRelevanceParamsWithContext(ctx context.Context) *JAccountByRelevanceParams {
	var ()
	return &JAccountByRelevanceParams{

		Context: ctx,
	}
}

/*JAccountByRelevanceParams contains all the parameters to send to the API endpoint
for the j account by relevance operation typically these are written to a http.Request
*/
type JAccountByRelevanceParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j account by relevance params
func (o *JAccountByRelevanceParams) WithTimeout(timeout time.Duration) *JAccountByRelevanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account by relevance params
func (o *JAccountByRelevanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account by relevance params
func (o *JAccountByRelevanceParams) WithContext(ctx context.Context) *JAccountByRelevanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account by relevance params
func (o *JAccountByRelevanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account by relevance params
func (o *JAccountByRelevanceParams) WithBody(body models.DefaultSelector) *JAccountByRelevanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account by relevance params
func (o *JAccountByRelevanceParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountByRelevanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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