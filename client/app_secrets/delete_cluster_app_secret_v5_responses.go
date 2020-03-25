// Code generated by go-swagger; DO NOT EDIT.

package app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// DeleteClusterAppSecretV5Reader is a Reader for the DeleteClusterAppSecretV5 structure.
type DeleteClusterAppSecretV5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterAppSecretV5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterAppSecretV5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteClusterAppSecretV5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteClusterAppSecretV5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteClusterAppSecretV5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterAppSecretV5OK creates a DeleteClusterAppSecretV5OK with default headers values
func NewDeleteClusterAppSecretV5OK() *DeleteClusterAppSecretV5OK {
	return &DeleteClusterAppSecretV5OK{}
}

/*DeleteClusterAppSecretV5OK handles this case with default header values.

Secret deleted
*/
type DeleteClusterAppSecretV5OK struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppSecretV5OK) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV5OK  %+v", 200, o.Payload)
}

func (o *DeleteClusterAppSecretV5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppSecretV5Unauthorized creates a DeleteClusterAppSecretV5Unauthorized with default headers values
func NewDeleteClusterAppSecretV5Unauthorized() *DeleteClusterAppSecretV5Unauthorized {
	return &DeleteClusterAppSecretV5Unauthorized{}
}

/*DeleteClusterAppSecretV5Unauthorized handles this case with default header values.

Permission denied
*/
type DeleteClusterAppSecretV5Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppSecretV5Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV5Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteClusterAppSecretV5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppSecretV5NotFound creates a DeleteClusterAppSecretV5NotFound with default headers values
func NewDeleteClusterAppSecretV5NotFound() *DeleteClusterAppSecretV5NotFound {
	return &DeleteClusterAppSecretV5NotFound{}
}

/*DeleteClusterAppSecretV5NotFound handles this case with default header values.

Secret not found
*/
type DeleteClusterAppSecretV5NotFound struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppSecretV5NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV5NotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterAppSecretV5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppSecretV5Default creates a DeleteClusterAppSecretV5Default with default headers values
func NewDeleteClusterAppSecretV5Default(code int) *DeleteClusterAppSecretV5Default {
	return &DeleteClusterAppSecretV5Default{
		_statusCode: code,
	}
}

/*DeleteClusterAppSecretV5Default handles this case with default header values.

Error
*/
type DeleteClusterAppSecretV5Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the delete cluster app secret v5 default response
func (o *DeleteClusterAppSecretV5Default) Code() int {
	return o._statusCode
}

func (o *DeleteClusterAppSecretV5Default) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV5 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterAppSecretV5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
