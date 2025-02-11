// Code generated by go-swagger; DO NOT EDIT.

package visualizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/dabeck/redash-go-sdk/gen/models"
)

// PostVisualizationsReader is a Reader for the PostVisualizations structure.
type PostVisualizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVisualizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostVisualizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVisualizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVisualizationsOK creates a PostVisualizationsOK with default headers values
func NewPostVisualizationsOK() *PostVisualizationsOK {
	return &PostVisualizationsOK{}
}

/*
PostVisualizationsOK describes a response with status code 200, with default header values.

OK
*/
type PostVisualizationsOK struct {
	Payload *models.Visualization
}

// IsSuccess returns true when this post visualizations o k response has a 2xx status code
func (o *PostVisualizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post visualizations o k response has a 3xx status code
func (o *PostVisualizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post visualizations o k response has a 4xx status code
func (o *PostVisualizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post visualizations o k response has a 5xx status code
func (o *PostVisualizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post visualizations o k response a status code equal to that given
func (o *PostVisualizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post visualizations o k response
func (o *PostVisualizationsOK) Code() int {
	return 200
}

func (o *PostVisualizationsOK) Error() string {
	return fmt.Sprintf("[POST /visualizations][%d] postVisualizationsOK  %+v", 200, o.Payload)
}

func (o *PostVisualizationsOK) String() string {
	return fmt.Sprintf("[POST /visualizations][%d] postVisualizationsOK  %+v", 200, o.Payload)
}

func (o *PostVisualizationsOK) GetPayload() *models.Visualization {
	return o.Payload
}

func (o *PostVisualizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Visualization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVisualizationsDefault creates a PostVisualizationsDefault with default headers values
func NewPostVisualizationsDefault(code int) *PostVisualizationsDefault {
	return &PostVisualizationsDefault{
		_statusCode: code,
	}
}

/*
PostVisualizationsDefault describes a response with status code -1, with default header values.

error
*/
type PostVisualizationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this post visualizations default response has a 2xx status code
func (o *PostVisualizationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post visualizations default response has a 3xx status code
func (o *PostVisualizationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post visualizations default response has a 4xx status code
func (o *PostVisualizationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post visualizations default response has a 5xx status code
func (o *PostVisualizationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post visualizations default response a status code equal to that given
func (o *PostVisualizationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post visualizations default response
func (o *PostVisualizationsDefault) Code() int {
	return o._statusCode
}

func (o *PostVisualizationsDefault) Error() string {
	return fmt.Sprintf("[POST /visualizations][%d] PostVisualizations default  %+v", o._statusCode, o.Payload)
}

func (o *PostVisualizationsDefault) String() string {
	return fmt.Sprintf("[POST /visualizations][%d] PostVisualizations default  %+v", o._statusCode, o.Payload)
}

func (o *PostVisualizationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVisualizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostVisualizationsBody post visualizations body
swagger:model PostVisualizationsBody
*/
type PostVisualizationsBody struct {

	// name
	Name string `json:"name,omitempty"`

	// options
	Options interface{} `json:"options,omitempty"`

	// query id
	QueryID int64 `json:"query_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this post visualizations body
func (o *PostVisualizationsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post visualizations body based on context it is used
func (o *PostVisualizationsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostVisualizationsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostVisualizationsBody) UnmarshalBinary(b []byte) error {
	var res PostVisualizationsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
