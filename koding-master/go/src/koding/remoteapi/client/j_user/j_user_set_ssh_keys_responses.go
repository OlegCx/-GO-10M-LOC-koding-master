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

// JUserSetSSHKeysReader is a Reader for the JUserSetSSHKeys structure.
type JUserSetSSHKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JUserSetSSHKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJUserSetSSHKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJUserSetSSHKeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJUserSetSSHKeysOK creates a JUserSetSSHKeysOK with default headers values
func NewJUserSetSSHKeysOK() *JUserSetSSHKeysOK {
	return &JUserSetSSHKeysOK{}
}

/*JUserSetSSHKeysOK handles this case with default header values.

Request processed successfully
*/
type JUserSetSSHKeysOK struct {
	Payload *models.DefaultResponse
}

func (o *JUserSetSSHKeysOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.setSSHKeys][%d] jUserSetSshKeysOK  %+v", 200, o.Payload)
}

func (o *JUserSetSSHKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJUserSetSSHKeysUnauthorized creates a JUserSetSSHKeysUnauthorized with default headers values
func NewJUserSetSSHKeysUnauthorized() *JUserSetSSHKeysUnauthorized {
	return &JUserSetSSHKeysUnauthorized{}
}

/*JUserSetSSHKeysUnauthorized handles this case with default header values.

Unauthorized request
*/
type JUserSetSSHKeysUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JUserSetSSHKeysUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.setSSHKeys][%d] jUserSetSshKeysUnauthorized  %+v", 401, o.Payload)
}

func (o *JUserSetSSHKeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}