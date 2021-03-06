// Code generated by go-swagger; DO NOT EDIT.

package app_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// DeleteClusterAppConfigV5Reader is a Reader for the DeleteClusterAppConfigV5 structure.
type DeleteClusterAppConfigV5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterAppConfigV5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterAppConfigV5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteClusterAppConfigV5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteClusterAppConfigV5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteClusterAppConfigV5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterAppConfigV5OK creates a DeleteClusterAppConfigV5OK with default headers values
func NewDeleteClusterAppConfigV5OK() *DeleteClusterAppConfigV5OK {
	return &DeleteClusterAppConfigV5OK{}
}

/*DeleteClusterAppConfigV5OK handles this case with default header values.

App Config deleted
*/
type DeleteClusterAppConfigV5OK struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppConfigV5OK) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] deleteClusterAppConfigV5OK  %+v", 200, o.Payload)
}

func (o *DeleteClusterAppConfigV5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppConfigV5Unauthorized creates a DeleteClusterAppConfigV5Unauthorized with default headers values
func NewDeleteClusterAppConfigV5Unauthorized() *DeleteClusterAppConfigV5Unauthorized {
	return &DeleteClusterAppConfigV5Unauthorized{}
}

/*DeleteClusterAppConfigV5Unauthorized handles this case with default header values.

Permission denied
*/
type DeleteClusterAppConfigV5Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppConfigV5Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] deleteClusterAppConfigV5Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteClusterAppConfigV5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppConfigV5NotFound creates a DeleteClusterAppConfigV5NotFound with default headers values
func NewDeleteClusterAppConfigV5NotFound() *DeleteClusterAppConfigV5NotFound {
	return &DeleteClusterAppConfigV5NotFound{}
}

/*DeleteClusterAppConfigV5NotFound handles this case with default header values.

App Config not found
*/
type DeleteClusterAppConfigV5NotFound struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppConfigV5NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] deleteClusterAppConfigV5NotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterAppConfigV5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppConfigV5Default creates a DeleteClusterAppConfigV5Default with default headers values
func NewDeleteClusterAppConfigV5Default(code int) *DeleteClusterAppConfigV5Default {
	return &DeleteClusterAppConfigV5Default{
		_statusCode: code,
	}
}

/*DeleteClusterAppConfigV5Default handles this case with default header values.

Error
*/
type DeleteClusterAppConfigV5Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the delete cluster app config v5 default response
func (o *DeleteClusterAppConfigV5Default) Code() int {
	return o._statusCode
}

func (o *DeleteClusterAppConfigV5Default) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] deleteClusterAppConfigV5 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterAppConfigV5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
