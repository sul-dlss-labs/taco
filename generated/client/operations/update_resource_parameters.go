// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/sul-dlss-labs/taco/generated/models"
)

// NewUpdateResourceParams creates a new UpdateResourceParams object
// with the default values initialized.
func NewUpdateResourceParams() *UpdateResourceParams {
	var ()
	return &UpdateResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateResourceParamsWithTimeout creates a new UpdateResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateResourceParamsWithTimeout(timeout time.Duration) *UpdateResourceParams {
	var ()
	return &UpdateResourceParams{

		timeout: timeout,
	}
}

// NewUpdateResourceParamsWithContext creates a new UpdateResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateResourceParamsWithContext(ctx context.Context) *UpdateResourceParams {
	var ()
	return &UpdateResourceParams{

		Context: ctx,
	}
}

// NewUpdateResourceParamsWithHTTPClient creates a new UpdateResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateResourceParamsWithHTTPClient(client *http.Client) *UpdateResourceParams {
	var ()
	return &UpdateResourceParams{
		HTTPClient: client,
	}
}

/*UpdateResourceParams contains all the parameters to send to the API endpoint
for the update resource operation typically these are written to a http.Request
*/
type UpdateResourceParams struct {

	/*ID
	  SDR Identifier for the Resource.

	*/
	ID string
	/*Payload
	  JSON-LD Representation of the resource metadata required fields and only the fields you wish to update (identified via its TACO identifier). Needs to fit the SDR 3.0 MAP requirements.

	*/
	Payload models.Resource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update resource params
func (o *UpdateResourceParams) WithTimeout(timeout time.Duration) *UpdateResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update resource params
func (o *UpdateResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update resource params
func (o *UpdateResourceParams) WithContext(ctx context.Context) *UpdateResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update resource params
func (o *UpdateResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update resource params
func (o *UpdateResourceParams) WithHTTPClient(client *http.Client) *UpdateResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update resource params
func (o *UpdateResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update resource params
func (o *UpdateResourceParams) WithID(id string) *UpdateResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update resource params
func (o *UpdateResourceParams) SetID(id string) {
	o.ID = id
}

// WithPayload adds the payload to the update resource params
func (o *UpdateResourceParams) WithPayload(payload models.Resource) *UpdateResourceParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the update resource params
func (o *UpdateResourceParams) SetPayload(payload models.Resource) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
