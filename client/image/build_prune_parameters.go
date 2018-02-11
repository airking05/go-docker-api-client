// Code generated by go-swagger; DO NOT EDIT.

package image

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

// NewBuildPruneParams creates a new BuildPruneParams object
// with the default values initialized.
func NewBuildPruneParams() *BuildPruneParams {

	return &BuildPruneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildPruneParamsWithTimeout creates a new BuildPruneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildPruneParamsWithTimeout(timeout time.Duration) *BuildPruneParams {

	return &BuildPruneParams{

		timeout: timeout,
	}
}

// NewBuildPruneParamsWithContext creates a new BuildPruneParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildPruneParamsWithContext(ctx context.Context) *BuildPruneParams {

	return &BuildPruneParams{

		Context: ctx,
	}
}

// NewBuildPruneParamsWithHTTPClient creates a new BuildPruneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildPruneParamsWithHTTPClient(client *http.Client) *BuildPruneParams {

	return &BuildPruneParams{
		HTTPClient: client,
	}
}

/*BuildPruneParams contains all the parameters to send to the API endpoint
for the build prune operation typically these are written to a http.Request
*/
type BuildPruneParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the build prune params
func (o *BuildPruneParams) WithTimeout(timeout time.Duration) *BuildPruneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build prune params
func (o *BuildPruneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build prune params
func (o *BuildPruneParams) WithContext(ctx context.Context) *BuildPruneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build prune params
func (o *BuildPruneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build prune params
func (o *BuildPruneParams) WithHTTPClient(client *http.Client) *BuildPruneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build prune params
func (o *BuildPruneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *BuildPruneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}