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

// PostRemoteAPIJGroupSearchMembersIDReader is a Reader for the PostRemoteAPIJGroupSearchMembersID structure.
type PostRemoteAPIJGroupSearchMembersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupSearchMembersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupSearchMembersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupSearchMembersIDOK creates a PostRemoteAPIJGroupSearchMembersIDOK with default headers values
func NewPostRemoteAPIJGroupSearchMembersIDOK() *PostRemoteAPIJGroupSearchMembersIDOK {
	return &PostRemoteAPIJGroupSearchMembersIDOK{}
}

/*PostRemoteAPIJGroupSearchMembersIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupSearchMembersIDOK struct {
	Payload PostRemoteAPIJGroupSearchMembersIDOKBody
}

func (o *PostRemoteAPIJGroupSearchMembersIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.searchMembers/{id}][%d] postRemoteApiJGroupSearchMembersIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupSearchMembersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupSearchMembersIDOKBody post remote API j group search members ID o k body
swagger:model PostRemoteAPIJGroupSearchMembersIDOKBody
*/
type PostRemoteAPIJGroupSearchMembersIDOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJGroupSearchMembersIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJGroupSearchMembersIDOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupSearchMembersIDOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = postRemoteAPIJGroupSearchMembersIDOKBodyAO0

	var postRemoteAPIJGroupSearchMembersIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupSearchMembersIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJGroupSearchMembersIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJGroupSearchMembersIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJGroupSearchMembersIDOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupSearchMembersIDOKBodyAO0)

	postRemoteAPIJGroupSearchMembersIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupSearchMembersIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j group search members ID o k body
func (o *PostRemoteAPIJGroupSearchMembersIDOKBody) Validate(formats strfmt.Registry) error {
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