package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDestroyParams creates a new DestroyParams object
// with the default values initialized.
func NewDestroyParams() *DestroyParams {
	var ()
	return &DestroyParams{}
}

/*DestroyParams contains all the parameters to send to the API endpoint
for the destroy operation typically these are written to a http.Request
*/
type DestroyParams struct {

	/*HookID*/
	HookID string
}

// WithHookID adds the hookId to the destroy params
func (o *DestroyParams) WithHookID(HookID string) *DestroyParams {
	o.HookID = HookID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DestroyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param hook_id
	if err := r.SetPathParam("hook_id", o.HookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
