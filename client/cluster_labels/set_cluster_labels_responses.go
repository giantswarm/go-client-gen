// Code generated by go-swagger; DO NOT EDIT.

package cluster_labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// SetClusterLabelsReader is a Reader for the SetClusterLabels structure.
type SetClusterLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetClusterLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetClusterLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSetClusterLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetClusterLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetClusterLabelsOK creates a SetClusterLabelsOK with default headers values
func NewSetClusterLabelsOK() *SetClusterLabelsOK {
	return &SetClusterLabelsOK{}
}

/*SetClusterLabelsOK handles this case with default header values.

Labels attached to this cluster after update
*/
type SetClusterLabelsOK struct {
	Payload *models.V5ClusterLabelsResponse
}

func (o *SetClusterLabelsOK) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/labels/][%d] setClusterLabelsOK  %+v", 200, o.Payload)
}

func (o *SetClusterLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V5ClusterLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetClusterLabelsUnauthorized creates a SetClusterLabelsUnauthorized with default headers values
func NewSetClusterLabelsUnauthorized() *SetClusterLabelsUnauthorized {
	return &SetClusterLabelsUnauthorized{}
}

/*SetClusterLabelsUnauthorized handles this case with default header values.

Permission denied
*/
type SetClusterLabelsUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *SetClusterLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/labels/][%d] setClusterLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *SetClusterLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetClusterLabelsNotFound creates a SetClusterLabelsNotFound with default headers values
func NewSetClusterLabelsNotFound() *SetClusterLabelsNotFound {
	return &SetClusterLabelsNotFound{}
}

/*SetClusterLabelsNotFound handles this case with default header values.

Not found
*/
type SetClusterLabelsNotFound struct {
	Payload *models.V4GenericResponse
}

func (o *SetClusterLabelsNotFound) Error() string {
	return fmt.Sprintf("[PUT /v5/clusters/{cluster_id}/labels/][%d] setClusterLabelsNotFound  %+v", 404, o.Payload)
}

func (o *SetClusterLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}