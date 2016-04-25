package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/swagger-api/models"
)

// PostSitesSiteIDSnippetsReader is a Reader for the PostSitesSiteIDSnippets structure.
type PostSitesSiteIDSnippetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostSitesSiteIDSnippetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSitesSiteIDSnippetsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostSitesSiteIDSnippetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostSitesSiteIDSnippetsCreated creates a PostSitesSiteIDSnippetsCreated with default headers values
func NewPostSitesSiteIDSnippetsCreated() *PostSitesSiteIDSnippetsCreated {
	return &PostSitesSiteIDSnippetsCreated{}
}

/*PostSitesSiteIDSnippetsCreated handles this case with default header values.

OK
*/
type PostSitesSiteIDSnippetsCreated struct {
	Payload models.Snippet
}

func (o *PostSitesSiteIDSnippetsCreated) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/snippets][%d] postSitesSiteIdSnippetsCreated  %+v", 201, o.Payload)
}

func (o *PostSitesSiteIDSnippetsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteIDSnippetsDefault creates a PostSitesSiteIDSnippetsDefault with default headers values
func NewPostSitesSiteIDSnippetsDefault(code int) *PostSitesSiteIDSnippetsDefault {
	return &PostSitesSiteIDSnippetsDefault{
		_statusCode: code,
	}
}

/*PostSitesSiteIDSnippetsDefault handles this case with default header values.

error
*/
type PostSitesSiteIDSnippetsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post sites site ID snippets default response
func (o *PostSitesSiteIDSnippetsDefault) Code() int {
	return o._statusCode
}

func (o *PostSitesSiteIDSnippetsDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/snippets][%d] PostSitesSiteIDSnippets default  %+v", o._statusCode, o.Payload)
}

func (o *PostSitesSiteIDSnippetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
