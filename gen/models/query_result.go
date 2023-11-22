// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QueryResult query result
//
// swagger:model queryResult
type QueryResult struct {

	// query result
	QueryResult *QueryResultQueryResult `json:"query_result,omitempty"`
}

// Validate validates this query result
func (m *QueryResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueryResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResult) validateQueryResult(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryResult) { // not required
		return nil
	}

	if m.QueryResult != nil {
		if err := m.QueryResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this query result based on the context it is used
func (m *QueryResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQueryResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResult) contextValidateQueryResult(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryResult != nil {

		if swag.IsZero(m.QueryResult) { // not required
			return nil
		}

		if err := m.QueryResult.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResult) UnmarshalBinary(b []byte) error {
	var res QueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryResultQueryResult query result query result
//
// swagger:model QueryResultQueryResult
type QueryResultQueryResult struct {

	// data
	Data *QueryResultQueryResultData `json:"data,omitempty"`

	// data source id
	DataSourceID int64 `json:"data_source_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// query hash
	QueryHash string `json:"query_hash,omitempty"`

	// retrieved at
	// Format: date-time
	RetrievedAt strfmt.DateTime `json:"retrieved_at,omitempty"`
}

// Validate validates this query result query result
func (m *QueryResultQueryResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetrievedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResultQueryResult) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_result" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_result" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResultQueryResult) validateRetrievedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RetrievedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("query_result"+"."+"retrieved_at", "body", "date-time", m.RetrievedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this query result query result based on the context it is used
func (m *QueryResultQueryResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResultQueryResult) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {

		if swag.IsZero(m.Data) { // not required
			return nil
		}

		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_result" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_result" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResultQueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResultQueryResult) UnmarshalBinary(b []byte) error {
	var res QueryResultQueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryResultQueryResultData query result query result data
//
// swagger:model QueryResultQueryResultData
type QueryResultQueryResultData struct {

	// columns
	Columns []*QueryResultQueryResultDataColumnsItems0 `json:"columns"`

	// metadata
	Metadata *QueryResultQueryResultDataMetadata `json:"metadata,omitempty"`

	// rows
	Rows []interface{} `json:"rows"`
}

// Validate validates this query result query result data
func (m *QueryResultQueryResultData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResultQueryResultData) validateColumns(formats strfmt.Registry) error {
	if swag.IsZero(m.Columns) { // not required
		return nil
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("query_result" + "." + "data" + "." + "columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("query_result" + "." + "data" + "." + "columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryResultQueryResultData) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_result" + "." + "data" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_result" + "." + "data" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this query result query result data based on the context it is used
func (m *QueryResultQueryResultData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResultQueryResultData) contextValidateColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Columns); i++ {

		if m.Columns[i] != nil {

			if swag.IsZero(m.Columns[i]) { // not required
				return nil
			}

			if err := m.Columns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("query_result" + "." + "data" + "." + "columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("query_result" + "." + "data" + "." + "columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryResultQueryResultData) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_result" + "." + "data" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_result" + "." + "data" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResultQueryResultData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResultQueryResultData) UnmarshalBinary(b []byte) error {
	var res QueryResultQueryResultData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryResultQueryResultDataColumnsItems0 query result query result data columns items0
//
// swagger:model QueryResultQueryResultDataColumnsItems0
type QueryResultQueryResultDataColumnsItems0 struct {

	// friendly name
	FriendlyName string `json:"friendly_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this query result query result data columns items0
func (m *QueryResultQueryResultDataColumnsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query result query result data columns items0 based on context it is used
func (m *QueryResultQueryResultDataColumnsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryResultQueryResultDataColumnsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResultQueryResultDataColumnsItems0) UnmarshalBinary(b []byte) error {
	var res QueryResultQueryResultDataColumnsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryResultQueryResultDataMetadata query result query result data metadata
//
// swagger:model QueryResultQueryResultDataMetadata
type QueryResultQueryResultDataMetadata struct {

	// athena query id
	AthenaQueryID string `json:"athena_query_id,omitempty"`

	// data scanned
	DataScanned int64 `json:"data_scanned,omitempty"`
}

// Validate validates this query result query result data metadata
func (m *QueryResultQueryResultDataMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query result query result data metadata based on context it is used
func (m *QueryResultQueryResultDataMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryResultQueryResultDataMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResultQueryResultDataMetadata) UnmarshalBinary(b []byte) error {
	var res QueryResultQueryResultDataMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
