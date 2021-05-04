package j_url_alias

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

// NewJURLAliasResolveParams creates a new JURLAliasResolveParams object
// with the default values initialized.
func NewJURLAliasResolveParams() *JURLAliasResolveParams {
	var ()
	return &JURLAliasResolveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJURLAliasResolveParamsWithTimeout creates a new JURLAliasResolveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJURLAliasResolveParamsWithTimeout(timeout time.Duration) *JURLAliasResolveParams {
	var ()
	return &JURLAliasResolveParams{

		timeout: timeout,
	}
}

// NewJURLAliasResolveParamsWithContext creates a new JURLAliasResolveParams object
// with the default values initialized, and the ability to set a context for a request
func NewJURLAliasResolveParamsWithContext(ctx context.Context) *JURLAliasResolveParams {
	var ()
	return &JURLAliasResolveParams{

		Context: ctx,
	}
}

/*JURLAliasResolveParams contains all the parameters to send to the API endpoint
for the j Url alias resolve operation typically these are written to a http.Request
*/
type JURLAliasResolveParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j Url alias resolve params
func (o *JURLAliasResolveParams) WithTimeout(timeout time.Duration) *JURLAliasResolveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j Url alias resolve params
func (o *JURLAliasResolveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j Url alias resolve params
func (o *JURLAliasResolveParams) WithContext(ctx context.Context) *JURLAliasResolveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j Url alias resolve params
func (o *JURLAliasResolveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j Url alias resolve params
func (o *JURLAliasResolveParams) WithBody(body models.DefaultSelector) *JURLAliasResolveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j Url alias resolve params
func (o *JURLAliasResolveParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JURLAliasResolveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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