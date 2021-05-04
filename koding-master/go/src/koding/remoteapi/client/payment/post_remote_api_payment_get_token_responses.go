package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIPaymentGetTokenReader is a Reader for the PostRemoteAPIPaymentGetToken structure.
type PostRemoteAPIPaymentGetTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIPaymentGetTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIPaymentGetTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIPaymentGetTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIPaymentGetTokenOK creates a PostRemoteAPIPaymentGetTokenOK with default headers values
func NewPostRemoteAPIPaymentGetTokenOK() *PostRemoteAPIPaymentGetTokenOK {
	return &PostRemoteAPIPaymentGetTokenOK{}
}

/*PostRemoteAPIPaymentGetTokenOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIPaymentGetTokenOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIPaymentGetTokenOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Payment.getToken][%d] postRemoteApiPaymentGetTokenOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIPaymentGetTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIPaymentGetTokenUnauthorized creates a PostRemoteAPIPaymentGetTokenUnauthorized with default headers values
func NewPostRemoteAPIPaymentGetTokenUnauthorized() *PostRemoteAPIPaymentGetTokenUnauthorized {
	return &PostRemoteAPIPaymentGetTokenUnauthorized{}
}

/*PostRemoteAPIPaymentGetTokenUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIPaymentGetTokenUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIPaymentGetTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Payment.getToken][%d] postRemoteApiPaymentGetTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIPaymentGetTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}