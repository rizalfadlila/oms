package api

import (
	"github.com/jatis/oms/config"
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/lib/router"
)

type (
	Options struct {
		Service *service.Service
		Router  *router.MyRouter
	}

	API struct {
		usecase *service.Usecases
		config  config.RestConfig
		router  *router.MyRouter
	}
)

func New(o *Options) *API {
	return &API{
		config:  o.Service.Config.Server.Rest,
		usecase: o.Service.UseCases,

		router: o.Router,
	}
}

func (a *API) Register() {
	a.v1API()
}

func (a *API) v1API() {
	a.router.Group("/v1", func(r *router.MyRouter) {
		r.Group("/order", a.v1Order)
	})
}

func (a *API) v1Order(r *router.MyRouter) {
	r.GET("/:orderId/detail", a.GetDetailOrder)
}
