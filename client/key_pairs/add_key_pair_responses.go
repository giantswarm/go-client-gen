// Code generated by go-swagger; DO NOT EDIT.

package key_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// AddKeyPairReader is a Reader for the AddKeyPair structure.
type AddKeyPairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddKeyPairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddKeyPairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAddKeyPairUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewAddKeyPairServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddKeyPairDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddKeyPairOK creates a AddKeyPairOK with default headers values
func NewAddKeyPairOK() *AddKeyPairOK {
	return &AddKeyPairOK{}
}

/*AddKeyPairOK handles this case with default header values.

Success
*/
type AddKeyPairOK struct {
	Payload *models.V4AddKeyPairResponse
}

func (o *AddKeyPairOK) Error() string {
	return fmt.Sprintf("[POST /v4/clusters/{cluster_id}/key-pairs/][%d] addKeyPairOK  %+v", 200, o.Payload)
}

func (o *AddKeyPairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4AddKeyPairResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKeyPairUnauthorized creates a AddKeyPairUnauthorized with default headers values
func NewAddKeyPairUnauthorized() *AddKeyPairUnauthorized {
	return &AddKeyPairUnauthorized{}
}

/*AddKeyPairUnauthorized handles this case with default header values.

Permission denied
*/
type AddKeyPairUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *AddKeyPairUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v4/clusters/{cluster_id}/key-pairs/][%d] addKeyPairUnauthorized  %+v", 401, o.Payload)
}

func (o *AddKeyPairUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKeyPairServiceUnavailable creates a AddKeyPairServiceUnavailable with default headers values
func NewAddKeyPairServiceUnavailable() *AddKeyPairServiceUnavailable {
	return &AddKeyPairServiceUnavailable{}
}

/*AddKeyPairServiceUnavailable handles this case with default header values.

Not yet available
*/
type AddKeyPairServiceUnavailable struct {
	Payload *models.V4GenericResponse
}

func (o *AddKeyPairServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v4/clusters/{cluster_id}/key-pairs/][%d] addKeyPairServiceUnavailable  %+v", 503, o.Payload)
}

func (o *AddKeyPairServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKeyPairDefault creates a AddKeyPairDefault with default headers values
func NewAddKeyPairDefault(code int) *AddKeyPairDefault {
	return &AddKeyPairDefault{
		_statusCode: code,
	}
}

/*AddKeyPairDefault handles this case with default header values.

error
*/
type AddKeyPairDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the add key pair default response
func (o *AddKeyPairDefault) Code() int {
	return o._statusCode
}

func (o *AddKeyPairDefault) Error() string {
	return fmt.Sprintf("[POST /v4/clusters/{cluster_id}/key-pairs/][%d] addKeyPair default  %+v", o._statusCode, o.Payload)
}

func (o *AddKeyPairDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
