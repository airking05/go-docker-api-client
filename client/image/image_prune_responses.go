// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// ImagePruneReader is a Reader for the ImagePrune structure.
type ImagePruneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagePruneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImagePruneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewImagePruneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImagePruneOK creates a ImagePruneOK with default headers values
func NewImagePruneOK() *ImagePruneOK {
	return &ImagePruneOK{}
}

/*ImagePruneOK handles this case with default header values.

No error
*/
type ImagePruneOK struct {
	Payload *models.ImagePruneOKBody
}

func (o *ImagePruneOK) Error() string {
	return fmt.Sprintf("[POST /images/prune][%d] imagePruneOK  %+v", 200, o.Payload)
}

func (o *ImagePruneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImagePruneOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePruneInternalServerError creates a ImagePruneInternalServerError with default headers values
func NewImagePruneInternalServerError() *ImagePruneInternalServerError {
	return &ImagePruneInternalServerError{}
}

/*ImagePruneInternalServerError handles this case with default header values.

Server error
*/
type ImagePruneInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImagePruneInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images/prune][%d] imagePruneInternalServerError  %+v", 500, o.Payload)
}

func (o *ImagePruneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
