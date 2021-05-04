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

// PostRemoteAPISocialChannelListParticipantsReader is a Reader for the PostRemoteAPISocialChannelListParticipants structure.
type PostRemoteAPISocialChannelListParticipantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialChannelListParticipantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialChannelListParticipantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialChannelListParticipantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialChannelListParticipantsOK creates a PostRemoteAPISocialChannelListParticipantsOK with default headers values
func NewPostRemoteAPISocialChannelListParticipantsOK() *PostRemoteAPISocialChannelListParticipantsOK {
	return &PostRemoteAPISocialChannelListParticipantsOK{}
}

/*PostRemoteAPISocialChannelListParticipantsOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPISocialChannelListParticipantsOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialChannelListParticipantsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.listParticipants][%d] postRemoteApiSocialChannelListParticipantsOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialChannelListParticipantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialChannelListParticipantsUnauthorized creates a PostRemoteAPISocialChannelListParticipantsUnauthorized with default headers values
func NewPostRemoteAPISocialChannelListParticipantsUnauthorized() *PostRemoteAPISocialChannelListParticipantsUnauthorized {
	return &PostRemoteAPISocialChannelListParticipantsUnauthorized{}
}

/*PostRemoteAPISocialChannelListParticipantsUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialChannelListParticipantsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialChannelListParticipantsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.listParticipants][%d] postRemoteApiSocialChannelListParticipantsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialChannelListParticipantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}