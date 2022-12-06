package api

import (
	"github.com/jatis/oms/lib/custerr"
	"github.com/jatis/oms/lib/response"
	"github.com/jatis/oms/lib/router"
	"net/http"
	"strconv"
)

func (a *API) GetDetailOrder(r *http.Request) *response.JSONResponse {
	var (
		ctx     = r.Context()
		message = "Failed to get notification"
	)

	orderID := router.GetHttpParam(ctx, "orderId")

	id, err := strconv.Atoi(orderID)
	if err != nil {
		return response.NewJSONResponse().
			SetError(custerr.New(err, custerr.WithErrCode(custerr.ErrCodeBadRequest))).
			SetMessage(message)
	}

	result, err := a.usecase.OrderUsecase.GetReportOrderByID(ctx, int64(id))

	if err != nil {
		if ok, v := custerr.IsCustErr(err); ok {
			return response.NewJSONResponse().
				SetError(v).
				SetMessage(message)
		}
		return response.NewJSONResponse().
			SetError(custerr.New(err)).
			SetMessage(message)
	}

	return response.NewJSONResponse().APIStatusSuccess().SetData(result)
}
