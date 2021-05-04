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

// NewJAccountFetchGroupsParams creates a new JAccountFetchGroupsParams object
// with the default values initialized.
func NewJAccountFetchGroupsParams() *JAccountFetchGroupsParams {
	var ()
	return &JAccountFetchGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountFetchGroupsParamsWithTimeout creates a new JAccountFetchGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountFetchGroupsParamsWithTimeout(timeout time.Duration) *JAccountFetchGroupsParams {
	var ()
	return &JAccountFetchGroupsParams{

		timeout: timeout,
	}
}

// NewJAccountFetchGroupsParamsWithContext creates a new JAccountFetchGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountFetchGroupsParamsWithContext(ctx context.Context) *JAccountFetchGroupsParams {
	var ()
	return &JAccountFetchGroupsParams{

		Context: ctx,
	}
}

/*JAccountFetchGroupsParams contains all the parameters to send to the API endpoint
for the j account fetch groups operation typically these are written to a http.Request
*/
type JAccountFetchGroupsParams struct {

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

// WithTimeout adds the timeout to the j account fetch groups params
func (o *JAccountFetchGroupsParams) WithTimeout(timeout time.Duration) *JAccountFetchGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account fetch groups params
func (o *JAccountFetchGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account fetch groups params
func (o *JAccountFetchGroupsParams) WithContext(ctx context.Context) *JAccountFetchGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account fetch groups params
func (o *JAccountFetchGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account fetch groups params
func (o *JAccountFetchGroupsParams) WithBody(body models.DefaultSelector) *JAccountFetchGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account fetch groups params
func (o *JAccountFetchGroupsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j account fetch groups params
func (o *JAccountFetchGroupsParams) WithID(id string) *JAccountFetchGroupsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j account fetch groups params
func (o *JAccountFetchGroupsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountFetchGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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