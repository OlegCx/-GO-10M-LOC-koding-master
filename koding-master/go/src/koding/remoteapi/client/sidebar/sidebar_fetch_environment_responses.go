package sidebar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SidebarFetchEnvironmentReader is a Reader for the SidebarFetchEnvironment structure.
type SidebarFetchEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SidebarFetchEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSidebarFetchEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSidebarFetchEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSidebarFetchEnvironmentOK creates a SidebarFetchEnvironmentOK with default headers values
func NewSidebarFetchEnvironmentOK() *SidebarFetchEnvironmentOK {
	return &SidebarFetchEnvironmentOK{}
}

/*SidebarFetchEnvironmentOK handles this case with default header values.

Request processed successfully
*/
type SidebarFetchEnvironmentOK struct {
	Payload *models.DefaultResponse
}

func (o *SidebarFetchEnvironmentOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Sidebar.fetchEnvironment][%d] sidebarFetchEnvironmentOK  %+v", 200, o.Payload)
}

func (o *SidebarFetchEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSidebarFetchEnvironmentUnauthorized creates a SidebarFetchEnvironmentUnauthorized with default headers values
func NewSidebarFetchEnvironmentUnauthorized() *SidebarFetchEnvironmentUnauthorized {
	return &SidebarFetchEnvironmentUnauthorized{}
}

/*SidebarFetchEnvironmentUnauthorized handles this case with default header values.

Unauthorized request
*/
type SidebarFetchEnvironmentUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SidebarFetchEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Sidebar.fetchEnvironment][%d] sidebarFetchEnvironmentUnauthorized  %+v", 401, o.Payload)
}

func (o *SidebarFetchEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}