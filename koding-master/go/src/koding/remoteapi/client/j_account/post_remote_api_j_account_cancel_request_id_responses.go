package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAccountCancelRequestIDReader is a Reader for the PostRemoteAPIJAccountCancelRequestID structure.
type PostRemoteAPIJAccountCancelRequestIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountCancelRequestIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountCancelRequestIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountCancelRequestIDOK creates a PostRemoteAPIJAccountCancelRequestIDOK with default headers values
func NewPostRemoteAPIJAccountCancelRequestIDOK() *PostRemoteAPIJAccountCancelRequestIDOK {
	return &PostRemoteAPIJAccountCancelRequestIDOK{}
}

/*PostRemoteAPIJAccountCancelRequestIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountCancelRequestIDOK struct {
	Payload PostRemoteAPIJAccountCancelRequestIDOKBody
}

func (o *PostRemoteAPIJAccountCancelRequestIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.cancelRequest/{id}][%d] postRemoteApiJAccountCancelRequestIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountCancelRequestIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJAccountCancelRequestIDOKBody post remote API j account cancel request ID o k body
swagger:model PostRemoteAPIJAccountCancelRequestIDOKBody
*/
type PostRemoteAPIJAccountCancelRequestIDOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJAccountCancelRequestIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJAccountCancelRequestIDOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountCancelRequestIDOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = postRemoteAPIJAccountCancelRequestIDOKBodyAO0

	var postRemoteAPIJAccountCancelRequestIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountCancelRequestIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJAccountCancelRequestIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJAccountCancelRequestIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJAccountCancelRequestIDOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountCancelRequestIDOKBodyAO0)

	postRemoteAPIJAccountCancelRequestIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountCancelRequestIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j account cancel request ID o k body
func (o *PostRemoteAPIJAccountCancelRequestIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}