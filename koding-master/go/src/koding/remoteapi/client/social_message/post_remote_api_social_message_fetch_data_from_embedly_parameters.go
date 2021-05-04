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

// NewPostRemoteAPISocialMessageFetchDataFromEmbedlyParams creates a new PostRemoteAPISocialMessageFetchDataFromEmbedlyParams object
// with the default values initialized.
func NewPostRemoteAPISocialMessageFetchDataFromEmbedlyParams() *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams {
	var ()
	return &PostRemoteAPISocialMessageFetchDataFromEmbedlyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialMessageFetchDataFromEmbedlyParamsWithTimeout creates a new PostRemoteAPISocialMessageFetchDataFromEmbedlyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialMessageFetchDataFromEmbedlyParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams {
	var ()
	return &PostRemoteAPISocialMessageFetchDataFromEmbedlyParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialMessageFetchDataFromEmbedlyParamsWithContext creates a new PostRemoteAPISocialMessageFetchDataFromEmbedlyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialMessageFetchDataFromEmbedlyParamsWithContext(ctx context.Context) *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams {
	var ()
	return &PostRemoteAPISocialMessageFetchDataFromEmbedlyParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialMessageFetchDataFromEmbedlyParams contains all the parameters to send to the API endpoint
for the post remote API social message fetch data from embedly operation typically these are written to a http.Request
*/
type PostRemoteAPISocialMessageFetchDataFromEmbedlyParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social message fetch data from embedly params
func (o *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social message fetch data from embedly params
func (o *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social message fetch data from embedly params
func (o *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams) WithContext(ctx context.Context) *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social message fetch data from embedly params
func (o *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social message fetch data from embedly params
func (o *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams) WithBody(body models.DefaultSelector) *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social message fetch data from embedly params
func (o *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialMessageFetchDataFromEmbedlyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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