// Code generated by go-swagger; DO NOT EDIT.

package queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dabeck/redash-go-sdk/gen/models"
)

// GetQueriesReader is a Reader for the GetQueries structure.
type GetQueriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetQueriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetQueriesOK creates a GetQueriesOK with default headers values
func NewGetQueriesOK() *GetQueriesOK {
	return &GetQueriesOK{}
}

/*
GetQueriesOK describes a response with status code 200, with default header values.

List queries
*/
type GetQueriesOK struct {
	Payload *models.QueryList
}

// IsSuccess returns true when this get queries o k response has a 2xx status code
func (o *GetQueriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get queries o k response has a 3xx status code
func (o *GetQueriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queries o k response has a 4xx status code
func (o *GetQueriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get queries o k response has a 5xx status code
func (o *GetQueriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get queries o k response a status code equal to that given
func (o *GetQueriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get queries o k response
func (o *GetQueriesOK) Code() int {
	return 200
}

func (o *GetQueriesOK) Error() string {
	return fmt.Sprintf("[GET /queries][%d] getQueriesOK  %+v", 200, o.Payload)
}

func (o *GetQueriesOK) String() string {
	return fmt.Sprintf("[GET /queries][%d] getQueriesOK  %+v", 200, o.Payload)
}

func (o *GetQueriesOK) GetPayload() *models.QueryList {
	return o.Payload
}

func (o *GetQueriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueriesDefault creates a GetQueriesDefault with default headers values
func NewGetQueriesDefault(code int) *GetQueriesDefault {
	return &GetQueriesDefault{
		_statusCode: code,
	}
}

/*
GetQueriesDefault describes a response with status code -1, with default header values.

error
*/
type GetQueriesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get queries default response has a 2xx status code
func (o *GetQueriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get queries default response has a 3xx status code
func (o *GetQueriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get queries default response has a 4xx status code
func (o *GetQueriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get queries default response has a 5xx status code
func (o *GetQueriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get queries default response a status code equal to that given
func (o *GetQueriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get queries default response
func (o *GetQueriesDefault) Code() int {
	return o._statusCode
}

func (o *GetQueriesDefault) Error() string {
	return fmt.Sprintf("[GET /queries][%d] GetQueries default  %+v", o._statusCode, o.Payload)
}

func (o *GetQueriesDefault) String() string {
	return fmt.Sprintf("[GET /queries][%d] GetQueries default  %+v", o._statusCode, o.Payload)
}

func (o *GetQueriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetQueriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
