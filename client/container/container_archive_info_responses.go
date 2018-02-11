// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// ContainerArchiveInfoReader is a Reader for the ContainerArchiveInfo structure.
type ContainerArchiveInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerArchiveInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerArchiveInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewContainerArchiveInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContainerArchiveInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerArchiveInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerArchiveInfoOK creates a ContainerArchiveInfoOK with default headers values
func NewContainerArchiveInfoOK() *ContainerArchiveInfoOK {
	return &ContainerArchiveInfoOK{}
}

/*ContainerArchiveInfoOK handles this case with default header values.

no error
*/
type ContainerArchiveInfoOK struct {
	/*TODO
	 */
	XDockerContainerPathStat string
}

func (o *ContainerArchiveInfoOK) Error() string {
	return fmt.Sprintf("[HEAD /containers/{id}/archive][%d] containerArchiveInfoOK ", 200)
}

func (o *ContainerArchiveInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Docker-Container-Path-Stat
	o.XDockerContainerPathStat = response.GetHeader("X-Docker-Container-Path-Stat")

	return nil
}

// NewContainerArchiveInfoBadRequest creates a ContainerArchiveInfoBadRequest with default headers values
func NewContainerArchiveInfoBadRequest() *ContainerArchiveInfoBadRequest {
	return &ContainerArchiveInfoBadRequest{}
}

/*ContainerArchiveInfoBadRequest handles this case with default header values.

Bad parameter
*/
type ContainerArchiveInfoBadRequest struct {
	Payload *models.ContainerArchiveInfoBadRequestBody
}

func (o *ContainerArchiveInfoBadRequest) Error() string {
	return fmt.Sprintf("[HEAD /containers/{id}/archive][%d] containerArchiveInfoBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerArchiveInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerArchiveInfoBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerArchiveInfoNotFound creates a ContainerArchiveInfoNotFound with default headers values
func NewContainerArchiveInfoNotFound() *ContainerArchiveInfoNotFound {
	return &ContainerArchiveInfoNotFound{}
}

/*ContainerArchiveInfoNotFound handles this case with default header values.

Container or path does not exist
*/
type ContainerArchiveInfoNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerArchiveInfoNotFound) Error() string {
	return fmt.Sprintf("[HEAD /containers/{id}/archive][%d] containerArchiveInfoNotFound  %+v", 404, o.Payload)
}

func (o *ContainerArchiveInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerArchiveInfoInternalServerError creates a ContainerArchiveInfoInternalServerError with default headers values
func NewContainerArchiveInfoInternalServerError() *ContainerArchiveInfoInternalServerError {
	return &ContainerArchiveInfoInternalServerError{}
}

/*ContainerArchiveInfoInternalServerError handles this case with default header values.

Server error
*/
type ContainerArchiveInfoInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerArchiveInfoInternalServerError) Error() string {
	return fmt.Sprintf("[HEAD /containers/{id}/archive][%d] containerArchiveInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerArchiveInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}