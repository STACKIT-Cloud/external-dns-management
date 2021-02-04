// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// GetServicesReader is a Reader for the GetServices structure.
type GetServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesOK creates a GetServicesOK with default headers values
func NewGetServicesOK() *GetServicesOK {
	return &GetServicesOK{}
}

/*GetServicesOK handles this case with default header values.

services
*/
type GetServicesOK struct {
	Payload []*models.Service
}

func (o *GetServicesOK) Error() string {
	return fmt.Sprintf("[GET /services/][%d] getServicesOK  %+v", 200, o.Payload)
}

func (o *GetServicesOK) GetPayload() []*models.Service {
	return o.Payload
}

func (o *GetServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesDefault creates a GetServicesDefault with default headers values
func NewGetServicesDefault(code int) *GetServicesDefault {
	return &GetServicesDefault{
		_statusCode: code,
	}
}

/*GetServicesDefault handles this case with default header values.

error
*/
type GetServicesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get services default response
func (o *GetServicesDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesDefault) Error() string {
	return fmt.Sprintf("[GET /services/][%d] getServices default  %+v", o._statusCode, o.Payload)
}

func (o *GetServicesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}