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

// NewPostRemoteAPISocialChannelAddParticipantsParams creates a new PostRemoteAPISocialChannelAddParticipantsParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelAddParticipantsParams() *PostRemoteAPISocialChannelAddParticipantsParams {
	var ()
	return &PostRemoteAPISocialChannelAddParticipantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelAddParticipantsParamsWithTimeout creates a new PostRemoteAPISocialChannelAddParticipantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelAddParticipantsParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelAddParticipantsParams {
	var ()
	return &PostRemoteAPISocialChannelAddParticipantsParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelAddParticipantsParamsWithContext creates a new PostRemoteAPISocialChannelAddParticipantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelAddParticipantsParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelAddParticipantsParams {
	var ()
	return &PostRemoteAPISocialChannelAddParticipantsParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelAddParticipantsParams contains all the parameters to send to the API endpoint
for the post remote API social channel add participants operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelAddParticipantsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social channel add participants params
func (o *PostRemoteAPISocialChannelAddParticipantsParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelAddParticipantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel add participants params
func (o *PostRemoteAPISocialChannelAddParticipantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel add participants params
func (o *PostRemoteAPISocialChannelAddParticipantsParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelAddParticipantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel add participants params
func (o *PostRemoteAPISocialChannelAddParticipantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel add participants params
func (o *PostRemoteAPISocialChannelAddParticipantsParams) WithBody(body models.DefaultSelector) *PostRemoteAPISocialChannelAddParticipantsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel add participants params
func (o *PostRemoteAPISocialChannelAddParticipantsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelAddParticipantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
