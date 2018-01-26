// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss-labs/taco/authorization"
)

// RetrieveResourceHandlerFunc turns a function with the right signature into a retrieve resource handler
type RetrieveResourceHandlerFunc func(RetrieveResourceParams, *authorization.Agent) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveResourceHandlerFunc) Handle(params RetrieveResourceParams, principal *authorization.Agent) middleware.Responder {
	return fn(params, principal)
}

// RetrieveResourceHandler interface for that can handle valid retrieve resource params
type RetrieveResourceHandler interface {
	Handle(RetrieveResourceParams, *authorization.Agent) middleware.Responder
}

// NewRetrieveResource creates a new http.Handler for the retrieve resource operation
func NewRetrieveResource(ctx *middleware.Context, handler RetrieveResourceHandler) *RetrieveResource {
	return &RetrieveResource{Context: ctx, Handler: handler}
}

/*RetrieveResource swagger:route GET /resource/{ID} retrieveResource

Retrieve TACO Resource Metadata.

Retrieves the metadata (as JSON-LD following our SDR3 MAP v.1) for an existing TACO resource (Collection, Digital Repository Object, File metadata object [not binary] or subclass of those). The resource is identified by the TACO identifier.

*/
type RetrieveResource struct {
	Context *middleware.Context
	Handler RetrieveResourceHandler
}

func (o *RetrieveResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRetrieveResourceParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *authorization.Agent
	if uprinc != nil {
		principal = uprinc.(*authorization.Agent) // this is really a authorization.Agent, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
