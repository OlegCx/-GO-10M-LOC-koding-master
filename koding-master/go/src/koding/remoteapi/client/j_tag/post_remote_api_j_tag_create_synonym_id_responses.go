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

// PostRemoteAPIJTagCreateSynonymIDReader is a Reader for the PostRemoteAPIJTagCreateSynonymID structure.
type PostRemoteAPIJTagCreateSynonymIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagCreateSynonymIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagCreateSynonymIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagCreateSynonymIDOK creates a PostRemoteAPIJTagCreateSynonymIDOK with default headers values
func NewPostRemoteAPIJTagCreateSynonymIDOK() *PostRemoteAPIJTagCreateSynonymIDOK {
	return &PostRemoteAPIJTagCreateSynonymIDOK{}
}

/*PostRemoteAPIJTagCreateSynonymIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJTagCreateSynonymIDOK struct {
	Payload PostRemoteAPIJTagCreateSynonymIDOKBody
}

func (o *PostRemoteAPIJTagCreateSynonymIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.createSynonym/{id}][%d] postRemoteApiJTagCreateSynonymIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagCreateSynonymIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJTagCreateSynonymIDOKBody post remote API j tag create synonym ID o k body
swagger:model PostRemoteAPIJTagCreateSynonymIDOKBody
*/
type PostRemoteAPIJTagCreateSynonymIDOKBody struct {
	models.JTag

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJTagCreateSynonymIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJTagCreateSynonymIDOKBodyAO0 models.JTag
	if err := swag.ReadJSON(raw, &postRemoteAPIJTagCreateSynonymIDOKBodyAO0); err != nil {
		return err
	}
	o.JTag = postRemoteAPIJTagCreateSynonymIDOKBodyAO0

	var postRemoteAPIJTagCreateSynonymIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJTagCreateSynonymIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJTagCreateSynonymIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJTagCreateSynonymIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJTagCreateSynonymIDOKBodyAO0, err := swag.WriteJSON(o.JTag)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTagCreateSynonymIDOKBodyAO0)

	postRemoteAPIJTagCreateSynonymIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTagCreateSynonymIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j tag create synonym ID o k body
func (o *PostRemoteAPIJTagCreateSynonymIDOKBody) Validate(formats strfmt.Registry) error {
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