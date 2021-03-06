// Code generated by go-swagger; DO NOT EDIT.

package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// AddNodePoolReader is a Reader for the AddNodePool structure.
type AddNodePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNodePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddNodePoolCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddNodePoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddNodePoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddNodePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddNodePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddNodePoolCreated creates a AddNodePoolCreated with default headers values
func NewAddNodePoolCreated() *AddNodePoolCreated {
	return &AddNodePoolCreated{}
}

/*AddNodePoolCreated handles this case with default header values.

Node pool created
*/
type AddNodePoolCreated struct {
	/*URI of the new node pool resource
	 */
	Location string

	Payload *models.V5GetNodePoolResponse
}

func (o *AddNodePoolCreated) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/{cluster_id}/nodepools/][%d] addNodePoolCreated  %+v", 201, o.Payload)
}

func (o *AddNodePoolCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.V5GetNodePoolResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNodePoolBadRequest creates a AddNodePoolBadRequest with default headers values
func NewAddNodePoolBadRequest() *AddNodePoolBadRequest {
	return &AddNodePoolBadRequest{}
}

/*AddNodePoolBadRequest handles this case with default header values.

Bad request
*/
type AddNodePoolBadRequest struct {
	Payload *models.V4GenericResponse
}

func (o *AddNodePoolBadRequest) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/{cluster_id}/nodepools/][%d] addNodePoolBadRequest  %+v", 400, o.Payload)
}

func (o *AddNodePoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNodePoolUnauthorized creates a AddNodePoolUnauthorized with default headers values
func NewAddNodePoolUnauthorized() *AddNodePoolUnauthorized {
	return &AddNodePoolUnauthorized{}
}

/*AddNodePoolUnauthorized handles this case with default header values.

Permission denied
*/
type AddNodePoolUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *AddNodePoolUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/{cluster_id}/nodepools/][%d] addNodePoolUnauthorized  %+v", 401, o.Payload)
}

func (o *AddNodePoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNodePoolNotFound creates a AddNodePoolNotFound with default headers values
func NewAddNodePoolNotFound() *AddNodePoolNotFound {
	return &AddNodePoolNotFound{}
}

/*AddNodePoolNotFound handles this case with default header values.

Cluster not found
*/
type AddNodePoolNotFound struct {
	Payload *models.V4GenericResponse
}

func (o *AddNodePoolNotFound) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/{cluster_id}/nodepools/][%d] addNodePoolNotFound  %+v", 404, o.Payload)
}

func (o *AddNodePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNodePoolDefault creates a AddNodePoolDefault with default headers values
func NewAddNodePoolDefault(code int) *AddNodePoolDefault {
	return &AddNodePoolDefault{
		_statusCode: code,
	}
}

/*AddNodePoolDefault handles this case with default header values.

error
*/
type AddNodePoolDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the add node pool default response
func (o *AddNodePoolDefault) Code() int {
	return o._statusCode
}

func (o *AddNodePoolDefault) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/{cluster_id}/nodepools/][%d] addNodePool default  %+v", o._statusCode, o.Payload)
}

func (o *AddNodePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
