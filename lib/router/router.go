package router

import (
	"context"
	"fmt"
	"github.com/felixge/httpsnoop"
	"github.com/jatis/oms/lib/log"
	"github.com/jatis/oms/lib/response"
	"github.com/julienschmidt/httprouter"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
	"net/http/httputil"
	"time"
)

type (
	MyRouter struct {
		Httprouter     *httprouter.Router
		WrappedHandler http.Handler
		Options        *Options
	}

	Options struct {
		Prefix  string
		Timeout time.Duration
	}

	captureConfig struct {
		captureHandler bool
	}

	httpParamsKey     struct{}
	responseWriterKey struct{}
	captureHandlerKey struct{}

	JSONHandle func(*http.Request) *response.JSONResponse
)

var (
	httpResponseLogKey = attribute.Key("rest.response.log")
)

func New(o *Options) *MyRouter {
	myrouter := &MyRouter{
		Options: o,
	}
	myrouter.Httprouter = httprouter.New()
	return myrouter
}

func (mr *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = httpsnoop.CaptureMetrics(mr.Httprouter, w, r.WithContext(r.Context()))
}

func (mr *MyRouter) Group(path string, fn func(r *MyRouter)) {
	sr := &MyRouter{
		Options: &Options{
			Prefix:  mr.Options.Prefix + path,
			Timeout: mr.Options.Timeout,
		},
		Httprouter: mr.Httprouter,
	}
	fn(sr)
}

func GetHttpParam(ctx context.Context, name string) string {
	ps := ctx.Value(httpParamsKey{}).(httprouter.Params)
	return ps.ByName(name)
}

func (mr *MyRouter) registerHandler(path, method string, handle httprouter.Handle) {
	fullPath := mr.Options.Prefix + path
	log.Println(fullPath)
	mr.Httprouter.Handle(method, fullPath, mr.handlePath(fullPath, handle))
}

func (mr *MyRouter) handlePath(fullPath string, handle httprouter.Handle) httprouter.Handle {
	var captureHandler = true

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx, cancel := context.WithTimeout(r.Context(), mr.Options.Timeout)
		defer cancel()

		ctx = context.WithValue(ctx, httpParamsKey{}, ps)

		if value, ok := ctx.Value(captureHandlerKey{}).(*captureConfig); ok {
			value.captureHandler = captureHandler
		}

		r.Header.Set("routePath", fullPath)

		r = r.WithContext(ctx)

		handle(w, r, ps)
	}
}

func dumpRequest(r *http.Request) []byte {
	httpDump, err := httputil.DumpRequest(r, true)
	if err == nil {
		return httpDump
	}

	log.WithFields(log.Fields{
		"url":    r.URL,
		"method": r.Method,
		"header": fmt.Sprintf("%+v", r.Header),
		"err":    err,
	}).Debugln("[Router] Failed to dump request with body, re-attempting to dump request without body")

	httpDump, err = httputil.DumpRequest(r, false)
	if err == nil {
		return httpDump
	}

	log.WithFields(log.Fields{
		"url":    r.URL,
		"method": r.Method,
		"header": fmt.Sprintf("%+v", r.Header),
		"err":    err,
	}).Infoln("[Router] Failed to dump request")
	return nil
}
