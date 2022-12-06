package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/jatis/oms/lib/custerr"
	"github.com/jatis/oms/lib/log"
	"github.com/jatis/oms/lib/response"
	"github.com/julienschmidt/httprouter"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"runtime/debug"
)

func (mr *MyRouter) GET(path string, handle JSONHandle) {
	mr.JSONHandle(path, http.MethodGet, handle)
}

func (mr *MyRouter) POST(path string, handle JSONHandle) {
	mr.JSONHandle(path, http.MethodPost, handle)
}

func (mr *MyRouter) PUT(path string, handle JSONHandle) {
	mr.JSONHandle(path, http.MethodPut, handle)
}

func (mr *MyRouter) PATCH(path string, handle JSONHandle) {
	mr.JSONHandle(path, http.MethodPatch, handle)
}

func (mr *MyRouter) DELETE(path string, handle JSONHandle) {
	mr.JSONHandle(path, http.MethodDelete, handle)
}

func (mr *MyRouter) JSONHandle(path, method string, handle JSONHandle) {
	mr.registerHandler(path, method, jsonHandleToHttpRouterHandle(handle))
}

func jsonHandleToHttpRouterHandle(handle JSONHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, httpParamsKey{}, ps)
		ctx = context.WithValue(ctx, responseWriterKey{}, w)

		respChan := make(chan *response.JSONResponse)

		r = r.WithContext(ctx)

		go func() {
			respChan <- jsonPanicHandleWrapper(handle)(r)
		}()

		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				response.NewJSONResponse().
					SetMessage("Oops, something went wrong").
					SetError(custerr.ToCustErr(errors.New(string(custerr.ErrCodeTimeoutError)))).
					Send(w)
			}
		case resp := <-respChan:
			if resp != nil {
				span := trace.SpanFromContext(ctx)
				span.SetAttributes(httpResponseLogKey.String(fmt.Sprintf("%+v", resp.Log)))

				resp.Send(w)
			} else {
				httpDump := dumpRequest(r)

				log.WithFields(log.Fields{
					"dump": string(httpDump),
				}).Errorln("[Router] Nil response received from the handler")

				response.NewJSONResponse().
					SetMessage("Oops, something went wrong").
					SetError(custerr.ToCustErr(errors.New(string(custerr.ErrCodeInternalError)))).
					Send(w)
			}
		}
	}
}

func jsonPanicHandleWrapper(handle JSONHandle) JSONHandle {
	return func(r *http.Request) (resp *response.JSONResponse) {
		defer func() {
			if err := recover(); err != nil {
				httpDump := dumpRequest(r)
				stackTrace := string(debug.Stack())
				err := fmt.Sprintf("%+v", err)

				log.WithFields(log.Fields{
					"path":       r.URL.Path,
					"httpDump":   string(httpDump),
					"stackTrace": stackTrace,
					"error":      err,
				}).Errorln("[router/jsonPanicHandleWrapper] panic have occurred")

				resp = response.NewJSONResponse().
					SetMessage("Oops, something went wrong").
					SetError(custerr.New(fmt.Errorf("internal erver error")))
				return
			}
		}()

		resp = handle(r)
		return
	}
}
