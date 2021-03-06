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

// FetchUserClustersReader is a Reader for the FetchUserClusters structure.
type FetchUserClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchUserClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFetchUserClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFetchUserClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFetchUserClustersOK creates a FetchUserClustersOK with default headers values
func NewFetchUserClustersOK() *FetchUserClustersOK {
	return &FetchUserClustersOK{}
}

/*FetchUserClustersOK handles this case with default header values.

200 response with the list of using clusters of user
*/
type FetchUserClustersOK struct {
	Payload []*models.Cluster
}

func (o *FetchUserClustersOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/clusters][%d] fetchUserClustersOK  %+v", 200, o.Payload)
}

func (o *FetchUserClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchUserClustersDefault creates a FetchUserClustersDefault with default headers values
func NewFetchUserClustersDefault(code int) *FetchUserClustersDefault {
	return &FetchUserClustersDefault{
		_statusCode: code,
	}
}

/*FetchUserClustersDefault handles this case with default header values.

Error
*/
type FetchUserClustersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the fetch user clusters default response
func (o *FetchUserClustersDefault) Code() int {
	return o._statusCode
}

func (o *FetchUserClustersDefault) Error() string {
	return fmt.Sprintf("[GET /users/{username}/clusters][%d] fetchUserClusters default  %+v", o._statusCode, o.Payload)
}

func (o *FetchUserClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
