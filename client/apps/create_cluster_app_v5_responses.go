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

// CreateClusterAppV5Reader is a Reader for the CreateClusterAppV5 structure.
type CreateClusterAppV5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterAppV5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateClusterAppV5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateClusterAppV5BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateClusterAppV5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateClusterAppV5Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateClusterAppV5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterAppV5OK creates a CreateClusterAppV5OK with default headers values
func NewCreateClusterAppV5OK() *CreateClusterAppV5OK {
	return &CreateClusterAppV5OK{}
}

/*CreateClusterAppV5OK handles this case with default header values.

Create cluster app
*/
type CreateClusterAppV5OK struct {
	Payload *models.V4App
}

func (o *CreateClusterAppV5OK) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppV5OK  %+v", 200, o.Payload)
}

func (o *CreateClusterAppV5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppV5BadRequest creates a CreateClusterAppV5BadRequest with default headers values
func NewCreateClusterAppV5BadRequest() *CreateClusterAppV5BadRequest {
	return &CreateClusterAppV5BadRequest{}
}

/*CreateClusterAppV5BadRequest handles this case with default header values.

Invalid input
*/
type CreateClusterAppV5BadRequest struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppV5BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppV5BadRequest  %+v", 400, o.Payload)
}

func (o *CreateClusterAppV5BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppV5Unauthorized creates a CreateClusterAppV5Unauthorized with default headers values
func NewCreateClusterAppV5Unauthorized() *CreateClusterAppV5Unauthorized {
	return &CreateClusterAppV5Unauthorized{}
}

/*CreateClusterAppV5Unauthorized handles this case with default header values.

Permission denied
*/
type CreateClusterAppV5Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppV5Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppV5Unauthorized  %+v", 401, o.Payload)
}

func (o *CreateClusterAppV5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppV5Conflict creates a CreateClusterAppV5Conflict with default headers values
func NewCreateClusterAppV5Conflict() *CreateClusterAppV5Conflict {
	return &CreateClusterAppV5Conflict{}
}

/*CreateClusterAppV5Conflict handles this case with default header values.

App already exists
*/
type CreateClusterAppV5Conflict struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppV5Conflict) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppV5Conflict  %+v", 409, o.Payload)
}

func (o *CreateClusterAppV5Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppV5Default creates a CreateClusterAppV5Default with default headers values
func NewCreateClusterAppV5Default(code int) *CreateClusterAppV5Default {
	return &CreateClusterAppV5Default{
		_statusCode: code,
	}
}

/*CreateClusterAppV5Default handles this case with default header values.

error
*/
type CreateClusterAppV5Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the create cluster app v5 default response
func (o *CreateClusterAppV5Default) Code() int {
	return o._statusCode
}

func (o *CreateClusterAppV5Default) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/apps/{app_name}/][%d] createClusterAppV5 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterAppV5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
