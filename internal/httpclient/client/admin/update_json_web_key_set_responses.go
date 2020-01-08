// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/internal/httpclient/models"
)

// UpdateJSONWebKeySetReader is a Reader for the UpdateJSONWebKeySet structure.
type UpdateJSONWebKeySetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateJSONWebKeySetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateJSONWebKeySetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateJSONWebKeySetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateJSONWebKeySetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateJSONWebKeySetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateJSONWebKeySetOK creates a UpdateJSONWebKeySetOK with default headers values
func NewUpdateJSONWebKeySetOK() *UpdateJSONWebKeySetOK {
	return &UpdateJSONWebKeySetOK{}
}

/*UpdateJSONWebKeySetOK handles this case with default header values.

JSONWebKeySet
*/
type UpdateJSONWebKeySetOK struct {
	Payload *models.JSONWebKeySet
}

func (o *UpdateJSONWebKeySetOK) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}][%d] updateJsonWebKeySetOK  %+v", 200, o.Payload)
}

func (o *UpdateJSONWebKeySetOK) GetPayload() *models.JSONWebKeySet {
	return o.Payload
}

func (o *UpdateJSONWebKeySetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONWebKeySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJSONWebKeySetUnauthorized creates a UpdateJSONWebKeySetUnauthorized with default headers values
func NewUpdateJSONWebKeySetUnauthorized() *UpdateJSONWebKeySetUnauthorized {
	return &UpdateJSONWebKeySetUnauthorized{}
}

/*UpdateJSONWebKeySetUnauthorized handles this case with default header values.

genericError
*/
type UpdateJSONWebKeySetUnauthorized struct {
	Payload *models.GenericError
}

func (o *UpdateJSONWebKeySetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}][%d] updateJsonWebKeySetUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateJSONWebKeySetUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateJSONWebKeySetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJSONWebKeySetForbidden creates a UpdateJSONWebKeySetForbidden with default headers values
func NewUpdateJSONWebKeySetForbidden() *UpdateJSONWebKeySetForbidden {
	return &UpdateJSONWebKeySetForbidden{}
}

/*UpdateJSONWebKeySetForbidden handles this case with default header values.

genericError
*/
type UpdateJSONWebKeySetForbidden struct {
	Payload *models.GenericError
}

func (o *UpdateJSONWebKeySetForbidden) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}][%d] updateJsonWebKeySetForbidden  %+v", 403, o.Payload)
}

func (o *UpdateJSONWebKeySetForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateJSONWebKeySetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJSONWebKeySetInternalServerError creates a UpdateJSONWebKeySetInternalServerError with default headers values
func NewUpdateJSONWebKeySetInternalServerError() *UpdateJSONWebKeySetInternalServerError {
	return &UpdateJSONWebKeySetInternalServerError{}
}

/*UpdateJSONWebKeySetInternalServerError handles this case with default header values.

genericError
*/
type UpdateJSONWebKeySetInternalServerError struct {
	Payload *models.GenericError
}

func (o *UpdateJSONWebKeySetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}][%d] updateJsonWebKeySetInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateJSONWebKeySetInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateJSONWebKeySetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
