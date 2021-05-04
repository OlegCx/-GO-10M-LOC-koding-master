package j_credential

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

// JCredentialDeleteReader is a Reader for the JCredentialDelete structure.
type JCredentialDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JCredentialDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJCredentialDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJCredentialDeleteOK creates a JCredentialDeleteOK with default headers values
func NewJCredentialDeleteOK() *JCredentialDeleteOK {
	return &JCredentialDeleteOK{}
}

/*JCredentialDeleteOK handles this case with default header values.

OK
*/
type JCredentialDeleteOK struct {
	Payload JCredentialDeleteOKBody
}

func (o *JCredentialDeleteOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCredential.delete/{id}][%d] jCredentialDeleteOK  %+v", 200, o.Payload)
}

func (o *JCredentialDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JCredentialDeleteOKBody j credential delete o k body
swagger:model JCredentialDeleteOKBody
*/
type JCredentialDeleteOKBody struct {
	models.JCredential

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JCredentialDeleteOKBody) UnmarshalJSON(raw []byte) error {

	var jCredentialDeleteOKBodyAO0 models.JCredential
	if err := swag.ReadJSON(raw, &jCredentialDeleteOKBodyAO0); err != nil {
		return err
	}
	o.JCredential = jCredentialDeleteOKBodyAO0

	var jCredentialDeleteOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jCredentialDeleteOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jCredentialDeleteOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JCredentialDeleteOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jCredentialDeleteOKBodyAO0, err := swag.WriteJSON(o.JCredential)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jCredentialDeleteOKBodyAO0)

	jCredentialDeleteOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jCredentialDeleteOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j credential delete o k body
func (o *JCredentialDeleteOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JCredential.Validate(formats); err != nil {
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