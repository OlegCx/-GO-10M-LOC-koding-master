package git_provider

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

// NewGitProviderFetchConfigParams creates a new GitProviderFetchConfigParams object
// with the default values initialized.
func NewGitProviderFetchConfigParams() *GitProviderFetchConfigParams {
	var ()
	return &GitProviderFetchConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGitProviderFetchConfigParamsWithTimeout creates a new GitProviderFetchConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGitProviderFetchConfigParamsWithTimeout(timeout time.Duration) *GitProviderFetchConfigParams {
	var ()
	return &GitProviderFetchConfigParams{

		timeout: timeout,
	}
}

// NewGitProviderFetchConfigParamsWithContext creates a new GitProviderFetchConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGitProviderFetchConfigParamsWithContext(ctx context.Context) *GitProviderFetchConfigParams {
	var ()
	return &GitProviderFetchConfigParams{

		Context: ctx,
	}
}

/*GitProviderFetchConfigParams contains all the parameters to send to the API endpoint
for the git provider fetch config operation typically these are written to a http.Request
*/
type GitProviderFetchConfigParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the git provider fetch config params
func (o *GitProviderFetchConfigParams) WithTimeout(timeout time.Duration) *GitProviderFetchConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the git provider fetch config params
func (o *GitProviderFetchConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the git provider fetch config params
func (o *GitProviderFetchConfigParams) WithContext(ctx context.Context) *GitProviderFetchConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the git provider fetch config params
func (o *GitProviderFetchConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the git provider fetch config params
func (o *GitProviderFetchConfigParams) WithBody(body models.DefaultSelector) *GitProviderFetchConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the git provider fetch config params
func (o *GitProviderFetchConfigParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GitProviderFetchConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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