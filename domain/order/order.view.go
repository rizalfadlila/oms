package order

import (
	"context"
	"github.com/jatis/oms/models"
)

func (m *module) GetReportOrderByOrderID(ctx context.Context, id int64) (*models.ReportOrder, error) {
	var (
		result models.ReportOrder
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &result, queryGetReportOrder, id); err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *module) GetReportDetailOrderByOrderID(ctx context.Context, id int64) ([]models.ReportOrderDetail, error) {
	var (
		result = make([]models.ReportOrderDetail, 0)
	)

	if err := m.GetQueryerFromContext(ctx).SelectContext(ctx, &result, queryGetReportOrderDetail, id); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *module) GetIDByPO(ctx context.Context, number string) (*int64, error) {
	var (
		id int64
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &id, queryGetIDByPO, number); err != nil {
		return nil, err
	}

	return &id, nil
}
