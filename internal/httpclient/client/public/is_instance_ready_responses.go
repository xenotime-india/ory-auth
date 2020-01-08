// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/internal/httpclient/models"
)

// IsInstanceReadyReader is a Reader for the IsInstanceReady structure.
type IsInstanceReadyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsInstanceReadyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsInstanceReadyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewIsInstanceReadyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIsInstanceReadyOK creates a IsInstanceReadyOK with default headers values
func NewIsInstanceReadyOK() *IsInstanceReadyOK {
	return &IsInstanceReadyOK{}
}

/*IsInstanceReadyOK handles this case with default header values.

healthStatus
*/
type IsInstanceReadyOK struct {
	Payload *models.HealthStatus
}

func (o *IsInstanceReadyOK) Error() string {
	return fmt.Sprintf("[GET /health/ready][%d] isInstanceReadyOK  %+v", 200, o.Payload)
}

func (o *IsInstanceReadyOK) GetPayload() *models.HealthStatus {
	return o.Payload
}

func (o *IsInstanceReadyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsInstanceReadyServiceUnavailable creates a IsInstanceReadyServiceUnavailable with default headers values
func NewIsInstanceReadyServiceUnavailable() *IsInstanceReadyServiceUnavailable {
	return &IsInstanceReadyServiceUnavailable{}
}

/*IsInstanceReadyServiceUnavailable handles this case with default header values.

healthNotReadyStatus
*/
type IsInstanceReadyServiceUnavailable struct {
	Payload *models.HealthNotReadyStatus
}

func (o *IsInstanceReadyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /health/ready][%d] isInstanceReadyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *IsInstanceReadyServiceUnavailable) GetPayload() *models.HealthNotReadyStatus {
	return o.Payload
}

func (o *IsInstanceReadyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthNotReadyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
