package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JUserUnregisterReader is a Reader for the JUserUnregister structure.
type JUserUnregisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JUserUnregisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJUserUnregisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJUserUnregisterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJUserUnregisterOK creates a JUserUnregisterOK with default headers values
func NewJUserUnregisterOK() *JUserUnregisterOK {
	return &JUserUnregisterOK{}
}

/*JUserUnregisterOK handles this case with default header values.

Request processed successfully
*/
type JUserUnregisterOK struct {
	Payload *models.DefaultResponse
}

func (o *JUserUnregisterOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.unregister][%d] jUserUnregisterOK  %+v", 200, o.Payload)
}

func (o *JUserUnregisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJUserUnregisterUnauthorized creates a JUserUnregisterUnauthorized with default headers values
func NewJUserUnregisterUnauthorized() *JUserUnregisterUnauthorized {
	return &JUserUnregisterUnauthorized{}
}

/*JUserUnregisterUnauthorized handles this case with default header values.

Unauthorized request
*/
type JUserUnregisterUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JUserUnregisterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.unregister][%d] jUserUnregisterUnauthorized  %+v", 401, o.Payload)
}

func (o *JUserUnregisterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}