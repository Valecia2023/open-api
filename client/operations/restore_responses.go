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

// RestoreReader is a Reader for the Restore structure.
type RestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *RestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRestoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewRestoreCreated creates a RestoreCreated with default headers values
func NewRestoreCreated() *RestoreCreated {
	return &RestoreCreated{}
}

/*RestoreCreated handles this case with default header values.

Created
*/
type RestoreCreated struct {
	Payload models.Deploy
}

func (o *RestoreCreated) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys/{deploy_id}/restore][%d] restoreCreated  %+v", 201, o.Payload)
}

func (o *RestoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDefault creates a RestoreDefault with default headers values
func NewRestoreDefault(code int) *RestoreDefault {
	return &RestoreDefault{
		_statusCode: code,
	}
}

/*RestoreDefault handles this case with default header values.

error
*/
type RestoreDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the restore default response
func (o *RestoreDefault) Code() int {
	return o._statusCode
}

func (o *RestoreDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys/{deploy_id}/restore][%d] restore default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
