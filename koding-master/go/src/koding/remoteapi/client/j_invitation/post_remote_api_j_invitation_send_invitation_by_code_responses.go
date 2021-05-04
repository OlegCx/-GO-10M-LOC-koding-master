package j_invitation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJInvitationSendInvitationByCodeReader is a Reader for the PostRemoteAPIJInvitationSendInvitationByCode structure.
type PostRemoteAPIJInvitationSendInvitationByCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJInvitationSendInvitationByCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJInvitationSendInvitationByCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJInvitationSendInvitationByCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJInvitationSendInvitationByCodeOK creates a PostRemoteAPIJInvitationSendInvitationByCodeOK with default headers values
func NewPostRemoteAPIJInvitationSendInvitationByCodeOK() *PostRemoteAPIJInvitationSendInvitationByCodeOK {
	return &PostRemoteAPIJInvitationSendInvitationByCodeOK{}
}

/*PostRemoteAPIJInvitationSendInvitationByCodeOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJInvitationSendInvitationByCodeOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJInvitationSendInvitationByCodeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JInvitation.sendInvitationByCode][%d] postRemoteApiJInvitationSendInvitationByCodeOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJInvitationSendInvitationByCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJInvitationSendInvitationByCodeUnauthorized creates a PostRemoteAPIJInvitationSendInvitationByCodeUnauthorized with default headers values
func NewPostRemoteAPIJInvitationSendInvitationByCodeUnauthorized() *PostRemoteAPIJInvitationSendInvitationByCodeUnauthorized {
	return &PostRemoteAPIJInvitationSendInvitationByCodeUnauthorized{}
}

/*PostRemoteAPIJInvitationSendInvitationByCodeUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJInvitationSendInvitationByCodeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJInvitationSendInvitationByCodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JInvitation.sendInvitationByCode][%d] postRemoteApiJInvitationSendInvitationByCodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJInvitationSendInvitationByCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}