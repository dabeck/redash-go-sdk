// Code generated by go-swagger; DO NOT EDIT.

package queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/dabeck/redash-go-sdk/gen/models"
)

// PostQueriesIDResultsReader is a Reader for the PostQueriesIDResults structure.
type PostQueriesIDResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQueriesIDResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQueriesIDResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostQueriesIDResultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostQueriesIDResultsOK creates a PostQueriesIDResultsOK with default headers values
func NewPostQueriesIDResultsOK() *PostQueriesIDResultsOK {
	return &PostQueriesIDResultsOK{}
}

/*
PostQueriesIDResultsOK describes a response with status code 200, with default header values.

Query is executed
*/
type PostQueriesIDResultsOK struct {
	Payload *models.JobResult
}

// IsSuccess returns true when this post queries Id results o k response has a 2xx status code
func (o *PostQueriesIDResultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post queries Id results o k response has a 3xx status code
func (o *PostQueriesIDResultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post queries Id results o k response has a 4xx status code
func (o *PostQueriesIDResultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post queries Id results o k response has a 5xx status code
func (o *PostQueriesIDResultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post queries Id results o k response a status code equal to that given
func (o *PostQueriesIDResultsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post queries Id results o k response
func (o *PostQueriesIDResultsOK) Code() int {
	return 200
}

func (o *PostQueriesIDResultsOK) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] postQueriesIdResultsOK  %+v", 200, o.Payload)
}

func (o *PostQueriesIDResultsOK) String() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] postQueriesIdResultsOK  %+v", 200, o.Payload)
}

func (o *PostQueriesIDResultsOK) GetPayload() *models.JobResult {
	return o.Payload
}

func (o *PostQueriesIDResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQueriesIDResultsDefault creates a PostQueriesIDResultsDefault with default headers values
func NewPostQueriesIDResultsDefault(code int) *PostQueriesIDResultsDefault {
	return &PostQueriesIDResultsDefault{
		_statusCode: code,
	}
}

/*
PostQueriesIDResultsDefault describes a response with status code -1, with default header values.

error
*/
type PostQueriesIDResultsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this post queries ID results default response has a 2xx status code
func (o *PostQueriesIDResultsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post queries ID results default response has a 3xx status code
func (o *PostQueriesIDResultsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post queries ID results default response has a 4xx status code
func (o *PostQueriesIDResultsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post queries ID results default response has a 5xx status code
func (o *PostQueriesIDResultsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post queries ID results default response a status code equal to that given
func (o *PostQueriesIDResultsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post queries ID results default response
func (o *PostQueriesIDResultsDefault) Code() int {
	return o._statusCode
}

func (o *PostQueriesIDResultsDefault) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] PostQueriesIDResults default  %+v", o._statusCode, o.Payload)
}

func (o *PostQueriesIDResultsDefault) String() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] PostQueriesIDResults default  %+v", o._statusCode, o.Payload)
}

func (o *PostQueriesIDResultsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostQueriesIDResultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostQueriesIDResultsBody post queries ID results body
swagger:model PostQueriesIDResultsBody
*/
type PostQueriesIDResultsBody struct {

	// apply auto limit
	// Required: true
	ApplyAutoLimit *bool `json:"apply_auto_limit"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// max age
	// Required: true
	MaxAge *int64 `json:"max_age"`

	// parameters
	Parameters interface{} `json:"parameters,omitempty"`
}

// Validate validates this post queries ID results body
func (o *PostQueriesIDResultsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateApplyAutoLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostQueriesIDResultsBody) validateApplyAutoLimit(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"apply_auto_limit", "body", o.ApplyAutoLimit); err != nil {
		return err
	}

	return nil
}

func (o *PostQueriesIDResultsBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *PostQueriesIDResultsBody) validateMaxAge(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"max_age", "body", o.MaxAge); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post queries ID results body based on context it is used
func (o *PostQueriesIDResultsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostQueriesIDResultsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostQueriesIDResultsBody) UnmarshalBinary(b []byte) error {
	var res PostQueriesIDResultsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
