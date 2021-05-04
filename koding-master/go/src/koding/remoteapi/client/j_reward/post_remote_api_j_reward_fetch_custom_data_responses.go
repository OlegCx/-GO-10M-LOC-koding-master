package j_reward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJRewardFetchCustomDataReader is a Reader for the PostRemoteAPIJRewardFetchCustomData structure.
type PostRemoteAPIJRewardFetchCustomDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJRewardFetchCustomDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJRewardFetchCustomDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJRewardFetchCustomDataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJRewardFetchCustomDataOK creates a PostRemoteAPIJRewardFetchCustomDataOK with default headers values
func NewPostRemoteAPIJRewardFetchCustomDataOK() *PostRemoteAPIJRewardFetchCustomDataOK {
	return &PostRemoteAPIJRewardFetchCustomDataOK{}
}

/*PostRemoteAPIJRewardFetchCustomDataOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJRewardFetchCustomDataOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJRewardFetchCustomDataOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JReward.fetchCustomData][%d] postRemoteApiJRewardFetchCustomDataOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJRewardFetchCustomDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJRewardFetchCustomDataUnauthorized creates a PostRemoteAPIJRewardFetchCustomDataUnauthorized with default headers values
func NewPostRemoteAPIJRewardFetchCustomDataUnauthorized() *PostRemoteAPIJRewardFetchCustomDataUnauthorized {
	return &PostRemoteAPIJRewardFetchCustomDataUnauthorized{}
}

/*PostRemoteAPIJRewardFetchCustomDataUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJRewardFetchCustomDataUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJRewardFetchCustomDataUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JReward.fetchCustomData][%d] postRemoteApiJRewardFetchCustomDataUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJRewardFetchCustomDataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
