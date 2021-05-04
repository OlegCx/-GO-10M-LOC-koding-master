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

// PostRemoteAPIJAccountFetchOAuthInfoIDReader is a Reader for the PostRemoteAPIJAccountFetchOAuthInfoID structure.
type PostRemoteAPIJAccountFetchOAuthInfoIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountFetchOAuthInfoIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountFetchOAuthInfoIDOK creates a PostRemoteAPIJAccountFetchOAuthInfoIDOK with default headers values
func NewPostRemoteAPIJAccountFetchOAuthInfoIDOK() *PostRemoteAPIJAccountFetchOAuthInfoIDOK {
	return &PostRemoteAPIJAccountFetchOAuthInfoIDOK{}
}

/*PostRemoteAPIJAccountFetchOAuthInfoIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountFetchOAuthInfoIDOK struct {
	Payload PostRemoteAPIJAccountFetchOAuthInfoIDOKBody
}

func (o *PostRemoteAPIJAccountFetchOAuthInfoIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.fetchOAuthInfo/{id}][%d] postRemoteApiJAccountFetchOAuthInfoIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountFetchOAuthInfoIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJAccountFetchOAuthInfoIDOKBody post remote API j account fetch o auth info ID o k body
swagger:model PostRemoteAPIJAccountFetchOAuthInfoIDOKBody
*/
type PostRemoteAPIJAccountFetchOAuthInfoIDOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO0

	var postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJAccountFetchOAuthInfoIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO0)

	postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchOAuthInfoIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j account fetch o auth info ID o k body
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDOKBody) Validate(formats strfmt.Registry) error {
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