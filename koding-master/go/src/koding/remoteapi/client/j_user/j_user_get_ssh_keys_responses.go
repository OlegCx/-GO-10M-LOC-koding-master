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

// JUserGetSSHKeysReader is a Reader for the JUserGetSSHKeys structure.
type JUserGetSSHKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JUserGetSSHKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJUserGetSSHKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJUserGetSSHKeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJUserGetSSHKeysOK creates a JUserGetSSHKeysOK with default headers values
func NewJUserGetSSHKeysOK() *JUserGetSSHKeysOK {
	return &JUserGetSSHKeysOK{}
}

/*JUserGetSSHKeysOK handles this case with default header values.

Request processed successfully
*/
type JUserGetSSHKeysOK struct {
	Payload *models.DefaultResponse
}

func (o *JUserGetSSHKeysOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.getSSHKeys][%d] jUserGetSshKeysOK  %+v", 200, o.Payload)
}

func (o *JUserGetSSHKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJUserGetSSHKeysUnauthorized creates a JUserGetSSHKeysUnauthorized with default headers values
func NewJUserGetSSHKeysUnauthorized() *JUserGetSSHKeysUnauthorized {
	return &JUserGetSSHKeysUnauthorized{}
}

/*JUserGetSSHKeysUnauthorized handles this case with default header values.

Unauthorized request
*/
type JUserGetSSHKeysUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JUserGetSSHKeysUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.getSSHKeys][%d] jUserGetSshKeysUnauthorized  %+v", 401, o.Payload)
}

func (o *JUserGetSSHKeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
