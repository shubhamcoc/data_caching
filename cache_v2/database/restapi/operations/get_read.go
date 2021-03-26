// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetReadHandlerFunc turns a function with the right signature into a get read handler
type GetReadHandlerFunc func(GetReadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReadHandlerFunc) Handle(params GetReadParams) middleware.Responder {
	return fn(params)
}

// GetReadHandler interface for that can handle valid get read params
type GetReadHandler interface {
	Handle(GetReadParams) middleware.Responder
}

// NewGetRead creates a new http.Handler for the get read operation
func NewGetRead(ctx *middleware.Context, handler GetReadHandler) *GetRead {
	return &GetRead{Context: ctx, Handler: handler}
}

/* GetRead swagger:route GET /read getRead

Return the query result

*/
type GetRead struct {
	Context *middleware.Context
	Handler GetReadHandler
}

func (o *GetRead) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReadParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}