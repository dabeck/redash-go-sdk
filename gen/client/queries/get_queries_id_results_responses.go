// Code generated by go-swagger; DO NOT EDIT.

package queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/recolabs/redash-go-sdk/gen/models"
)

// GetQueriesIDResultsReader is a Reader for the GetQueriesIDResults structure.
type GetQueriesIDResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueriesIDResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueriesIDResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetQueriesIDResultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetQueriesIDResultsOK creates a GetQueriesIDResultsOK with default headers values
func NewGetQueriesIDResultsOK() *GetQueriesIDResultsOK {
	return &GetQueriesIDResultsOK{}
}

/*
GetQueriesIDResultsOK describes a response with status code 200, with default header values.

Query is executed
*/
type GetQueriesIDResultsOK struct {
	Payload *models.QueryResult
}

// IsSuccess returns true when this get queries Id results o k response has a 2xx status code
func (o *GetQueriesIDResultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get queries Id results o k response has a 3xx status code
func (o *GetQueriesIDResultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queries Id results o k response has a 4xx status code
func (o *GetQueriesIDResultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get queries Id results o k response has a 5xx status code
func (o *GetQueriesIDResultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get queries Id results o k response a status code equal to that given
func (o *GetQueriesIDResultsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get queries Id results o k response
func (o *GetQueriesIDResultsOK) Code() int {
	return 200
}

func (o *GetQueriesIDResultsOK) Error() string {
	return fmt.Sprintf("[GET /queries/{id}/results][%d] getQueriesIdResultsOK  %+v", 200, o.Payload)
}

func (o *GetQueriesIDResultsOK) String() string {
	return fmt.Sprintf("[GET /queries/{id}/results][%d] getQueriesIdResultsOK  %+v", 200, o.Payload)
}

func (o *GetQueriesIDResultsOK) GetPayload() *models.QueryResult {
	return o.Payload
}

func (o *GetQueriesIDResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueriesIDResultsDefault creates a GetQueriesIDResultsDefault with default headers values
func NewGetQueriesIDResultsDefault(code int) *GetQueriesIDResultsDefault {
	return &GetQueriesIDResultsDefault{
		_statusCode: code,
	}
}

/*
GetQueriesIDResultsDefault describes a response with status code -1, with default header values.

error
*/
type GetQueriesIDResultsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get queries ID results default response has a 2xx status code
func (o *GetQueriesIDResultsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get queries ID results default response has a 3xx status code
func (o *GetQueriesIDResultsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get queries ID results default response has a 4xx status code
func (o *GetQueriesIDResultsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get queries ID results default response has a 5xx status code
func (o *GetQueriesIDResultsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get queries ID results default response a status code equal to that given
func (o *GetQueriesIDResultsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get queries ID results default response
func (o *GetQueriesIDResultsDefault) Code() int {
	return o._statusCode
}

func (o *GetQueriesIDResultsDefault) Error() string {
	return fmt.Sprintf("[GET /queries/{id}/results][%d] GetQueriesIDResults default  %+v", o._statusCode, o.Payload)
}

func (o *GetQueriesIDResultsDefault) String() string {
	return fmt.Sprintf("[GET /queries/{id}/results][%d] GetQueriesIDResults default  %+v", o._statusCode, o.Payload)
}

func (o *GetQueriesIDResultsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetQueriesIDResultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
