// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/recolabs/redash-go-sdk/gen/models"
)

// GetUsersIDReader is a Reader for the GetUsersID structure.
type GetUsersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUsersIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsersIDOK creates a GetUsersIDOK with default headers values
func NewGetUsersIDOK() *GetUsersIDOK {
	return &GetUsersIDOK{}
}

/*
GetUsersIDOK describes a response with status code 200, with default header values.

Get user by ID
*/
type GetUsersIDOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get users Id o k response has a 2xx status code
func (o *GetUsersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users Id o k response has a 3xx status code
func (o *GetUsersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users Id o k response has a 4xx status code
func (o *GetUsersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users Id o k response has a 5xx status code
func (o *GetUsersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users Id o k response a status code equal to that given
func (o *GetUsersIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users Id o k response
func (o *GetUsersIDOK) Code() int {
	return 200
}

func (o *GetUsersIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdOK  %+v", 200, o.Payload)
}

func (o *GetUsersIDOK) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdOK  %+v", 200, o.Payload)
}

func (o *GetUsersIDOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUsersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersIDDefault creates a GetUsersIDDefault with default headers values
func NewGetUsersIDDefault(code int) *GetUsersIDDefault {
	return &GetUsersIDDefault{
		_statusCode: code,
	}
}

/*
GetUsersIDDefault describes a response with status code -1, with default header values.

error
*/
type GetUsersIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get users ID default response has a 2xx status code
func (o *GetUsersIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get users ID default response has a 3xx status code
func (o *GetUsersIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get users ID default response has a 4xx status code
func (o *GetUsersIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get users ID default response has a 5xx status code
func (o *GetUsersIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get users ID default response a status code equal to that given
func (o *GetUsersIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get users ID default response
func (o *GetUsersIDDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersIDDefault) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] GetUsersID default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersIDDefault) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] GetUsersID default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsersIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
