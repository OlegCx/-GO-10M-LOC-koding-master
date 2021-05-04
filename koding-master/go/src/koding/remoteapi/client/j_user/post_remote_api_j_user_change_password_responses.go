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

// PostRemoteAPIJUserChangePasswordReader is a Reader for the PostRemoteAPIJUserChangePassword structure.
type PostRemoteAPIJUserChangePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJUserChangePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJUserChangePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJUserChangePasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJUserChangePasswordOK creates a PostRemoteAPIJUserChangePasswordOK with default headers values
func NewPostRemoteAPIJUserChangePasswordOK() *PostRemoteAPIJUserChangePasswordOK {
	return &PostRemoteAPIJUserChangePasswordOK{}
}

/*PostRemoteAPIJUserChangePasswordOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJUserChangePasswordOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJUserChangePasswordOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.changePassword][%d] postRemoteApiJUserChangePasswordOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJUserChangePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJUserChangePasswordUnauthorized creates a PostRemoteAPIJUserChangePasswordUnauthorized with default headers values
func NewPostRemoteAPIJUserChangePasswordUnauthorized() *PostRemoteAPIJUserChangePasswordUnauthorized {
	return &PostRemoteAPIJUserChangePasswordUnauthorized{}
}

/*PostRemoteAPIJUserChangePasswordUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJUserChangePasswordUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJUserChangePasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.changePassword][%d] postRemoteApiJUserChangePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJUserChangePasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}