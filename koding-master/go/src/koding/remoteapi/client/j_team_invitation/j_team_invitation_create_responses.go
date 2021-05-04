package j_team_invitation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JTeamInvitationCreateReader is a Reader for the JTeamInvitationCreate structure.
type JTeamInvitationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JTeamInvitationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJTeamInvitationCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJTeamInvitationCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJTeamInvitationCreateOK creates a JTeamInvitationCreateOK with default headers values
func NewJTeamInvitationCreateOK() *JTeamInvitationCreateOK {
	return &JTeamInvitationCreateOK{}
}

/*JTeamInvitationCreateOK handles this case with default header values.

Request processed successfully
*/
type JTeamInvitationCreateOK struct {
	Payload *models.DefaultResponse
}

func (o *JTeamInvitationCreateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTeamInvitation.create][%d] jTeamInvitationCreateOK  %+v", 200, o.Payload)
}

func (o *JTeamInvitationCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJTeamInvitationCreateUnauthorized creates a JTeamInvitationCreateUnauthorized with default headers values
func NewJTeamInvitationCreateUnauthorized() *JTeamInvitationCreateUnauthorized {
	return &JTeamInvitationCreateUnauthorized{}
}

/*JTeamInvitationCreateUnauthorized handles this case with default header values.

Unauthorized request
*/
type JTeamInvitationCreateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JTeamInvitationCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTeamInvitation.create][%d] jTeamInvitationCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *JTeamInvitationCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
