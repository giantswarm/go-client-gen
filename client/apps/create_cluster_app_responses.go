// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// CreateClusterAppReader is a Reader for the CreateClusterApp structure.
type CreateClusterAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateClusterAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateClusterAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateClusterAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateClusterAppConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateClusterAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterAppOK creates a CreateClusterAppOK with default headers values
func NewCreateClusterAppOK() *CreateClusterAppOK {
	return &CreateClusterAppOK{}
}

/*CreateClusterAppOK handles this case with default header values.

Create cluster app
*/
type CreateClusterAppOK struct {
	Payload *models.V4App
}

func (o *CreateClusterAppOK) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppOK  %+v", 200, o.Payload)
}

func (o *CreateClusterAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppBadRequest creates a CreateClusterAppBadRequest with default headers values
func NewCreateClusterAppBadRequest() *CreateClusterAppBadRequest {
	return &CreateClusterAppBadRequest{}
}

/*CreateClusterAppBadRequest handles this case with default header values.

Invalid input
*/
type CreateClusterAppBadRequest struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppBadRequest  %+v", 400, o.Payload)
}

func (o *CreateClusterAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppUnauthorized creates a CreateClusterAppUnauthorized with default headers values
func NewCreateClusterAppUnauthorized() *CreateClusterAppUnauthorized {
	return &CreateClusterAppUnauthorized{}
}

/*CreateClusterAppUnauthorized handles this case with default header values.

Permission denied
*/
type CreateClusterAppUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateClusterAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppConflict creates a CreateClusterAppConflict with default headers values
func NewCreateClusterAppConflict() *CreateClusterAppConflict {
	return &CreateClusterAppConflict{}
}

/*CreateClusterAppConflict handles this case with default header values.

App already exists
*/
type CreateClusterAppConflict struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppConflict) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppConflict  %+v", 409, o.Payload)
}

func (o *CreateClusterAppConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppDefault creates a CreateClusterAppDefault with default headers values
func NewCreateClusterAppDefault(code int) *CreateClusterAppDefault {
	return &CreateClusterAppDefault{
		_statusCode: code,
	}
}

/*CreateClusterAppDefault handles this case with default header values.

error
*/
type CreateClusterAppDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the create cluster app default response
func (o *CreateClusterAppDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterAppDefault) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterApp default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
