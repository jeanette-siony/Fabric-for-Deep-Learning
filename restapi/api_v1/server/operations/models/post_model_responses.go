// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/IBM/FfDL/restapi/api_v1/restmodels"
)

// PostModelCreatedCode is the HTTP code returned for type PostModelCreated
const PostModelCreatedCode int = 201

/*PostModelCreated Deep learning model successfully accepted.

swagger:response postModelCreated
*/
type PostModelCreated struct {
	/*Location header containing the model id.
	  Required: true
	*/
	Location string `json:"Location"`

	/*
	  In: Body
	*/
	Payload *restmodels.BasicNewModel `json:"body,omitempty"`
}

// NewPostModelCreated creates PostModelCreated with default headers values
func NewPostModelCreated() *PostModelCreated {
	return &PostModelCreated{}
}

// WithLocation adds the location to the post model created response
func (o *PostModelCreated) WithLocation(location string) *PostModelCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the post model created response
func (o *PostModelCreated) SetLocation(location string) {
	o.Location = location
}

// WithPayload adds the payload to the post model created response
func (o *PostModelCreated) WithPayload(payload *restmodels.BasicNewModel) *PostModelCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post model created response
func (o *PostModelCreated) SetPayload(payload *restmodels.BasicNewModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostModelCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostModelBadRequestCode is the HTTP code returned for type PostModelBadRequest
const PostModelBadRequestCode int = 400

/*PostModelBadRequest Error in the the model_definition or manifest.

swagger:response postModelBadRequest
*/
type PostModelBadRequest struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewPostModelBadRequest creates PostModelBadRequest with default headers values
func NewPostModelBadRequest() *PostModelBadRequest {
	return &PostModelBadRequest{}
}

// WithPayload adds the payload to the post model bad request response
func (o *PostModelBadRequest) WithPayload(payload *restmodels.Error) *PostModelBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post model bad request response
func (o *PostModelBadRequest) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostModelBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostModelUnauthorizedCode is the HTTP code returned for type PostModelUnauthorized
const PostModelUnauthorizedCode int = 401

/*PostModelUnauthorized Unauthorized

swagger:response postModelUnauthorized
*/
type PostModelUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewPostModelUnauthorized creates PostModelUnauthorized with default headers values
func NewPostModelUnauthorized() *PostModelUnauthorized {
	return &PostModelUnauthorized{}
}

// WithPayload adds the payload to the post model unauthorized response
func (o *PostModelUnauthorized) WithPayload(payload *restmodels.Error) *PostModelUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post model unauthorized response
func (o *PostModelUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostModelUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
