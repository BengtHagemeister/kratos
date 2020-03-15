// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// GetSelfServiceBrowserProfileManagementRequestReader is a Reader for the GetSelfServiceBrowserProfileManagementRequest structure.
type GetSelfServiceBrowserProfileManagementRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceBrowserProfileManagementRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceBrowserProfileManagementRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSelfServiceBrowserProfileManagementRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSelfServiceBrowserProfileManagementRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSelfServiceBrowserProfileManagementRequestGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceBrowserProfileManagementRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSelfServiceBrowserProfileManagementRequestOK creates a GetSelfServiceBrowserProfileManagementRequestOK with default headers values
func NewGetSelfServiceBrowserProfileManagementRequestOK() *GetSelfServiceBrowserProfileManagementRequestOK {
	return &GetSelfServiceBrowserProfileManagementRequestOK{}
}

/*GetSelfServiceBrowserProfileManagementRequestOK handles this case with default header values.

profileManagementRequest
*/
type GetSelfServiceBrowserProfileManagementRequestOK struct {
	Payload *models.ProfileManagementRequest
}

func (o *GetSelfServiceBrowserProfileManagementRequestOK) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/profile][%d] getSelfServiceBrowserProfileManagementRequestOK  %+v", 200, o.Payload)
}

func (o *GetSelfServiceBrowserProfileManagementRequestOK) GetPayload() *models.ProfileManagementRequest {
	return o.Payload
}

func (o *GetSelfServiceBrowserProfileManagementRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProfileManagementRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserProfileManagementRequestForbidden creates a GetSelfServiceBrowserProfileManagementRequestForbidden with default headers values
func NewGetSelfServiceBrowserProfileManagementRequestForbidden() *GetSelfServiceBrowserProfileManagementRequestForbidden {
	return &GetSelfServiceBrowserProfileManagementRequestForbidden{}
}

/*GetSelfServiceBrowserProfileManagementRequestForbidden handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserProfileManagementRequestForbidden struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserProfileManagementRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/profile][%d] getSelfServiceBrowserProfileManagementRequestForbidden  %+v", 403, o.Payload)
}

func (o *GetSelfServiceBrowserProfileManagementRequestForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserProfileManagementRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserProfileManagementRequestNotFound creates a GetSelfServiceBrowserProfileManagementRequestNotFound with default headers values
func NewGetSelfServiceBrowserProfileManagementRequestNotFound() *GetSelfServiceBrowserProfileManagementRequestNotFound {
	return &GetSelfServiceBrowserProfileManagementRequestNotFound{}
}

/*GetSelfServiceBrowserProfileManagementRequestNotFound handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserProfileManagementRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserProfileManagementRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/profile][%d] getSelfServiceBrowserProfileManagementRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetSelfServiceBrowserProfileManagementRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserProfileManagementRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserProfileManagementRequestGone creates a GetSelfServiceBrowserProfileManagementRequestGone with default headers values
func NewGetSelfServiceBrowserProfileManagementRequestGone() *GetSelfServiceBrowserProfileManagementRequestGone {
	return &GetSelfServiceBrowserProfileManagementRequestGone{}
}

/*GetSelfServiceBrowserProfileManagementRequestGone handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserProfileManagementRequestGone struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserProfileManagementRequestGone) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/profile][%d] getSelfServiceBrowserProfileManagementRequestGone  %+v", 410, o.Payload)
}

func (o *GetSelfServiceBrowserProfileManagementRequestGone) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserProfileManagementRequestGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserProfileManagementRequestInternalServerError creates a GetSelfServiceBrowserProfileManagementRequestInternalServerError with default headers values
func NewGetSelfServiceBrowserProfileManagementRequestInternalServerError() *GetSelfServiceBrowserProfileManagementRequestInternalServerError {
	return &GetSelfServiceBrowserProfileManagementRequestInternalServerError{}
}

/*GetSelfServiceBrowserProfileManagementRequestInternalServerError handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserProfileManagementRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserProfileManagementRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/profile][%d] getSelfServiceBrowserProfileManagementRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSelfServiceBrowserProfileManagementRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserProfileManagementRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
