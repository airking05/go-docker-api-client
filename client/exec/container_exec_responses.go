// Code generated by go-swagger; DO NOT EDIT.

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// ContainerExecReader is a Reader for the ContainerExec structure.
type ContainerExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewContainerExecCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerExecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewContainerExecConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerExecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerExecCreated creates a ContainerExecCreated with default headers values
func NewContainerExecCreated() *ContainerExecCreated {
	return &ContainerExecCreated{}
}

/*ContainerExecCreated handles this case with default header values.

no error
*/
type ContainerExecCreated struct {
	Payload *models.IDResponse
}

func (o *ContainerExecCreated) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecCreated  %+v", 201, o.Payload)
}

func (o *ContainerExecCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExecNotFound creates a ContainerExecNotFound with default headers values
func NewContainerExecNotFound() *ContainerExecNotFound {
	return &ContainerExecNotFound{}
}

/*ContainerExecNotFound handles this case with default header values.

no such container
*/
type ContainerExecNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerExecNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecNotFound  %+v", 404, o.Payload)
}

func (o *ContainerExecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExecConflict creates a ContainerExecConflict with default headers values
func NewContainerExecConflict() *ContainerExecConflict {
	return &ContainerExecConflict{}
}

/*ContainerExecConflict handles this case with default header values.

container is paused
*/
type ContainerExecConflict struct {
	Payload *models.ErrorResponse
}

func (o *ContainerExecConflict) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecConflict  %+v", 409, o.Payload)
}

func (o *ContainerExecConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExecInternalServerError creates a ContainerExecInternalServerError with default headers values
func NewContainerExecInternalServerError() *ContainerExecInternalServerError {
	return &ContainerExecInternalServerError{}
}

/*ContainerExecInternalServerError handles this case with default header values.

Server error
*/
type ContainerExecInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerExecInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/exec][%d] containerExecInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerExecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
