package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JGroupSomeWithRelationshipReader is a Reader for the JGroupSomeWithRelationship structure.
type JGroupSomeWithRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JGroupSomeWithRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJGroupSomeWithRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJGroupSomeWithRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJGroupSomeWithRelationshipOK creates a JGroupSomeWithRelationshipOK with default headers values
func NewJGroupSomeWithRelationshipOK() *JGroupSomeWithRelationshipOK {
	return &JGroupSomeWithRelationshipOK{}
}

/*JGroupSomeWithRelationshipOK handles this case with default header values.

Request processed successfully
*/
type JGroupSomeWithRelationshipOK struct {
	Payload *models.DefaultResponse
}

func (o *JGroupSomeWithRelationshipOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.someWithRelationship][%d] jGroupSomeWithRelationshipOK  %+v", 200, o.Payload)
}

func (o *JGroupSomeWithRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJGroupSomeWithRelationshipUnauthorized creates a JGroupSomeWithRelationshipUnauthorized with default headers values
func NewJGroupSomeWithRelationshipUnauthorized() *JGroupSomeWithRelationshipUnauthorized {
	return &JGroupSomeWithRelationshipUnauthorized{}
}

/*JGroupSomeWithRelationshipUnauthorized handles this case with default header values.

Unauthorized request
*/
type JGroupSomeWithRelationshipUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JGroupSomeWithRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.someWithRelationship][%d] jGroupSomeWithRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *JGroupSomeWithRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}