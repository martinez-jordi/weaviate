/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsListOKCode is the HTTP code returned for type WeaviateActionsListOK
const WeaviateActionsListOKCode int = 200

/*WeaviateActionsListOK Successful response.

swagger:response weaviateActionsListOK
*/
type WeaviateActionsListOK struct {

	/*
	  In: Body
	*/
	Payload *models.ActionsListResponse `json:"body,omitempty"`
}

// NewWeaviateActionsListOK creates WeaviateActionsListOK with default headers values
func NewWeaviateActionsListOK() *WeaviateActionsListOK {

	return &WeaviateActionsListOK{}
}

// WithPayload adds the payload to the weaviate actions list o k response
func (o *WeaviateActionsListOK) WithPayload(payload *models.ActionsListResponse) *WeaviateActionsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions list o k response
func (o *WeaviateActionsListOK) SetPayload(payload *models.ActionsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsListUnauthorizedCode is the HTTP code returned for type WeaviateActionsListUnauthorized
const WeaviateActionsListUnauthorizedCode int = 401

/*WeaviateActionsListUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsListUnauthorized
*/
type WeaviateActionsListUnauthorized struct {
}

// NewWeaviateActionsListUnauthorized creates WeaviateActionsListUnauthorized with default headers values
func NewWeaviateActionsListUnauthorized() *WeaviateActionsListUnauthorized {

	return &WeaviateActionsListUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsListForbiddenCode is the HTTP code returned for type WeaviateActionsListForbidden
const WeaviateActionsListForbiddenCode int = 403

/*WeaviateActionsListForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionsListForbidden
*/
type WeaviateActionsListForbidden struct {
}

// NewWeaviateActionsListForbidden creates WeaviateActionsListForbidden with default headers values
func NewWeaviateActionsListForbidden() *WeaviateActionsListForbidden {

	return &WeaviateActionsListForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsListNotFoundCode is the HTTP code returned for type WeaviateActionsListNotFound
const WeaviateActionsListNotFoundCode int = 404

/*WeaviateActionsListNotFound Successful query result but no resource was found.

swagger:response weaviateActionsListNotFound
*/
type WeaviateActionsListNotFound struct {
}

// NewWeaviateActionsListNotFound creates WeaviateActionsListNotFound with default headers values
func NewWeaviateActionsListNotFound() *WeaviateActionsListNotFound {

	return &WeaviateActionsListNotFound{}
}

// WriteResponse to the client
func (o *WeaviateActionsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// WeaviateActionsListInternalServerErrorCode is the HTTP code returned for type WeaviateActionsListInternalServerError
const WeaviateActionsListInternalServerErrorCode int = 500

/*WeaviateActionsListInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateActionsListInternalServerError
*/
type WeaviateActionsListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsListInternalServerError creates WeaviateActionsListInternalServerError with default headers values
func NewWeaviateActionsListInternalServerError() *WeaviateActionsListInternalServerError {

	return &WeaviateActionsListInternalServerError{}
}

// WithPayload adds the payload to the weaviate actions list internal server error response
func (o *WeaviateActionsListInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionsListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions list internal server error response
func (o *WeaviateActionsListInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
