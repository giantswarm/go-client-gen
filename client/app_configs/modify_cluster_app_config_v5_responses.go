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

// ModifyClusterAppConfigV5Reader is a Reader for the ModifyClusterAppConfigV5 structure.
type ModifyClusterAppConfigV5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyClusterAppConfigV5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyClusterAppConfigV5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyClusterAppConfigV5BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewModifyClusterAppConfigV5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewModifyClusterAppConfigV5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyClusterAppConfigV5OK creates a ModifyClusterAppConfigV5OK with default headers values
func NewModifyClusterAppConfigV5OK() *ModifyClusterAppConfigV5OK {
	return &ModifyClusterAppConfigV5OK{}
}

/*ModifyClusterAppConfigV5OK handles this case with default header values.

Success
*/
type ModifyClusterAppConfigV5OK struct {
	Payload *models.V4GenericResponse
}

func (o *ModifyClusterAppConfigV5OK) Error() string {
	return fmt.Sprintf("[PATCH /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] modifyClusterAppConfigV5OK  %+v", 200, o.Payload)
}

func (o *ModifyClusterAppConfigV5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyClusterAppConfigV5BadRequest creates a ModifyClusterAppConfigV5BadRequest with default headers values
func NewModifyClusterAppConfigV5BadRequest() *ModifyClusterAppConfigV5BadRequest {
	return &ModifyClusterAppConfigV5BadRequest{}
}

/*ModifyClusterAppConfigV5BadRequest handles this case with default header values.

Invalid input
*/
type ModifyClusterAppConfigV5BadRequest struct {
	Payload *models.V4GenericResponse
}

func (o *ModifyClusterAppConfigV5BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] modifyClusterAppConfigV5BadRequest  %+v", 400, o.Payload)
}

func (o *ModifyClusterAppConfigV5BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyClusterAppConfigV5Unauthorized creates a ModifyClusterAppConfigV5Unauthorized with default headers values
func NewModifyClusterAppConfigV5Unauthorized() *ModifyClusterAppConfigV5Unauthorized {
	return &ModifyClusterAppConfigV5Unauthorized{}
}

/*ModifyClusterAppConfigV5Unauthorized handles this case with default header values.

Permission denied
*/
type ModifyClusterAppConfigV5Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *ModifyClusterAppConfigV5Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] modifyClusterAppConfigV5Unauthorized  %+v", 401, o.Payload)
}

func (o *ModifyClusterAppConfigV5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyClusterAppConfigV5Default creates a ModifyClusterAppConfigV5Default with default headers values
func NewModifyClusterAppConfigV5Default(code int) *ModifyClusterAppConfigV5Default {
	return &ModifyClusterAppConfigV5Default{
		_statusCode: code,
	}
}

/*ModifyClusterAppConfigV5Default handles this case with default header values.

Error
*/
type ModifyClusterAppConfigV5Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the modify cluster app config v5 default response
func (o *ModifyClusterAppConfigV5Default) Code() int {
	return o._statusCode
}

func (o *ModifyClusterAppConfigV5Default) Error() string {
	return fmt.Sprintf("[PATCH /v5/clusters/{cluster_id}/apps/{app_name}/config/][%d] modifyClusterAppConfigV5 default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyClusterAppConfigV5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}