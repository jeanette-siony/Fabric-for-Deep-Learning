// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteEventEndpointHandlerFunc turns a function with the right signature into a delete event endpoint handler
type DeleteEventEndpointHandlerFunc func(DeleteEventEndpointParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteEventEndpointHandlerFunc) Handle(params DeleteEventEndpointParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteEventEndpointHandler interface for that can handle valid delete event endpoint params
type DeleteEventEndpointHandler interface {
	Handle(DeleteEventEndpointParams, interface{}) middleware.Responder
}

// NewDeleteEventEndpoint creates a new http.Handler for the delete event endpoint operation
func NewDeleteEventEndpoint(ctx *middleware.Context, handler DeleteEventEndpointHandler) *DeleteEventEndpoint {
	return &DeleteEventEndpoint{Context: ctx, Handler: handler}
}

/*DeleteEventEndpoint swagger:route DELETE /v1/models/{model_id}/events/{event_type}/{endpoint_id} Events deleteEventEndpoint

Deletes an event notification endpoint.

Deletes a specific event type's URL notification endpoint.

*/
type DeleteEventEndpoint struct {
	Context *middleware.Context
	Handler DeleteEventEndpointHandler
}

func (o *DeleteEventEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteEventEndpointParams()

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
