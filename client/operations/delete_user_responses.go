// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/superkomputer/apiclient/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteUserAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserAccepted creates a DeleteUserAccepted with default headers values
func NewDeleteUserAccepted() *DeleteUserAccepted {
	return &DeleteUserAccepted{}
}

/*DeleteUserAccepted handles this case with default header values.

delete user task has been accepted
*/
type DeleteUserAccepted struct {
}

func (o *DeleteUserAccepted) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}][%d] deleteUserAccepted ", 202)
}

func (o *DeleteUserAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/*DeleteUserNotFound handles this case with default header values.

The user was not found
*/
type DeleteUserNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}][%d] deleteUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserDefault creates a DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	return &DeleteUserDefault{
		_statusCode: code,
	}
}

/*DeleteUserDefault handles this case with default header values.

Error
*/
type DeleteUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete user default response
func (o *DeleteUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}][%d] deleteUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
