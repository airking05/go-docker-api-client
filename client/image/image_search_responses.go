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

// ImageSearchReader is a Reader for the ImageSearch structure.
type ImageSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewImageSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageSearchOK creates a ImageSearchOK with default headers values
func NewImageSearchOK() *ImageSearchOK {
	return &ImageSearchOK{}
}

/*ImageSearchOK handles this case with default header values.

No error
*/
type ImageSearchOK struct {
	Payload models.ImageSearchOKBody
}

func (o *ImageSearchOK) Error() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchOK  %+v", 200, o.Payload)
}

func (o *ImageSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageSearchInternalServerError creates a ImageSearchInternalServerError with default headers values
func NewImageSearchInternalServerError() *ImageSearchInternalServerError {
	return &ImageSearchInternalServerError{}
}

/*ImageSearchInternalServerError handles this case with default header values.

Server error
*/
type ImageSearchInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImageSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
