// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostLoginOKCode is the HTTP code returned for type PostLoginOK
const PostLoginOKCode int = 200

/*PostLoginOK Ok

swagger:response postLoginOK
*/
type PostLoginOK struct {

	/*
	  In: Body
	*/
	Payload *PostLoginOKBody `json:"body,omitempty"`
}

// NewPostLoginOK creates PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {

	return &PostLoginOK{}
}

// WithPayload adds the payload to the post login o k response
func (o *PostLoginOK) WithPayload(payload *PostLoginOKBody) *PostLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login o k response
func (o *PostLoginOK) SetPayload(payload *PostLoginOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLoginBadRequestCode is the HTTP code returned for type PostLoginBadRequest
const PostLoginBadRequestCode int = 400

/*PostLoginBadRequest Bad Request

swagger:response postLoginBadRequest
*/
type PostLoginBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostLoginBadRequestBody `json:"body,omitempty"`
}

// NewPostLoginBadRequest creates PostLoginBadRequest with default headers values
func NewPostLoginBadRequest() *PostLoginBadRequest {

	return &PostLoginBadRequest{}
}

// WithPayload adds the payload to the post login bad request response
func (o *PostLoginBadRequest) WithPayload(payload *PostLoginBadRequestBody) *PostLoginBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login bad request response
func (o *PostLoginBadRequest) SetPayload(payload *PostLoginBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLoginInternalServerErrorCode is the HTTP code returned for type PostLoginInternalServerError
const PostLoginInternalServerErrorCode int = 500

/*PostLoginInternalServerError Internal Server Error

swagger:response postLoginInternalServerError
*/
type PostLoginInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostLoginInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostLoginInternalServerError creates PostLoginInternalServerError with default headers values
func NewPostLoginInternalServerError() *PostLoginInternalServerError {

	return &PostLoginInternalServerError{}
}

// WithPayload adds the payload to the post login internal server error response
func (o *PostLoginInternalServerError) WithPayload(payload *PostLoginInternalServerErrorBody) *PostLoginInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login internal server error response
func (o *PostLoginInternalServerError) SetPayload(payload *PostLoginInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
