package j_proposed_domain

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

// JProposedDomainBindMachineReader is a Reader for the JProposedDomainBindMachine structure.
type JProposedDomainBindMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JProposedDomainBindMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJProposedDomainBindMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJProposedDomainBindMachineOK creates a JProposedDomainBindMachineOK with default headers values
func NewJProposedDomainBindMachineOK() *JProposedDomainBindMachineOK {
	return &JProposedDomainBindMachineOK{}
}

/*JProposedDomainBindMachineOK handles this case with default header values.

OK
*/
type JProposedDomainBindMachineOK struct {
	Payload JProposedDomainBindMachineOKBody
}

func (o *JProposedDomainBindMachineOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.bindMachine/{id}][%d] jProposedDomainBindMachineOK  %+v", 200, o.Payload)
}

func (o *JProposedDomainBindMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JProposedDomainBindMachineOKBody j proposed domain bind machine o k body
swagger:model JProposedDomainBindMachineOKBody
*/
type JProposedDomainBindMachineOKBody struct {
	models.JProposedDomain

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JProposedDomainBindMachineOKBody) UnmarshalJSON(raw []byte) error {

	var jProposedDomainBindMachineOKBodyAO0 models.JProposedDomain
	if err := swag.ReadJSON(raw, &jProposedDomainBindMachineOKBodyAO0); err != nil {
		return err
	}
	o.JProposedDomain = jProposedDomainBindMachineOKBodyAO0

	var jProposedDomainBindMachineOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jProposedDomainBindMachineOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jProposedDomainBindMachineOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JProposedDomainBindMachineOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jProposedDomainBindMachineOKBodyAO0, err := swag.WriteJSON(o.JProposedDomain)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jProposedDomainBindMachineOKBodyAO0)

	jProposedDomainBindMachineOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jProposedDomainBindMachineOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j proposed domain bind machine o k body
func (o *JProposedDomainBindMachineOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JProposedDomain.Validate(formats); err != nil {
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