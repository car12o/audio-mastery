// Code generated by go-swagger; DO NOT EDIT.

package audios

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListAudiosOKCode is the HTTP code returned for type ListAudiosOK
const ListAudiosOKCode int = 200

/*ListAudiosOK Ok

swagger:response listAudiosOK
*/
type ListAudiosOK struct {

	/*
	  In: Body
	*/
	Payload []*ListAudiosOKBodyItems0 `json:"body,omitempty"`
}

// NewListAudiosOK creates ListAudiosOK with default headers values
func NewListAudiosOK() *ListAudiosOK {

	return &ListAudiosOK{}
}

// WithPayload adds the payload to the list audios o k response
func (o *ListAudiosOK) WithPayload(payload []*ListAudiosOKBodyItems0) *ListAudiosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list audios o k response
func (o *ListAudiosOK) SetPayload(payload []*ListAudiosOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAudiosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ListAudiosOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListAudiosInternalServerErrorCode is the HTTP code returned for type ListAudiosInternalServerError
const ListAudiosInternalServerErrorCode int = 500

/*ListAudiosInternalServerError Internal Server Error

swagger:response listAudiosInternalServerError
*/
type ListAudiosInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ListAudiosInternalServerErrorBody `json:"body,omitempty"`
}

// NewListAudiosInternalServerError creates ListAudiosInternalServerError with default headers values
func NewListAudiosInternalServerError() *ListAudiosInternalServerError {

	return &ListAudiosInternalServerError{}
}

// WithPayload adds the payload to the list audios internal server error response
func (o *ListAudiosInternalServerError) WithPayload(payload *ListAudiosInternalServerErrorBody) *ListAudiosInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list audios internal server error response
func (o *ListAudiosInternalServerError) SetPayload(payload *ListAudiosInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAudiosInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
