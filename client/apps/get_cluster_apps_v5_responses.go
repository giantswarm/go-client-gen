// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// GetClusterAppsV5Reader is a Reader for the GetClusterAppsV5 structure.
type GetClusterAppsV5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterAppsV5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterAppsV5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetClusterAppsV5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetClusterAppsV5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterAppsV5OK creates a GetClusterAppsV5OK with default headers values
func NewGetClusterAppsV5OK() *GetClusterAppsV5OK {
	return &GetClusterAppsV5OK{}
}

/*GetClusterAppsV5OK handles this case with default header values.

Cluster apps
*/
type GetClusterAppsV5OK struct {
	Payload models.V4GetClusterAppsResponse
}

func (o *GetClusterAppsV5OK) Error() string {
	return fmt.Sprintf("[GET /v5/clusters/{cluster_id}/apps/][%d] getClusterAppsV5OK  %+v", 200, o.Payload)
}

func (o *GetClusterAppsV5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterAppsV5Unauthorized creates a GetClusterAppsV5Unauthorized with default headers values
func NewGetClusterAppsV5Unauthorized() *GetClusterAppsV5Unauthorized {
	return &GetClusterAppsV5Unauthorized{}
}

/*GetClusterAppsV5Unauthorized handles this case with default header values.

Permission denied
*/
type GetClusterAppsV5Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *GetClusterAppsV5Unauthorized) Error() string {
	return fmt.Sprintf("[GET /v5/clusters/{cluster_id}/apps/][%d] getClusterAppsV5Unauthorized  %+v", 401, o.Payload)
}

func (o *GetClusterAppsV5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterAppsV5Default creates a GetClusterAppsV5Default with default headers values
func NewGetClusterAppsV5Default(code int) *GetClusterAppsV5Default {
	return &GetClusterAppsV5Default{
		_statusCode: code,
	}
}

/*GetClusterAppsV5Default handles this case with default header values.

error
*/
type GetClusterAppsV5Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get cluster apps v5 default response
func (o *GetClusterAppsV5Default) Code() int {
	return o._statusCode
}

func (o *GetClusterAppsV5Default) Error() string {
	return fmt.Sprintf("[GET /v5/clusters/{cluster_id}/apps/][%d] getClusterAppsV5 default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterAppsV5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
