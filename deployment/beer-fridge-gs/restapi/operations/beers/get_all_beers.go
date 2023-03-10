// Code generated by go-swagger; DO NOT EDIT.

package beers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllBeersHandlerFunc turns a function with the right signature into a get all beers handler
type GetAllBeersHandlerFunc func(GetAllBeersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllBeersHandlerFunc) Handle(params GetAllBeersParams) middleware.Responder {
	return fn(params)
}

// GetAllBeersHandler interface for that can handle valid get all beers params
type GetAllBeersHandler interface {
	Handle(GetAllBeersParams) middleware.Responder
}

// NewGetAllBeers creates a new http.Handler for the get all beers operation
func NewGetAllBeers(ctx *middleware.Context, handler GetAllBeersHandler) *GetAllBeers {
	return &GetAllBeers{Context: ctx, Handler: handler}
}

/* GetAllBeers swagger:route GET /beers beers getAllBeers

GetAllBeers get all beers API

*/
type GetAllBeers struct {
	Context *middleware.Context
	Handler GetAllBeersHandler
}

func (o *GetAllBeers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllBeersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
