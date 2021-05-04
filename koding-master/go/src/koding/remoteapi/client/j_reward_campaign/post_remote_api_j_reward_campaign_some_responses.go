package j_reward_campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJRewardCampaignSomeReader is a Reader for the PostRemoteAPIJRewardCampaignSome structure.
type PostRemoteAPIJRewardCampaignSomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJRewardCampaignSomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJRewardCampaignSomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJRewardCampaignSomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJRewardCampaignSomeOK creates a PostRemoteAPIJRewardCampaignSomeOK with default headers values
func NewPostRemoteAPIJRewardCampaignSomeOK() *PostRemoteAPIJRewardCampaignSomeOK {
	return &PostRemoteAPIJRewardCampaignSomeOK{}
}

/*PostRemoteAPIJRewardCampaignSomeOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJRewardCampaignSomeOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJRewardCampaignSomeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.some][%d] postRemoteApiJRewardCampaignSomeOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJRewardCampaignSomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJRewardCampaignSomeUnauthorized creates a PostRemoteAPIJRewardCampaignSomeUnauthorized with default headers values
func NewPostRemoteAPIJRewardCampaignSomeUnauthorized() *PostRemoteAPIJRewardCampaignSomeUnauthorized {
	return &PostRemoteAPIJRewardCampaignSomeUnauthorized{}
}

/*PostRemoteAPIJRewardCampaignSomeUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJRewardCampaignSomeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJRewardCampaignSomeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.some][%d] postRemoteApiJRewardCampaignSomeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJRewardCampaignSomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
