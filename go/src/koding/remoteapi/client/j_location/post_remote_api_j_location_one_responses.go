package j_location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJLocationOneReader is a Reader for the PostRemoteAPIJLocationOne structure.
type PostRemoteAPIJLocationOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJLocationOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJLocationOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJLocationOneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJLocationOneOK creates a PostRemoteAPIJLocationOneOK with default headers values
func NewPostRemoteAPIJLocationOneOK() *PostRemoteAPIJLocationOneOK {
	return &PostRemoteAPIJLocationOneOK{}
}

/*PostRemoteAPIJLocationOneOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJLocationOneOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJLocationOneOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.one][%d] postRemoteApiJLocationOneOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJLocationOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJLocationOneUnauthorized creates a PostRemoteAPIJLocationOneUnauthorized with default headers values
func NewPostRemoteAPIJLocationOneUnauthorized() *PostRemoteAPIJLocationOneUnauthorized {
	return &PostRemoteAPIJLocationOneUnauthorized{}
}

/*PostRemoteAPIJLocationOneUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJLocationOneUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJLocationOneUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.one][%d] postRemoteApiJLocationOneUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJLocationOneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
