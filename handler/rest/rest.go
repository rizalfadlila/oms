package rest

import (
	"fmt"
	"github.com/jatis/oms/handler/rest/api"
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/lib/httputil"
	"github.com/jatis/oms/lib/router"
	"log"
	"time"
)

type (
	Opts struct {
		Service *service.Service
	}

	Handler struct {
		options *Opts
		err     chan error
		router  *router.MyRouter
	}

	RestHandler interface {
		Run()
		Error() <-chan error
	}
)

func New(o *Opts) RestHandler {
	routerTimeout, err := time.ParseDuration(o.Service.Config.Server.Rest.RouterTimeout)
	if err != nil {
		log.Fatalf("error when parse gracefulTimeout: %v", err)
	}

	r := router.New(&router.Options{Timeout: routerTimeout})

	handler := &Handler{
		options: o,
		router:  r,
	}

	api.New(&api.Options{
		Service: o.Service,
		Router:  r,
	}).Register()

	return handler
}

func (h *Handler) Run() {
	cfg := h.options.Service.Config.Server

	gracefulTimeout, err := time.ParseDuration(cfg.Rest.GracefulTimeout)
	if err != nil {
		h.err <- fmt.Errorf("error when parse gracefulTimeout: %v", err)
	}

	readTimeout, err := time.ParseDuration(cfg.Rest.ReadTimeout)
	if err != nil {
		h.err <- fmt.Errorf("error when parse readTimeout: %v", err)
	}

	writeTimeout, err := time.ParseDuration(cfg.Rest.WriteTimeout)
	if err != nil {
		h.err <- fmt.Errorf("error when parse writeTimeout: %v", err)
	}

	h.err <- httputil.HTTPServe(
		fmt.Sprintf(":%v", cfg.Rest.Port),
		h.router,
		gracefulTimeout,
		readTimeout,
		writeTimeout,
	)
}

func (h *Handler) Error() <-chan error {
	return h.err
}
