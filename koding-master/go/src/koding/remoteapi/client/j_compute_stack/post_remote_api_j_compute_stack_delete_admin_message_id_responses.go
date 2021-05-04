package j_compute_stack

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

// PostRemoteAPIJComputeStackDeleteAdminMessageIDReader is a Reader for the PostRemoteAPIJComputeStackDeleteAdminMessageID structure.
type PostRemoteAPIJComputeStackDeleteAdminMessageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJComputeStackDeleteAdminMessageIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJComputeStackDeleteAdminMessageIDOK creates a PostRemoteAPIJComputeStackDeleteAdminMessageIDOK with default headers values
func NewPostRemoteAPIJComputeStackDeleteAdminMessageIDOK() *PostRemoteAPIJComputeStackDeleteAdminMessageIDOK {
	return &PostRemoteAPIJComputeStackDeleteAdminMessageIDOK{}
}

/*PostRemoteAPIJComputeStackDeleteAdminMessageIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJComputeStackDeleteAdminMessageIDOK struct {
	Payload PostRemoteAPIJComputeStackDeleteAdminMessageIDOKBody
}

func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JComputeStack.deleteAdminMessage/{id}][%d] postRemoteApiJComputeStackDeleteAdminMessageIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJComputeStackDeleteAdminMessageIDOKBody post remote API j compute stack delete admin message ID o k body
swagger:model PostRemoteAPIJComputeStackDeleteAdminMessageIDOKBody
*/
type PostRemoteAPIJComputeStackDeleteAdminMessageIDOKBody struct {
	models.JComputeStack

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO0 models.JComputeStack
	if err := swag.ReadJSON(raw, &postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO0); err != nil {
		return err
	}
	o.JComputeStack = postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO0

	var postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJComputeStackDeleteAdminMessageIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO0, err := swag.WriteJSON(o.JComputeStack)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO0)

	postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJComputeStackDeleteAdminMessageIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j compute stack delete admin message ID o k body
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JComputeStack.Validate(formats); err != nil {
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