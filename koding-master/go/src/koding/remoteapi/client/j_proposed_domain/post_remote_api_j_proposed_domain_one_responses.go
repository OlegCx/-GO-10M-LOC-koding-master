package j_proposed_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJProposedDomainOneReader is a Reader for the PostRemoteAPIJProposedDomainOne structure.
type PostRemoteAPIJProposedDomainOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJProposedDomainOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJProposedDomainOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJProposedDomainOneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJProposedDomainOneOK creates a PostRemoteAPIJProposedDomainOneOK with default headers values
func NewPostRemoteAPIJProposedDomainOneOK() *PostRemoteAPIJProposedDomainOneOK {
	return &PostRemoteAPIJProposedDomainOneOK{}
}

/*PostRemoteAPIJProposedDomainOneOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJProposedDomainOneOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJProposedDomainOneOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.one][%d] postRemoteApiJProposedDomainOneOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJProposedDomainOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJProposedDomainOneUnauthorized creates a PostRemoteAPIJProposedDomainOneUnauthorized with default headers values
func NewPostRemoteAPIJProposedDomainOneUnauthorized() *PostRemoteAPIJProposedDomainOneUnauthorized {
	return &PostRemoteAPIJProposedDomainOneUnauthorized{}
}

/*PostRemoteAPIJProposedDomainOneUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJProposedDomainOneUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJProposedDomainOneUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.one][%d] postRemoteApiJProposedDomainOneUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJProposedDomainOneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}