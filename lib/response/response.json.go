package response

import (
	"encoding/json"
	"fmt"
	"github.com/jatis/oms/lib/custerr"
	"github.com/jatis/oms/lib/log"
	"net/http"
	"runtime"
)

type JSONResponse struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"`
}

func NewJSONResponse() *JSONResponse {
	return &JSONResponse{
		BaseResponse: BaseResponse{
			Log: map[string]interface{}{},
		},
	}
}

func (r *JSONResponse) APIStatusSuccess() *JSONResponse {
	r.StatusCode = http.StatusOK
	r.Code = fmt.Sprintf("%v000", r.StatusCode)
	r.Message = http.StatusText(r.StatusCode)
	return r
}

func (r *JSONResponse) setLog(key string, val interface{}) *JSONResponse {
	_, file, no, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"code":            r.Code,
		"err":             val,
		"function_caller": fmt.Sprintf("file %v line no %v", file, no),
	}).Errorln("Error API")
	r.Log[key] = val
	return r
}

func (r *JSONResponse) SetError(err *custerr.ErrChain) *JSONResponse {
	r.StatusCode = err.Code.GetStatusCode()
	r.Code = string(err.Code)
	r.Error = err.Error()

	r.setLog("error", err.Stacktrace)

	return r
}

func (r *JSONResponse) SetMessage(msg string) *JSONResponse {
	r.Message = msg
	return r
}

func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.Data = data
	return r
}

func (r *JSONResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)

	err := json.NewEncoder(w).Encode(r)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Errorln("[JSONResponse] Error encoding response")
	}
}
