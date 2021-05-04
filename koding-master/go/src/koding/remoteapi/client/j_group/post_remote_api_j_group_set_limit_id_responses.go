package j_group

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

// PostRemoteAPIJGroupSetLimitIDReader is a Reader for the PostRemoteAPIJGroupSetLimitID structure.
type PostRemoteAPIJGroupSetLimitIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupSetLimitIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupSetLimitIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupSetLimitIDOK creates a PostRemoteAPIJGroupSetLimitIDOK with default headers values
func NewPostRemoteAPIJGroupSetLimitIDOK() *PostRemoteAPIJGroupSetLimitIDOK {
	return &PostRemoteAPIJGroupSetLimitIDOK{}
}

/*PostRemoteAPIJGroupSetLimitIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupSetLimitIDOK struct {
	Payload PostRemoteAPIJGroupSetLimitIDOKBody
}

func (o *PostRemoteAPIJGroupSetLimitIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.setLimit/{id}][%d] postRemoteApiJGroupSetLimitIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupSetLimitIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupSetLimitIDOKBody post remote API j group set limit ID o k body
swagger:model PostRemoteAPIJGroupSetLimitIDOKBody
*/
type PostRemoteAPIJGroupSetLimitIDOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJGroupSetLimitIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJGroupSetLimitIDOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupSetLimitIDOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = postRemoteAPIJGroupSetLimitIDOKBodyAO0

	var postRemoteAPIJGroupSetLimitIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupSetLimitIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJGroupSetLimitIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJGroupSetLimitIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJGroupSetLimitIDOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupSetLimitIDOKBodyAO0)

	postRemoteAPIJGroupSetLimitIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupSetLimitIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j group set limit ID o k body
func (o *PostRemoteAPIJGroupSetLimitIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JGroup.Validate(formats); err != nil {
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