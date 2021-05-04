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

// NewJProposedDomainCreateDomainParams creates a new JProposedDomainCreateDomainParams object
// with the default values initialized.
func NewJProposedDomainCreateDomainParams() *JProposedDomainCreateDomainParams {
	var ()
	return &JProposedDomainCreateDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJProposedDomainCreateDomainParamsWithTimeout creates a new JProposedDomainCreateDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJProposedDomainCreateDomainParamsWithTimeout(timeout time.Duration) *JProposedDomainCreateDomainParams {
	var ()
	return &JProposedDomainCreateDomainParams{

		timeout: timeout,
	}
}

// NewJProposedDomainCreateDomainParamsWithContext creates a new JProposedDomainCreateDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewJProposedDomainCreateDomainParamsWithContext(ctx context.Context) *JProposedDomainCreateDomainParams {
	var ()
	return &JProposedDomainCreateDomainParams{

		Context: ctx,
	}
}

/*JProposedDomainCreateDomainParams contains all the parameters to send to the API endpoint
for the j proposed domain create domain operation typically these are written to a http.Request
*/
type JProposedDomainCreateDomainParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j proposed domain create domain params
func (o *JProposedDomainCreateDomainParams) WithTimeout(timeout time.Duration) *JProposedDomainCreateDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j proposed domain create domain params
func (o *JProposedDomainCreateDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j proposed domain create domain params
func (o *JProposedDomainCreateDomainParams) WithContext(ctx context.Context) *JProposedDomainCreateDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j proposed domain create domain params
func (o *JProposedDomainCreateDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j proposed domain create domain params
func (o *JProposedDomainCreateDomainParams) WithBody(body models.DefaultSelector) *JProposedDomainCreateDomainParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j proposed domain create domain params
func (o *JProposedDomainCreateDomainParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JProposedDomainCreateDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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