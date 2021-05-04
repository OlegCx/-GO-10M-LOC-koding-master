package j_stack_template

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

// PostRemoteAPIJStackTemplateForceStacksToReinitIDReader is a Reader for the PostRemoteAPIJStackTemplateForceStacksToReinitID structure.
type PostRemoteAPIJStackTemplateForceStacksToReinitIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJStackTemplateForceStacksToReinitIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJStackTemplateForceStacksToReinitIDOK creates a PostRemoteAPIJStackTemplateForceStacksToReinitIDOK with default headers values
func NewPostRemoteAPIJStackTemplateForceStacksToReinitIDOK() *PostRemoteAPIJStackTemplateForceStacksToReinitIDOK {
	return &PostRemoteAPIJStackTemplateForceStacksToReinitIDOK{}
}

/*PostRemoteAPIJStackTemplateForceStacksToReinitIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJStackTemplateForceStacksToReinitIDOK struct {
	Payload PostRemoteAPIJStackTemplateForceStacksToReinitIDOKBody
}

func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.forceStacksToReinit/{id}][%d] postRemoteApiJStackTemplateForceStacksToReinitIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJStackTemplateForceStacksToReinitIDOKBody post remote API j stack template force stacks to reinit ID o k body
swagger:model PostRemoteAPIJStackTemplateForceStacksToReinitIDOKBody
*/
type PostRemoteAPIJStackTemplateForceStacksToReinitIDOKBody struct {
	models.JStackTemplate

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO0 models.JStackTemplate
	if err := swag.ReadJSON(raw, &postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO0); err != nil {
		return err
	}
	o.JStackTemplate = postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO0

	var postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJStackTemplateForceStacksToReinitIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO0, err := swag.WriteJSON(o.JStackTemplate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO0)

	postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJStackTemplateForceStacksToReinitIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j stack template force stacks to reinit ID o k body
func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JStackTemplate.Validate(formats); err != nil {
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