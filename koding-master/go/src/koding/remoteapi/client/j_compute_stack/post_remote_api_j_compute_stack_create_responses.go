package j_compute_stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJComputeStackCreateReader is a Reader for the PostRemoteAPIJComputeStackCreate structure.
type PostRemoteAPIJComputeStackCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJComputeStackCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJComputeStackCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJComputeStackCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJComputeStackCreateOK creates a PostRemoteAPIJComputeStackCreateOK with default headers values
func NewPostRemoteAPIJComputeStackCreateOK() *PostRemoteAPIJComputeStackCreateOK {
	return &PostRemoteAPIJComputeStackCreateOK{}
}

/*PostRemoteAPIJComputeStackCreateOK handles this case with default header values.

created JComputeStack instance
*/
type PostRemoteAPIJComputeStackCreateOK struct {
	Payload *models.JComputeStack
}

func (o *PostRemoteAPIJComputeStackCreateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JComputeStack.create][%d] postRemoteApiJComputeStackCreateOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJComputeStackCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JComputeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJComputeStackCreateUnauthorized creates a PostRemoteAPIJComputeStackCreateUnauthorized with default headers values
func NewPostRemoteAPIJComputeStackCreateUnauthorized() *PostRemoteAPIJComputeStackCreateUnauthorized {
	return &PostRemoteAPIJComputeStackCreateUnauthorized{}
}

/*PostRemoteAPIJComputeStackCreateUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJComputeStackCreateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJComputeStackCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JComputeStack.create][%d] postRemoteApiJComputeStackCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJComputeStackCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}