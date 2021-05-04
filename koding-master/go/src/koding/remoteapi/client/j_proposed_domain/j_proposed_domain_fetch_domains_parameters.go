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
)

// NewJProposedDomainFetchDomainsParams creates a new JProposedDomainFetchDomainsParams object
// with the default values initialized.
func NewJProposedDomainFetchDomainsParams() *JProposedDomainFetchDomainsParams {

	return &JProposedDomainFetchDomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJProposedDomainFetchDomainsParamsWithTimeout creates a new JProposedDomainFetchDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJProposedDomainFetchDomainsParamsWithTimeout(timeout time.Duration) *JProposedDomainFetchDomainsParams {

	return &JProposedDomainFetchDomainsParams{

		timeout: timeout,
	}
}

// NewJProposedDomainFetchDomainsParamsWithContext creates a new JProposedDomainFetchDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJProposedDomainFetchDomainsParamsWithContext(ctx context.Context) *JProposedDomainFetchDomainsParams {

	return &JProposedDomainFetchDomainsParams{

		Context: ctx,
	}
}

/*JProposedDomainFetchDomainsParams contains all the parameters to send to the API endpoint
for the j proposed domain fetch domains operation typically these are written to a http.Request
*/
type JProposedDomainFetchDomainsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j proposed domain fetch domains params
func (o *JProposedDomainFetchDomainsParams) WithTimeout(timeout time.Duration) *JProposedDomainFetchDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j proposed domain fetch domains params
func (o *JProposedDomainFetchDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j proposed domain fetch domains params
func (o *JProposedDomainFetchDomainsParams) WithContext(ctx context.Context) *JProposedDomainFetchDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j proposed domain fetch domains params
func (o *JProposedDomainFetchDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *JProposedDomainFetchDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
