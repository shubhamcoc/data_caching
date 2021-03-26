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

// PostLoadOKCode is the HTTP code returned for type PostLoadOK
const PostLoadOKCode int = 200

/*PostLoadOK OK

swagger:response postLoadOK
*/
type PostLoadOK struct {

	/*
	  In: Body
	*/
	Payload models.Validstringresponse `json:"body,omitempty"`
}

// NewPostLoadOK creates PostLoadOK with default headers values
func NewPostLoadOK() *PostLoadOK {

	return &PostLoadOK{}
}

// WithPayload adds the payload to the post load o k response
func (o *PostLoadOK) WithPayload(payload models.Validstringresponse) *PostLoadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post load o k response
func (o *PostLoadOK) SetPayload(payload models.Validstringresponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostLoadBadRequestCode is the HTTP code returned for type PostLoadBadRequest
const PostLoadBadRequestCode int = 400

/*PostLoadBadRequest The specified offset is not valid

swagger:response postLoadBadRequest
*/
type PostLoadBadRequest struct {
}

// NewPostLoadBadRequest creates PostLoadBadRequest with default headers values
func NewPostLoadBadRequest() *PostLoadBadRequest {

	return &PostLoadBadRequest{}
}

// WriteResponse to the client
func (o *PostLoadBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostLoadMethodNotAllowedCode is the HTTP code returned for type PostLoadMethodNotAllowed
const PostLoadMethodNotAllowedCode int = 405

/*PostLoadMethodNotAllowed Method not allowed

swagger:response postLoadMethodNotAllowed
*/
type PostLoadMethodNotAllowed struct {
}

// NewPostLoadMethodNotAllowed creates PostLoadMethodNotAllowed with default headers values
func NewPostLoadMethodNotAllowed() *PostLoadMethodNotAllowed {

	return &PostLoadMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PostLoadMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}