// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewGetResolvedParameterParams creates a new GetResolvedParameterParams object
// with the default values initialized.
func NewGetResolvedParameterParams() *GetResolvedParameterParams {
	var ()
	return &GetResolvedParameterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResolvedParameterParamsWithTimeout creates a new GetResolvedParameterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResolvedParameterParamsWithTimeout(timeout time.Duration) *GetResolvedParameterParams {
	var ()
	return &GetResolvedParameterParams{

		timeout: timeout,
	}
}

// NewGetResolvedParameterParamsWithContext creates a new GetResolvedParameterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResolvedParameterParamsWithContext(ctx context.Context) *GetResolvedParameterParams {
	var ()
	return &GetResolvedParameterParams{

		Context: ctx,
	}
}

// NewGetResolvedParameterParamsWithHTTPClient creates a new GetResolvedParameterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResolvedParameterParamsWithHTTPClient(client *http.Client) *GetResolvedParameterParams {
	var ()
	return &GetResolvedParameterParams{
		HTTPClient: client,
	}
}

/*GetResolvedParameterParams contains all the parameters to send to the API endpoint
for the get resolved parameter operation typically these are written to a http.Request
*/
type GetResolvedParameterParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*Value*/
	Value string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get resolved parameter params
func (o *GetResolvedParameterParams) WithTimeout(timeout time.Duration) *GetResolvedParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resolved parameter params
func (o *GetResolvedParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resolved parameter params
func (o *GetResolvedParameterParams) WithContext(ctx context.Context) *GetResolvedParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resolved parameter params
func (o *GetResolvedParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resolved parameter params
func (o *GetResolvedParameterParams) WithHTTPClient(client *http.Client) *GetResolvedParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resolved parameter params
func (o *GetResolvedParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get resolved parameter params
func (o *GetResolvedParameterParams) WithBuildLocator(buildLocator string) *GetResolvedParameterParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get resolved parameter params
func (o *GetResolvedParameterParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithValue adds the value to the get resolved parameter params
func (o *GetResolvedParameterParams) WithValue(value string) *GetResolvedParameterParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the get resolved parameter params
func (o *GetResolvedParameterParams) SetValue(value string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *GetResolvedParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	// path param value
	if err := r.SetPathParam("value", o.Value); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}