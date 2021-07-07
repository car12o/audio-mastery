// Code generated by go-swagger; DO NOT EDIT.

package audios

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostAudioHandlerFunc turns a function with the right signature into a post audio handler
type PostAudioHandlerFunc func(PostAudioParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAudioHandlerFunc) Handle(params PostAudioParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostAudioHandler interface for that can handle valid post audio params
type PostAudioHandler interface {
	Handle(PostAudioParams, interface{}) middleware.Responder
}

// NewPostAudio creates a new http.Handler for the post audio operation
func NewPostAudio(ctx *middleware.Context, handler PostAudioHandler) *PostAudio {
	return &PostAudio{Context: ctx, Handler: handler}
}

/* PostAudio swagger:route POST /v1/audio audios postAudio

Creates an audio

*/
type PostAudio struct {
	Context *middleware.Context
	Handler PostAudioHandler
}

func (o *PostAudio) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAudioParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAudioBody post audio body
//
// swagger:model PostAudioBody
type PostAudioBody struct {

	// record
	// Required: true
	Record []strfmt.Base64 `json:"record"`
}

// Validate validates this post audio body
func (o *PostAudioBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecord(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAudioBody) validateRecord(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"record", "body", o.Record); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post audio body based on context it is used
func (o *PostAudioBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAudioBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAudioBody) UnmarshalBinary(b []byte) error {
	var res PostAudioBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAudioInternalServerErrorBody post audio internal server error body
//
// swagger:model PostAudioInternalServerErrorBody
type PostAudioInternalServerErrorBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this post audio internal server error body
func (o *PostAudioInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post audio internal server error body based on context it is used
func (o *PostAudioInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAudioInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAudioInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostAudioInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAudioNoContentBody post audio no content body
//
// swagger:model PostAudioNoContentBody
type PostAudioNoContentBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this post audio no content body
func (o *PostAudioNoContentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post audio no content body based on context it is used
func (o *PostAudioNoContentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAudioNoContentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAudioNoContentBody) UnmarshalBinary(b []byte) error {
	var res PostAudioNoContentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAudioUnauthorizedBody post audio unauthorized body
//
// swagger:model PostAudioUnauthorizedBody
type PostAudioUnauthorizedBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this post audio unauthorized body
func (o *PostAudioUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post audio unauthorized body based on context it is used
func (o *PostAudioUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAudioUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAudioUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostAudioUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}