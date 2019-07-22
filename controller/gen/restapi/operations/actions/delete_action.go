// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

// DeleteActionHandlerFunc turns a function with the right signature into a delete action handler
type DeleteActionHandlerFunc func(DeleteActionParams, *models.Auth) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteActionHandlerFunc) Handle(params DeleteActionParams, principal *models.Auth) middleware.Responder {
	return fn(params, principal)
}

// DeleteActionHandler interface for that can handle valid delete action params
type DeleteActionHandler interface {
	Handle(DeleteActionParams, *models.Auth) middleware.Responder
}

// NewDeleteAction creates a new http.Handler for the delete action operation
func NewDeleteAction(ctx *middleware.Context, handler DeleteActionHandler) *DeleteAction {
	return &DeleteAction{Context: ctx, Handler: handler}
}

/*DeleteAction swagger:route DELETE /namespaces/{namespace}/actions/{actionName} Actions deleteAction

Delete an action

Delete an action

*/
type DeleteAction struct {
	Context *middleware.Context
	Handler DeleteActionHandler
}

func (o *DeleteAction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteActionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Auth
	if uprinc != nil {
		principal = uprinc.(*models.Auth) // this is really a models.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}