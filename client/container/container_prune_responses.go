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

// ContainerPruneReader is a Reader for the ContainerPrune structure.
type ContainerPruneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerPruneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerPruneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewContainerPruneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerPruneOK creates a ContainerPruneOK with default headers values
func NewContainerPruneOK() *ContainerPruneOK {
	return &ContainerPruneOK{}
}

/*ContainerPruneOK handles this case with default header values.

No error
*/
type ContainerPruneOK struct {
	Payload *models.ContainerPruneOKBody
}

func (o *ContainerPruneOK) Error() string {
	return fmt.Sprintf("[POST /containers/prune][%d] containerPruneOK  %+v", 200, o.Payload)
}

func (o *ContainerPruneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerPruneOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerPruneInternalServerError creates a ContainerPruneInternalServerError with default headers values
func NewContainerPruneInternalServerError() *ContainerPruneInternalServerError {
	return &ContainerPruneInternalServerError{}
}

/*ContainerPruneInternalServerError handles this case with default header values.

Server error
*/
type ContainerPruneInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerPruneInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/prune][%d] containerPruneInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerPruneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}