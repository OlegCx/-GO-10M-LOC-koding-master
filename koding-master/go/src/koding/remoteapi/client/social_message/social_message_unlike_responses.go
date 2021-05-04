package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SocialMessageUnlikeReader is a Reader for the SocialMessageUnlike structure.
type SocialMessageUnlikeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialMessageUnlikeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialMessageUnlikeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialMessageUnlikeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialMessageUnlikeOK creates a SocialMessageUnlikeOK with default headers values
func NewSocialMessageUnlikeOK() *SocialMessageUnlikeOK {
	return &SocialMessageUnlikeOK{}
}

/*SocialMessageUnlikeOK handles this case with default header values.

Request processed successfully
*/
type SocialMessageUnlikeOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialMessageUnlikeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.unlike][%d] socialMessageUnlikeOK  %+v", 200, o.Payload)
}

func (o *SocialMessageUnlikeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialMessageUnlikeUnauthorized creates a SocialMessageUnlikeUnauthorized with default headers values
func NewSocialMessageUnlikeUnauthorized() *SocialMessageUnlikeUnauthorized {
	return &SocialMessageUnlikeUnauthorized{}
}

/*SocialMessageUnlikeUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialMessageUnlikeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialMessageUnlikeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.unlike][%d] socialMessageUnlikeUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialMessageUnlikeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}