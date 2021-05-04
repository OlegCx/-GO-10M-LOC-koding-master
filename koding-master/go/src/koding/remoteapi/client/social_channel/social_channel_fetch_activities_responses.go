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

// SocialChannelFetchActivitiesReader is a Reader for the SocialChannelFetchActivities structure.
type SocialChannelFetchActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialChannelFetchActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialChannelFetchActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialChannelFetchActivitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialChannelFetchActivitiesOK creates a SocialChannelFetchActivitiesOK with default headers values
func NewSocialChannelFetchActivitiesOK() *SocialChannelFetchActivitiesOK {
	return &SocialChannelFetchActivitiesOK{}
}

/*SocialChannelFetchActivitiesOK handles this case with default header values.

Request processed successfully
*/
type SocialChannelFetchActivitiesOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialChannelFetchActivitiesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.fetchActivities][%d] socialChannelFetchActivitiesOK  %+v", 200, o.Payload)
}

func (o *SocialChannelFetchActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialChannelFetchActivitiesUnauthorized creates a SocialChannelFetchActivitiesUnauthorized with default headers values
func NewSocialChannelFetchActivitiesUnauthorized() *SocialChannelFetchActivitiesUnauthorized {
	return &SocialChannelFetchActivitiesUnauthorized{}
}

/*SocialChannelFetchActivitiesUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialChannelFetchActivitiesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialChannelFetchActivitiesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.fetchActivities][%d] socialChannelFetchActivitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialChannelFetchActivitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
