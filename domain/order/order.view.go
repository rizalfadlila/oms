package order

import (
	"context"
	"github.com/jatis/oms/errorx"
	"github.com/jatis/oms/models"
)

func (m *module) GetReportOrderByOrderID(ctx context.Context, id int64) (*models.ReportOrder, error) {
	var (
		result models.ReportOrder
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &result, queryGetReportOrder, id); err != nil {
		return nil, errorx.SqlError(err, errorx.SqlQuery)
	}

	return &result, nil
}

func (m *module) GetReportDetailOrderByOrderID(ctx context.Context, id int64) ([]models.ReportOrderDetail, error) {
	var (
		result = make([]models.ReportOrderDetail, 0)
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &result, queryGetReportOrderDetail, id); err != nil {
		return nil, errorx.SqlError(err, errorx.SqlQuery)
	}

	return result, nil
}
