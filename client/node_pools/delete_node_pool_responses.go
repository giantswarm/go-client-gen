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

// DeleteNodePoolReader is a Reader for the DeleteNodePool structure.
type DeleteNodePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteNodePoolAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteNodePoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteNodePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteNodePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNodePoolAccepted creates a DeleteNodePoolAccepted with default headers values
func NewDeleteNodePoolAccepted() *DeleteNodePoolAccepted {
	return &DeleteNodePoolAccepted{}
}

/*DeleteNodePoolAccepted handles this case with default header values.

Deleting node pool
*/
type DeleteNodePoolAccepted struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteNodePoolAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] deleteNodePoolAccepted  %+v", 202, o.Payload)
}

func (o *DeleteNodePoolAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodePoolUnauthorized creates a DeleteNodePoolUnauthorized with default headers values
func NewDeleteNodePoolUnauthorized() *DeleteNodePoolUnauthorized {
	return &DeleteNodePoolUnauthorized{}
}

/*DeleteNodePoolUnauthorized handles this case with default header values.

Permission denied
*/
type DeleteNodePoolUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteNodePoolUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] deleteNodePoolUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteNodePoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodePoolNotFound creates a DeleteNodePoolNotFound with default headers values
func NewDeleteNodePoolNotFound() *DeleteNodePoolNotFound {
	return &DeleteNodePoolNotFound{}
}

/*DeleteNodePoolNotFound handles this case with default header values.

Not found
*/
type DeleteNodePoolNotFound struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteNodePoolNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] deleteNodePoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNodePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodePoolDefault creates a DeleteNodePoolDefault with default headers values
func NewDeleteNodePoolDefault(code int) *DeleteNodePoolDefault {
	return &DeleteNodePoolDefault{
		_statusCode: code,
	}
}

/*DeleteNodePoolDefault handles this case with default header values.

error
*/
type DeleteNodePoolDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the delete node pool default response
func (o *DeleteNodePoolDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNodePoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] deleteNodePool default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNodePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
