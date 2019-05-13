// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// ModifyUserReader is a Reader for the ModifyUser structure.
type ModifyUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewModifyUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewModifyUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyUserOK creates a ModifyUserOK with default headers values
func NewModifyUserOK() *ModifyUserOK {
	return &ModifyUserOK{}
}

/*ModifyUserOK handles this case with default header values.

Success
*/
type ModifyUserOK struct {
	Payload *models.V4UserListItem
}

func (o *ModifyUserOK) Error() string {
	return fmt.Sprintf("[PATCH /v4/users/{email}/][%d] modifyUserOK  %+v", 200, o.Payload)
}

func (o *ModifyUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4UserListItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyUserUnauthorized creates a ModifyUserUnauthorized with default headers values
func NewModifyUserUnauthorized() *ModifyUserUnauthorized {
	return &ModifyUserUnauthorized{}
}

/*ModifyUserUnauthorized handles this case with default header values.

Permission denied
*/
type ModifyUserUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *ModifyUserUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v4/users/{email}/][%d] modifyUserUnauthorized  %+v", 401, o.Payload)
}

func (o *ModifyUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyUserDefault creates a ModifyUserDefault with default headers values
func NewModifyUserDefault(code int) *ModifyUserDefault {
	return &ModifyUserDefault{
		_statusCode: code,
	}
}

/*ModifyUserDefault handles this case with default header values.

Error
*/
type ModifyUserDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the modify user default response
func (o *ModifyUserDefault) Code() int {
	return o._statusCode
}

func (o *ModifyUserDefault) Error() string {
	return fmt.Sprintf("[PATCH /v4/users/{email}/][%d] modifyUser default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
