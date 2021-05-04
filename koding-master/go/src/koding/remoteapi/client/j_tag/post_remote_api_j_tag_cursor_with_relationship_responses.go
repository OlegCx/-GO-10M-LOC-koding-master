package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTagCursorWithRelationshipReader is a Reader for the PostRemoteAPIJTagCursorWithRelationship structure.
type PostRemoteAPIJTagCursorWithRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagCursorWithRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagCursorWithRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJTagCursorWithRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagCursorWithRelationshipOK creates a PostRemoteAPIJTagCursorWithRelationshipOK with default headers values
func NewPostRemoteAPIJTagCursorWithRelationshipOK() *PostRemoteAPIJTagCursorWithRelationshipOK {
	return &PostRemoteAPIJTagCursorWithRelationshipOK{}
}

/*PostRemoteAPIJTagCursorWithRelationshipOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJTagCursorWithRelationshipOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJTagCursorWithRelationshipOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.cursorWithRelationship][%d] postRemoteApiJTagCursorWithRelationshipOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagCursorWithRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJTagCursorWithRelationshipUnauthorized creates a PostRemoteAPIJTagCursorWithRelationshipUnauthorized with default headers values
func NewPostRemoteAPIJTagCursorWithRelationshipUnauthorized() *PostRemoteAPIJTagCursorWithRelationshipUnauthorized {
	return &PostRemoteAPIJTagCursorWithRelationshipUnauthorized{}
}

/*PostRemoteAPIJTagCursorWithRelationshipUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJTagCursorWithRelationshipUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJTagCursorWithRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.cursorWithRelationship][%d] postRemoteApiJTagCursorWithRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJTagCursorWithRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}