// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	//"mariadb_testing/models"
	"database/models"
)

// GetReadOKCode is the HTTP code returned for type GetReadOK
const GetReadOKCode int = 200

/*GetReadOK Data found

swagger:response getReadOK
*/
type GetReadOK struct {

	/*
	  In: Body
	*/
	Payload *models.Record `json:"body,omitempty"`
}

// NewGetReadOK creates GetReadOK with default headers values
func NewGetReadOK() *GetReadOK {

	return &GetReadOK{}
}

// WithPayload adds the payload to the get read o k response
func (o *GetReadOK) WithPayload(payload *models.Record) *GetReadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get read o k response
func (o *GetReadOK) SetPayload(payload *models.Record) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReadBadRequestCode is the HTTP code returned for type GetReadBadRequest
const GetReadBadRequestCode int = 400

/*GetReadBadRequest Key not found

swagger:response getReadBadRequest
*/
type GetReadBadRequest struct {
}

// NewGetReadBadRequest creates GetReadBadRequest with default headers values
func NewGetReadBadRequest() *GetReadBadRequest {

	return &GetReadBadRequest{}
}

// WriteResponse to the client
func (o *GetReadBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetReadMethodNotAllowedCode is the HTTP code returned for type GetReadMethodNotAllowed
const GetReadMethodNotAllowedCode int = 405

/*GetReadMethodNotAllowed Method not allowed

swagger:response getReadMethodNotAllowed
*/
type GetReadMethodNotAllowed struct {
}

// NewGetReadMethodNotAllowed creates GetReadMethodNotAllowed with default headers values
func NewGetReadMethodNotAllowed() *GetReadMethodNotAllowed {

	return &GetReadMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetReadMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
