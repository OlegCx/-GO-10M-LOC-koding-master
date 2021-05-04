package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialChannelRemoveParticipantsReader is a Reader for the PostRemoteAPISocialChannelRemoveParticipants structure.
type PostRemoteAPISocialChannelRemoveParticipantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialChannelRemoveParticipantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialChannelRemoveParticipantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialChannelRemoveParticipantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialChannelRemoveParticipantsOK creates a PostRemoteAPISocialChannelRemoveParticipantsOK with default headers values
func NewPostRemoteAPISocialChannelRemoveParticipantsOK() *PostRemoteAPISocialChannelRemoveParticipantsOK {
	return &PostRemoteAPISocialChannelRemoveParticipantsOK{}
}

/*PostRemoteAPISocialChannelRemoveParticipantsOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPISocialChannelRemoveParticipantsOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialChannelRemoveParticipantsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.removeParticipants][%d] postRemoteApiSocialChannelRemoveParticipantsOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialChannelRemoveParticipantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialChannelRemoveParticipantsUnauthorized creates a PostRemoteAPISocialChannelRemoveParticipantsUnauthorized with default headers values
func NewPostRemoteAPISocialChannelRemoveParticipantsUnauthorized() *PostRemoteAPISocialChannelRemoveParticipantsUnauthorized {
	return &PostRemoteAPISocialChannelRemoveParticipantsUnauthorized{}
}

/*PostRemoteAPISocialChannelRemoveParticipantsUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialChannelRemoveParticipantsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialChannelRemoveParticipantsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.removeParticipants][%d] postRemoteApiSocialChannelRemoveParticipantsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialChannelRemoveParticipantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}