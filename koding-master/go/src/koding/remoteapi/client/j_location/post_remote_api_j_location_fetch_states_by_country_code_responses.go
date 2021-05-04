package j_location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJLocationFetchStatesByCountryCodeReader is a Reader for the PostRemoteAPIJLocationFetchStatesByCountryCode structure.
type PostRemoteAPIJLocationFetchStatesByCountryCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJLocationFetchStatesByCountryCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJLocationFetchStatesByCountryCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJLocationFetchStatesByCountryCodeOK creates a PostRemoteAPIJLocationFetchStatesByCountryCodeOK with default headers values
func NewPostRemoteAPIJLocationFetchStatesByCountryCodeOK() *PostRemoteAPIJLocationFetchStatesByCountryCodeOK {
	return &PostRemoteAPIJLocationFetchStatesByCountryCodeOK{}
}

/*PostRemoteAPIJLocationFetchStatesByCountryCodeOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJLocationFetchStatesByCountryCodeOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJLocationFetchStatesByCountryCodeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.fetchStatesByCountryCode][%d] postRemoteApiJLocationFetchStatesByCountryCodeOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJLocationFetchStatesByCountryCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized creates a PostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized with default headers values
func NewPostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized() *PostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized {
	return &PostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized{}
}

/*PostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.fetchStatesByCountryCode][%d] postRemoteApiJLocationFetchStatesByCountryCodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJLocationFetchStatesByCountryCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}