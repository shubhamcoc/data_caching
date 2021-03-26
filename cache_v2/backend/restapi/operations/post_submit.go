// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostSubmitHandlerFunc turns a function with the right signature into a post submit handler
type PostSubmitHandlerFunc func(PostSubmitParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSubmitHandlerFunc) Handle(params PostSubmitParams) middleware.Responder {
	return fn(params)
}

// PostSubmitHandler interface for that can handle valid post submit params
type PostSubmitHandler interface {
	Handle(PostSubmitParams) middleware.Responder
}

// NewPostSubmit creates a new http.Handler for the post submit operation
func NewPostSubmit(ctx *middleware.Context, handler PostSubmitHandler) *PostSubmit {
	return &PostSubmit{Context: ctx, Handler: handler}
}

/* PostSubmit swagger:route POST /submit postSubmit

Save data in database

*/
type PostSubmit struct {
	Context *middleware.Context
	Handler PostSubmitHandler
}

func (o *PostSubmit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostSubmitParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
