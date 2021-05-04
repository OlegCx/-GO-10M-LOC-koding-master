package j_reward_campaign

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

// JRewardCampaignUpdateReader is a Reader for the JRewardCampaignUpdate structure.
type JRewardCampaignUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JRewardCampaignUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJRewardCampaignUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJRewardCampaignUpdateOK creates a JRewardCampaignUpdateOK with default headers values
func NewJRewardCampaignUpdateOK() *JRewardCampaignUpdateOK {
	return &JRewardCampaignUpdateOK{}
}

/*JRewardCampaignUpdateOK handles this case with default header values.

OK
*/
type JRewardCampaignUpdateOK struct {
	Payload JRewardCampaignUpdateOKBody
}

func (o *JRewardCampaignUpdateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.update/{id}][%d] jRewardCampaignUpdateOK  %+v", 200, o.Payload)
}

func (o *JRewardCampaignUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JRewardCampaignUpdateOKBody j reward campaign update o k body
swagger:model JRewardCampaignUpdateOKBody
*/
type JRewardCampaignUpdateOKBody struct {
	models.JRewardCampaign

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JRewardCampaignUpdateOKBody) UnmarshalJSON(raw []byte) error {

	var jRewardCampaignUpdateOKBodyAO0 models.JRewardCampaign
	if err := swag.ReadJSON(raw, &jRewardCampaignUpdateOKBodyAO0); err != nil {
		return err
	}
	o.JRewardCampaign = jRewardCampaignUpdateOKBodyAO0

	var jRewardCampaignUpdateOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jRewardCampaignUpdateOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jRewardCampaignUpdateOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JRewardCampaignUpdateOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jRewardCampaignUpdateOKBodyAO0, err := swag.WriteJSON(o.JRewardCampaign)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jRewardCampaignUpdateOKBodyAO0)

	jRewardCampaignUpdateOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jRewardCampaignUpdateOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j reward campaign update o k body
func (o *JRewardCampaignUpdateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JRewardCampaign.Validate(formats); err != nil {
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