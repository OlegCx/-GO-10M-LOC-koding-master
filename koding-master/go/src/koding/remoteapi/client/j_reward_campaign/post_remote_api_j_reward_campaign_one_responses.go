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

// PostRemoteAPIJRewardCampaignOneReader is a Reader for the PostRemoteAPIJRewardCampaignOne structure.
type PostRemoteAPIJRewardCampaignOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJRewardCampaignOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJRewardCampaignOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJRewardCampaignOneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJRewardCampaignOneOK creates a PostRemoteAPIJRewardCampaignOneOK with default headers values
func NewPostRemoteAPIJRewardCampaignOneOK() *PostRemoteAPIJRewardCampaignOneOK {
	return &PostRemoteAPIJRewardCampaignOneOK{}
}

/*PostRemoteAPIJRewardCampaignOneOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJRewardCampaignOneOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJRewardCampaignOneOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.one][%d] postRemoteApiJRewardCampaignOneOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJRewardCampaignOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJRewardCampaignOneUnauthorized creates a PostRemoteAPIJRewardCampaignOneUnauthorized with default headers values
func NewPostRemoteAPIJRewardCampaignOneUnauthorized() *PostRemoteAPIJRewardCampaignOneUnauthorized {
	return &PostRemoteAPIJRewardCampaignOneUnauthorized{}
}

/*PostRemoteAPIJRewardCampaignOneUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJRewardCampaignOneUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJRewardCampaignOneUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.one][%d] postRemoteApiJRewardCampaignOneUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJRewardCampaignOneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
