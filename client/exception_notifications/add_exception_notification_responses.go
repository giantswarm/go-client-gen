// Code generated by go-swagger; DO NOT EDIT.

package exception_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// AddExceptionNotificationReader is a Reader for the AddExceptionNotification structure.
type AddExceptionNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddExceptionNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddExceptionNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddExceptionNotificationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddExceptionNotificationOK creates a AddExceptionNotificationOK with default headers values
func NewAddExceptionNotificationOK() *AddExceptionNotificationOK {
	return &AddExceptionNotificationOK{}
}

/*AddExceptionNotificationOK handles this case with default header values.

Exception notification created
*/
type AddExceptionNotificationOK struct {
	Payload *models.V4GenericResponse
}

func (o *AddExceptionNotificationOK) Error() string {
	return fmt.Sprintf("[POST /v5/exception-notifications/][%d] addExceptionNotificationOK  %+v", 200, o.Payload)
}

func (o *AddExceptionNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddExceptionNotificationDefault creates a AddExceptionNotificationDefault with default headers values
func NewAddExceptionNotificationDefault(code int) *AddExceptionNotificationDefault {
	return &AddExceptionNotificationDefault{
		_statusCode: code,
	}
}

/*AddExceptionNotificationDefault handles this case with default header values.

error
*/
type AddExceptionNotificationDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the add exception notification default response
func (o *AddExceptionNotificationDefault) Code() int {
	return o._statusCode
}

func (o *AddExceptionNotificationDefault) Error() string {
	return fmt.Sprintf("[POST /v5/exception-notifications/][%d] addExceptionNotification default  %+v", o._statusCode, o.Payload)
}

func (o *AddExceptionNotificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
