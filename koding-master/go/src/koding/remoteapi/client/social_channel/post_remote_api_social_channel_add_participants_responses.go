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

// PostRemoteAPISocialChannelAddParticipantsReader is a Reader for the PostRemoteAPISocialChannelAddParticipants structure.
type PostRemoteAPISocialChannelAddParticipantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialChannelAddParticipantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialChannelAddParticipantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialChannelAddParticipantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialChannelAddParticipantsOK creates a PostRemoteAPISocialChannelAddParticipantsOK with default headers values
func NewPostRemoteAPISocialChannelAddParticipantsOK() *PostRemoteAPISocialChannelAddParticipantsOK {
	return &PostRemoteAPISocialChannelAddParticipantsOK{}
}

/*PostRemoteAPISocialChannelAddParticipantsOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPISocialChannelAddParticipantsOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialChannelAddParticipantsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.addParticipants][%d] postRemoteApiSocialChannelAddParticipantsOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialChannelAddParticipantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialChannelAddParticipantsUnauthorized creates a PostRemoteAPISocialChannelAddParticipantsUnauthorized with default headers values
func NewPostRemoteAPISocialChannelAddParticipantsUnauthorized() *PostRemoteAPISocialChannelAddParticipantsUnauthorized {
	return &PostRemoteAPISocialChannelAddParticipantsUnauthorized{}
}

/*PostRemoteAPISocialChannelAddParticipantsUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialChannelAddParticipantsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialChannelAddParticipantsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.addParticipants][%d] postRemoteApiSocialChannelAddParticipantsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialChannelAddParticipantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}