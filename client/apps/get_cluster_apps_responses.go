// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	models "github.com/giantswarm/gsclientgen/models"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// GetClusterAppsReader is a Reader for the GetClusterApps structure.
type GetClusterAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetClusterAppsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetClusterAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterAppsOK creates a GetClusterAppsOK with default headers values
func NewGetClusterAppsOK() *GetClusterAppsOK {
	return &GetClusterAppsOK{}
}

/*GetClusterAppsOK handles this case with default header values.

Cluster apps
*/
type GetClusterAppsOK struct {
	Payload models.V4GetClusterAppsResponse
}

func (o *GetClusterAppsOK) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/apps/][%d] getClusterAppsOK  %+v", 200, o.Payload)
}

func (o *GetClusterAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterAppsUnauthorized creates a GetClusterAppsUnauthorized with default headers values
func NewGetClusterAppsUnauthorized() *GetClusterAppsUnauthorized {
	return &GetClusterAppsUnauthorized{}
}

/*GetClusterAppsUnauthorized handles this case with default header values.

Permission denied
*/
type GetClusterAppsUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *GetClusterAppsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/apps/][%d] getClusterAppsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetClusterAppsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterAppsDefault creates a GetClusterAppsDefault with default headers values
func NewGetClusterAppsDefault(code int) *GetClusterAppsDefault {
	return &GetClusterAppsDefault{
		_statusCode: code,
	}
}

/*GetClusterAppsDefault handles this case with default header values.

error
*/
type GetClusterAppsDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get cluster apps default response
func (o *GetClusterAppsDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterAppsDefault) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/apps/][%d] getClusterApps default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
