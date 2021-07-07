// Code generated by go-swagger; DO NOT EDIT.

package audios

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteAudioNoContentCode is the HTTP code returned for type DeleteAudioNoContent
const DeleteAudioNoContentCode int = 204

/*DeleteAudioNoContent No Content

swagger:response deleteAudioNoContent
*/
type DeleteAudioNoContent struct {

	/*
	  In: Body
	*/
	Payload *DeleteAudioNoContentBody `json:"body,omitempty"`
}

// NewDeleteAudioNoContent creates DeleteAudioNoContent with default headers values
func NewDeleteAudioNoContent() *DeleteAudioNoContent {

	return &DeleteAudioNoContent{}
}

// WithPayload adds the payload to the delete audio no content response
func (o *DeleteAudioNoContent) WithPayload(payload *DeleteAudioNoContentBody) *DeleteAudioNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete audio no content response
func (o *DeleteAudioNoContent) SetPayload(payload *DeleteAudioNoContentBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAudioNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAudioBadRequestCode is the HTTP code returned for type DeleteAudioBadRequest
const DeleteAudioBadRequestCode int = 400

/*DeleteAudioBadRequest Bad Request

swagger:response deleteAudioBadRequest
*/
type DeleteAudioBadRequest struct {

	/*
	  In: Body
	*/
	Payload *DeleteAudioBadRequestBody `json:"body,omitempty"`
}

// NewDeleteAudioBadRequest creates DeleteAudioBadRequest with default headers values
func NewDeleteAudioBadRequest() *DeleteAudioBadRequest {

	return &DeleteAudioBadRequest{}
}

// WithPayload adds the payload to the delete audio bad request response
func (o *DeleteAudioBadRequest) WithPayload(payload *DeleteAudioBadRequestBody) *DeleteAudioBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete audio bad request response
func (o *DeleteAudioBadRequest) SetPayload(payload *DeleteAudioBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAudioBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAudioUnauthorizedCode is the HTTP code returned for type DeleteAudioUnauthorized
const DeleteAudioUnauthorizedCode int = 401

/*DeleteAudioUnauthorized Unauthorized

swagger:response deleteAudioUnauthorized
*/
type DeleteAudioUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *DeleteAudioUnauthorizedBody `json:"body,omitempty"`
}

// NewDeleteAudioUnauthorized creates DeleteAudioUnauthorized with default headers values
func NewDeleteAudioUnauthorized() *DeleteAudioUnauthorized {

	return &DeleteAudioUnauthorized{}
}

// WithPayload adds the payload to the delete audio unauthorized response
func (o *DeleteAudioUnauthorized) WithPayload(payload *DeleteAudioUnauthorizedBody) *DeleteAudioUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete audio unauthorized response
func (o *DeleteAudioUnauthorized) SetPayload(payload *DeleteAudioUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAudioUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAudioInternalServerErrorCode is the HTTP code returned for type DeleteAudioInternalServerError
const DeleteAudioInternalServerErrorCode int = 500

/*DeleteAudioInternalServerError Internal Server Error

swagger:response deleteAudioInternalServerError
*/
type DeleteAudioInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *DeleteAudioInternalServerErrorBody `json:"body,omitempty"`
}

// NewDeleteAudioInternalServerError creates DeleteAudioInternalServerError with default headers values
func NewDeleteAudioInternalServerError() *DeleteAudioInternalServerError {

	return &DeleteAudioInternalServerError{}
}

// WithPayload adds the payload to the delete audio internal server error response
func (o *DeleteAudioInternalServerError) WithPayload(payload *DeleteAudioInternalServerErrorBody) *DeleteAudioInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete audio internal server error response
func (o *DeleteAudioInternalServerError) SetPayload(payload *DeleteAudioInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAudioInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}