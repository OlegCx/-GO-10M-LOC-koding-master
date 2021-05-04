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

// SocialMessagePostReader is a Reader for the SocialMessagePost structure.
type SocialMessagePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialMessagePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialMessagePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialMessagePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialMessagePostOK creates a SocialMessagePostOK with default headers values
func NewSocialMessagePostOK() *SocialMessagePostOK {
	return &SocialMessagePostOK{}
}

/*SocialMessagePostOK handles this case with default header values.

Request processed successfully
*/
type SocialMessagePostOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialMessagePostOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.post][%d] socialMessagePostOK  %+v", 200, o.Payload)
}

func (o *SocialMessagePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialMessagePostUnauthorized creates a SocialMessagePostUnauthorized with default headers values
func NewSocialMessagePostUnauthorized() *SocialMessagePostUnauthorized {
	return &SocialMessagePostUnauthorized{}
}

/*SocialMessagePostUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialMessagePostUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialMessagePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.post][%d] socialMessagePostUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialMessagePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}