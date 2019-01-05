/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaActionsUpdateHandlerFunc turns a function with the right signature into a weaviate schema actions update handler
type WeaviateSchemaActionsUpdateHandlerFunc func(WeaviateSchemaActionsUpdateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateSchemaActionsUpdateHandlerFunc) Handle(params WeaviateSchemaActionsUpdateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateSchemaActionsUpdateHandler interface for that can handle valid weaviate schema actions update params
type WeaviateSchemaActionsUpdateHandler interface {
	Handle(WeaviateSchemaActionsUpdateParams, interface{}) middleware.Responder
}

// NewWeaviateSchemaActionsUpdate creates a new http.Handler for the weaviate schema actions update operation
func NewWeaviateSchemaActionsUpdate(ctx *middleware.Context, handler WeaviateSchemaActionsUpdateHandler) *WeaviateSchemaActionsUpdate {
	return &WeaviateSchemaActionsUpdate{Context: ctx, Handler: handler}
}

/*WeaviateSchemaActionsUpdate swagger:route PUT /schema/actions/{className} schema weaviateSchemaActionsUpdate

Rename, or replace the keywords of the Action.

*/
type WeaviateSchemaActionsUpdate struct {
	Context *middleware.Context
	Handler WeaviateSchemaActionsUpdateHandler
}

func (o *WeaviateSchemaActionsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateSchemaActionsUpdateParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WeaviateSchemaActionsUpdateBody weaviate schema actions update body
// swagger:model WeaviateSchemaActionsUpdateBody
type WeaviateSchemaActionsUpdateBody struct {

	// keywords
	Keywords models.SemanticSchemaKeywords `json:"keywords,omitempty"`

	// The new name of the Action.
	NewName string `json:"newName,omitempty"`
}

// Validate validates this weaviate schema actions update body
func (o *WeaviateSchemaActionsUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateSchemaActionsUpdateBody) validateKeywords(formats strfmt.Registry) error {

	if swag.IsZero(o.Keywords) { // not required
		return nil
	}

	if err := o.Keywords.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "keywords")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateSchemaActionsUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateSchemaActionsUpdateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateSchemaActionsUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
