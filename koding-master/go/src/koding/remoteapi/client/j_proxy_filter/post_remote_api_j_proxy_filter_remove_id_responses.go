package j_proxy_filter

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

// PostRemoteAPIJProxyFilterRemoveIDReader is a Reader for the PostRemoteAPIJProxyFilterRemoveID structure.
type PostRemoteAPIJProxyFilterRemoveIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJProxyFilterRemoveIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJProxyFilterRemoveIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJProxyFilterRemoveIDOK creates a PostRemoteAPIJProxyFilterRemoveIDOK with default headers values
func NewPostRemoteAPIJProxyFilterRemoveIDOK() *PostRemoteAPIJProxyFilterRemoveIDOK {
	return &PostRemoteAPIJProxyFilterRemoveIDOK{}
}

/*PostRemoteAPIJProxyFilterRemoveIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJProxyFilterRemoveIDOK struct {
	Payload PostRemoteAPIJProxyFilterRemoveIDOKBody
}

func (o *PostRemoteAPIJProxyFilterRemoveIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProxyFilter.remove/{id}][%d] postRemoteApiJProxyFilterRemoveIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJProxyFilterRemoveIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJProxyFilterRemoveIDOKBody post remote API j proxy filter remove ID o k body
swagger:model PostRemoteAPIJProxyFilterRemoveIDOKBody
*/
type PostRemoteAPIJProxyFilterRemoveIDOKBody struct {
	models.JProxyFilter

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJProxyFilterRemoveIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJProxyFilterRemoveIDOKBodyAO0 models.JProxyFilter
	if err := swag.ReadJSON(raw, &postRemoteAPIJProxyFilterRemoveIDOKBodyAO0); err != nil {
		return err
	}
	o.JProxyFilter = postRemoteAPIJProxyFilterRemoveIDOKBodyAO0

	var postRemoteAPIJProxyFilterRemoveIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJProxyFilterRemoveIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJProxyFilterRemoveIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJProxyFilterRemoveIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJProxyFilterRemoveIDOKBodyAO0, err := swag.WriteJSON(o.JProxyFilter)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJProxyFilterRemoveIDOKBodyAO0)

	postRemoteAPIJProxyFilterRemoveIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJProxyFilterRemoveIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j proxy filter remove ID o k body
func (o *PostRemoteAPIJProxyFilterRemoveIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JProxyFilter.Validate(formats); err != nil {
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