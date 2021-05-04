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

// PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDReader is a Reader for the PostRemoteAPIJGroupFetchBlockedAccountsWithEmailID structure.
type PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK creates a PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK with default headers values
func NewPostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK() *PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK {
	return &PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK{}
}

/*PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK struct {
	Payload PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBody
}

func (o *PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.fetchBlockedAccountsWithEmail/{id}][%d] postRemoteApiJGroupFetchBlockedAccountsWithEmailIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBody post remote API j group fetch blocked accounts with email ID o k body
swagger:model PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBody
*/
type PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO0

	var postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO0)

	postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j group fetch blocked accounts with email ID o k body
func (o *PostRemoteAPIJGroupFetchBlockedAccountsWithEmailIDOKBody) Validate(formats strfmt.Registry) error {
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