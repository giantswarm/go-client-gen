// Code generated by go-swagger; DO NOT EDIT.

package app_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// CreateClusterAppConfigV4Reader is a Reader for the CreateClusterAppConfigV4 structure.
type CreateClusterAppConfigV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterAppConfigV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateClusterAppConfigV4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateClusterAppConfigV4BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateClusterAppConfigV4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateClusterAppConfigV4Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateClusterAppConfigV4Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterAppConfigV4OK creates a CreateClusterAppConfigV4OK with default headers values
func NewCreateClusterAppConfigV4OK() *CreateClusterAppConfigV4OK {
	return &CreateClusterAppConfigV4OK{}
}

/*CreateClusterAppConfigV4OK handles this case with default header values.

Success
*/
type CreateClusterAppConfigV4OK struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppConfigV4OK) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/config/][%d] createClusterAppConfigV4OK  %+v", 200, o.Payload)
}

func (o *CreateClusterAppConfigV4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppConfigV4BadRequest creates a CreateClusterAppConfigV4BadRequest with default headers values
func NewCreateClusterAppConfigV4BadRequest() *CreateClusterAppConfigV4BadRequest {
	return &CreateClusterAppConfigV4BadRequest{}
}

/*CreateClusterAppConfigV4BadRequest handles this case with default header values.

Invalid input
*/
type CreateClusterAppConfigV4BadRequest struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppConfigV4BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/config/][%d] createClusterAppConfigV4BadRequest  %+v", 400, o.Payload)
}

func (o *CreateClusterAppConfigV4BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppConfigV4Unauthorized creates a CreateClusterAppConfigV4Unauthorized with default headers values
func NewCreateClusterAppConfigV4Unauthorized() *CreateClusterAppConfigV4Unauthorized {
	return &CreateClusterAppConfigV4Unauthorized{}
}

/*CreateClusterAppConfigV4Unauthorized handles this case with default header values.

Permission denied
*/
type CreateClusterAppConfigV4Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppConfigV4Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/config/][%d] createClusterAppConfigV4Unauthorized  %+v", 401, o.Payload)
}

func (o *CreateClusterAppConfigV4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppConfigV4Conflict creates a CreateClusterAppConfigV4Conflict with default headers values
func NewCreateClusterAppConfigV4Conflict() *CreateClusterAppConfigV4Conflict {
	return &CreateClusterAppConfigV4Conflict{}
}

/*CreateClusterAppConfigV4Conflict handles this case with default header values.

App config already exists
*/
type CreateClusterAppConfigV4Conflict struct {
	Payload *models.V4GenericResponse
}

func (o *CreateClusterAppConfigV4Conflict) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/config/][%d] createClusterAppConfigV4Conflict  %+v", 409, o.Payload)
}

func (o *CreateClusterAppConfigV4Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAppConfigV4Default creates a CreateClusterAppConfigV4Default with default headers values
func NewCreateClusterAppConfigV4Default(code int) *CreateClusterAppConfigV4Default {
	return &CreateClusterAppConfigV4Default{
		_statusCode: code,
	}
}

/*CreateClusterAppConfigV4Default handles this case with default header values.

Error
*/
type CreateClusterAppConfigV4Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the create cluster app config v4 default response
func (o *CreateClusterAppConfigV4Default) Code() int {
	return o._statusCode
}

func (o *CreateClusterAppConfigV4Default) Error() string {
	return fmt.Sprintf("[PUT /v4/clusters/{cluster_id}/apps/{app_name}/config/][%d] createClusterAppConfigV4 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterAppConfigV4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
