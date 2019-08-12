// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

// UpdateRuleOKCode is the HTTP code returned for type UpdateRuleOK
const UpdateRuleOKCode int = 200

/*UpdateRuleOK Updated rule

swagger:response updateRuleOK
*/
type UpdateRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.Rule `json:"body,omitempty"`
}

// NewUpdateRuleOK creates UpdateRuleOK with default headers values
func NewUpdateRuleOK() *UpdateRuleOK {

	return &UpdateRuleOK{}
}

// WithPayload adds the payload to the update rule o k response
func (o *UpdateRuleOK) WithPayload(payload *models.Rule) *UpdateRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update rule o k response
func (o *UpdateRuleOK) SetPayload(payload *models.Rule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRuleBadRequestCode is the HTTP code returned for type UpdateRuleBadRequest
const UpdateRuleBadRequestCode int = 400

/*UpdateRuleBadRequest Bad request

swagger:response updateRuleBadRequest
*/
type UpdateRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewUpdateRuleBadRequest creates UpdateRuleBadRequest with default headers values
func NewUpdateRuleBadRequest() *UpdateRuleBadRequest {

	return &UpdateRuleBadRequest{}
}

// WithPayload adds the payload to the update rule bad request response
func (o *UpdateRuleBadRequest) WithPayload(payload *models.ErrorMessage) *UpdateRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update rule bad request response
func (o *UpdateRuleBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRuleUnauthorizedCode is the HTTP code returned for type UpdateRuleUnauthorized
const UpdateRuleUnauthorizedCode int = 401

/*UpdateRuleUnauthorized Unauthorized request

swagger:response updateRuleUnauthorized
*/
type UpdateRuleUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewUpdateRuleUnauthorized creates UpdateRuleUnauthorized with default headers values
func NewUpdateRuleUnauthorized() *UpdateRuleUnauthorized {

	return &UpdateRuleUnauthorized{}
}

// WithPayload adds the payload to the update rule unauthorized response
func (o *UpdateRuleUnauthorized) WithPayload(payload *models.ErrorMessage) *UpdateRuleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update rule unauthorized response
func (o *UpdateRuleUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRuleNotFoundCode is the HTTP code returned for type UpdateRuleNotFound
const UpdateRuleNotFoundCode int = 404

/*UpdateRuleNotFound Item not found

swagger:response updateRuleNotFound
*/
type UpdateRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewUpdateRuleNotFound creates UpdateRuleNotFound with default headers values
func NewUpdateRuleNotFound() *UpdateRuleNotFound {

	return &UpdateRuleNotFound{}
}

// WithPayload adds the payload to the update rule not found response
func (o *UpdateRuleNotFound) WithPayload(payload *models.ErrorMessage) *UpdateRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update rule not found response
func (o *UpdateRuleNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRuleConflictCode is the HTTP code returned for type UpdateRuleConflict
const UpdateRuleConflictCode int = 409

/*UpdateRuleConflict Conflicting item already exists

swagger:response updateRuleConflict
*/
type UpdateRuleConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewUpdateRuleConflict creates UpdateRuleConflict with default headers values
func NewUpdateRuleConflict() *UpdateRuleConflict {

	return &UpdateRuleConflict{}
}

// WithPayload adds the payload to the update rule conflict response
func (o *UpdateRuleConflict) WithPayload(payload *models.ErrorMessage) *UpdateRuleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update rule conflict response
func (o *UpdateRuleConflict) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRuleRequestEntityTooLargeCode is the HTTP code returned for type UpdateRuleRequestEntityTooLarge
const UpdateRuleRequestEntityTooLargeCode int = 413

/*UpdateRuleRequestEntityTooLarge Request entity too large

swagger:response updateRuleRequestEntityTooLarge
*/
type UpdateRuleRequestEntityTooLarge struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewUpdateRuleRequestEntityTooLarge creates UpdateRuleRequestEntityTooLarge with default headers values
func NewUpdateRuleRequestEntityTooLarge() *UpdateRuleRequestEntityTooLarge {

	return &UpdateRuleRequestEntityTooLarge{}
}

// WithPayload adds the payload to the update rule request entity too large response
func (o *UpdateRuleRequestEntityTooLarge) WithPayload(payload *models.ErrorMessage) *UpdateRuleRequestEntityTooLarge {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update rule request entity too large response
func (o *UpdateRuleRequestEntityTooLarge) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuleRequestEntityTooLarge) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(413)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRuleInternalServerErrorCode is the HTTP code returned for type UpdateRuleInternalServerError
const UpdateRuleInternalServerErrorCode int = 500

/*UpdateRuleInternalServerError Server error

swagger:response updateRuleInternalServerError
*/
type UpdateRuleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewUpdateRuleInternalServerError creates UpdateRuleInternalServerError with default headers values
func NewUpdateRuleInternalServerError() *UpdateRuleInternalServerError {

	return &UpdateRuleInternalServerError{}
}

// WithPayload adds the payload to the update rule internal server error response
func (o *UpdateRuleInternalServerError) WithPayload(payload *models.ErrorMessage) *UpdateRuleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update rule internal server error response
func (o *UpdateRuleInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}