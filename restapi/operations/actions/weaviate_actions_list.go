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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateActionsListHandlerFunc turns a function with the right signature into a weaviate actions list handler
type WeaviateActionsListHandlerFunc func(context.Context, WeaviateActionsListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsListHandlerFunc) Handle(ctx context.Context, params WeaviateActionsListParams) middleware.Responder {
	return fn(ctx, params)
}

// WeaviateActionsListHandler interface for that can handle valid weaviate actions list params
type WeaviateActionsListHandler interface {
	Handle(context.Context, WeaviateActionsListParams) middleware.Responder
}

// NewWeaviateActionsList creates a new http.Handler for the weaviate actions list operation
func NewWeaviateActionsList(ctx *middleware.Context, handler WeaviateActionsListHandler) *WeaviateActionsList {
	return &WeaviateActionsList{Context: ctx, Handler: handler}
}

/*WeaviateActionsList swagger:route GET /actions actions weaviateActionsList

Get a list of Actions related to this key.

Lists all Actions in reverse order of creation, owned by the user that belongs to the used token.

*/
type WeaviateActionsList struct {
	Context *middleware.Context
	Handler WeaviateActionsListHandler
}

func (o *WeaviateActionsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(r.Context(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
