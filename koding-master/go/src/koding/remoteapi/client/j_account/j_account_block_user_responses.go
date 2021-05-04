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

// JAccountBlockUserReader is a Reader for the JAccountBlockUser structure.
type JAccountBlockUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountBlockUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountBlockUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountBlockUserOK creates a JAccountBlockUserOK with default headers values
func NewJAccountBlockUserOK() *JAccountBlockUserOK {
	return &JAccountBlockUserOK{}
}

/*JAccountBlockUserOK handles this case with default header values.

OK
*/
type JAccountBlockUserOK struct {
	Payload JAccountBlockUserOKBody
}

func (o *JAccountBlockUserOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.blockUser/{id}][%d] jAccountBlockUserOK  %+v", 200, o.Payload)
}

func (o *JAccountBlockUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JAccountBlockUserOKBody j account block user o k body
swagger:model JAccountBlockUserOKBody
*/
type JAccountBlockUserOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JAccountBlockUserOKBody) UnmarshalJSON(raw []byte) error {

	var jAccountBlockUserOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &jAccountBlockUserOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = jAccountBlockUserOKBodyAO0

	var jAccountBlockUserOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jAccountBlockUserOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jAccountBlockUserOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JAccountBlockUserOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jAccountBlockUserOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountBlockUserOKBodyAO0)

	jAccountBlockUserOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountBlockUserOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j account block user o k body
func (o *JAccountBlockUserOKBody) Validate(formats strfmt.Registry) error {
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
