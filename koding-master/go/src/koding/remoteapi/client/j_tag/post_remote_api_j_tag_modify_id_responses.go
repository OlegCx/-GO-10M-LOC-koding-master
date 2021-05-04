package j_tag

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

// PostRemoteAPIJTagModifyIDReader is a Reader for the PostRemoteAPIJTagModifyID structure.
type PostRemoteAPIJTagModifyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagModifyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagModifyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagModifyIDOK creates a PostRemoteAPIJTagModifyIDOK with default headers values
func NewPostRemoteAPIJTagModifyIDOK() *PostRemoteAPIJTagModifyIDOK {
	return &PostRemoteAPIJTagModifyIDOK{}
}

/*PostRemoteAPIJTagModifyIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJTagModifyIDOK struct {
	Payload PostRemoteAPIJTagModifyIDOKBody
}

func (o *PostRemoteAPIJTagModifyIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.modify/{id}][%d] postRemoteApiJTagModifyIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagModifyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJTagModifyIDOKBody post remote API j tag modify ID o k body
swagger:model PostRemoteAPIJTagModifyIDOKBody
*/
type PostRemoteAPIJTagModifyIDOKBody struct {
	models.JTag

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJTagModifyIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJTagModifyIDOKBodyAO0 models.JTag
	if err := swag.ReadJSON(raw, &postRemoteAPIJTagModifyIDOKBodyAO0); err != nil {
		return err
	}
	o.JTag = postRemoteAPIJTagModifyIDOKBodyAO0

	var postRemoteAPIJTagModifyIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJTagModifyIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJTagModifyIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJTagModifyIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJTagModifyIDOKBodyAO0, err := swag.WriteJSON(o.JTag)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTagModifyIDOKBodyAO0)

	postRemoteAPIJTagModifyIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTagModifyIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j tag modify ID o k body
func (o *PostRemoteAPIJTagModifyIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JTag.Validate(formats); err != nil {
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