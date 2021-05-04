package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JAccountSomeWithRelationshipReader is a Reader for the JAccountSomeWithRelationship structure.
type JAccountSomeWithRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountSomeWithRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountSomeWithRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJAccountSomeWithRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountSomeWithRelationshipOK creates a JAccountSomeWithRelationshipOK with default headers values
func NewJAccountSomeWithRelationshipOK() *JAccountSomeWithRelationshipOK {
	return &JAccountSomeWithRelationshipOK{}
}

/*JAccountSomeWithRelationshipOK handles this case with default header values.

Request processed successfully
*/
type JAccountSomeWithRelationshipOK struct {
	Payload *models.DefaultResponse
}

func (o *JAccountSomeWithRelationshipOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.someWithRelationship][%d] jAccountSomeWithRelationshipOK  %+v", 200, o.Payload)
}

func (o *JAccountSomeWithRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJAccountSomeWithRelationshipUnauthorized creates a JAccountSomeWithRelationshipUnauthorized with default headers values
func NewJAccountSomeWithRelationshipUnauthorized() *JAccountSomeWithRelationshipUnauthorized {
	return &JAccountSomeWithRelationshipUnauthorized{}
}

/*JAccountSomeWithRelationshipUnauthorized handles this case with default header values.

Unauthorized request
*/
type JAccountSomeWithRelationshipUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JAccountSomeWithRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.someWithRelationship][%d] jAccountSomeWithRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *JAccountSomeWithRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}