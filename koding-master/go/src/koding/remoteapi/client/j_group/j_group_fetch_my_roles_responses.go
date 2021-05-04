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

// JGroupFetchMyRolesReader is a Reader for the JGroupFetchMyRoles structure.
type JGroupFetchMyRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JGroupFetchMyRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJGroupFetchMyRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJGroupFetchMyRolesOK creates a JGroupFetchMyRolesOK with default headers values
func NewJGroupFetchMyRolesOK() *JGroupFetchMyRolesOK {
	return &JGroupFetchMyRolesOK{}
}

/*JGroupFetchMyRolesOK handles this case with default header values.

OK
*/
type JGroupFetchMyRolesOK struct {
	Payload JGroupFetchMyRolesOKBody
}

func (o *JGroupFetchMyRolesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.fetchMyRoles/{id}][%d] jGroupFetchMyRolesOK  %+v", 200, o.Payload)
}

func (o *JGroupFetchMyRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JGroupFetchMyRolesOKBody j group fetch my roles o k body
swagger:model JGroupFetchMyRolesOKBody
*/
type JGroupFetchMyRolesOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JGroupFetchMyRolesOKBody) UnmarshalJSON(raw []byte) error {

	var jGroupFetchMyRolesOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &jGroupFetchMyRolesOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = jGroupFetchMyRolesOKBodyAO0

	var jGroupFetchMyRolesOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jGroupFetchMyRolesOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jGroupFetchMyRolesOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JGroupFetchMyRolesOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jGroupFetchMyRolesOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupFetchMyRolesOKBodyAO0)

	jGroupFetchMyRolesOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupFetchMyRolesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j group fetch my roles o k body
func (o *JGroupFetchMyRolesOKBody) Validate(formats strfmt.Registry) error {
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