package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSitesSiteIDSubmissionsParams creates a new GetSitesSiteIDSubmissionsParams object
// with the default values initialized.
func NewGetSitesSiteIDSubmissionsParams() *GetSitesSiteIDSubmissionsParams {
	var ()
	return &GetSitesSiteIDSubmissionsParams{}
}

/*GetSitesSiteIDSubmissionsParams contains all the parameters to send to the API endpoint
for the get sites site ID submissions operation typically these are written to a http.Request
*/
type GetSitesSiteIDSubmissionsParams struct {

	/*SiteID*/
	SiteID string
}

// WithSiteID adds the siteId to the get sites site ID submissions params
func (o *GetSitesSiteIDSubmissionsParams) WithSiteID(SiteID string) *GetSitesSiteIDSubmissionsParams {
	o.SiteID = SiteID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesSiteIDSubmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
