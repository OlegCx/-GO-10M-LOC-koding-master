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

// PostRemoteAPISocialChannelFetchPopularTopicsReader is a Reader for the PostRemoteAPISocialChannelFetchPopularTopics structure.
type PostRemoteAPISocialChannelFetchPopularTopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialChannelFetchPopularTopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialChannelFetchPopularTopicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialChannelFetchPopularTopicsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialChannelFetchPopularTopicsOK creates a PostRemoteAPISocialChannelFetchPopularTopicsOK with default headers values
func NewPostRemoteAPISocialChannelFetchPopularTopicsOK() *PostRemoteAPISocialChannelFetchPopularTopicsOK {
	return &PostRemoteAPISocialChannelFetchPopularTopicsOK{}
}

/*PostRemoteAPISocialChannelFetchPopularTopicsOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPISocialChannelFetchPopularTopicsOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialChannelFetchPopularTopicsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.fetchPopularTopics][%d] postRemoteApiSocialChannelFetchPopularTopicsOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialChannelFetchPopularTopicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialChannelFetchPopularTopicsUnauthorized creates a PostRemoteAPISocialChannelFetchPopularTopicsUnauthorized with default headers values
func NewPostRemoteAPISocialChannelFetchPopularTopicsUnauthorized() *PostRemoteAPISocialChannelFetchPopularTopicsUnauthorized {
	return &PostRemoteAPISocialChannelFetchPopularTopicsUnauthorized{}
}

/*PostRemoteAPISocialChannelFetchPopularTopicsUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialChannelFetchPopularTopicsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialChannelFetchPopularTopicsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.fetchPopularTopics][%d] postRemoteApiSocialChannelFetchPopularTopicsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialChannelFetchPopularTopicsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}