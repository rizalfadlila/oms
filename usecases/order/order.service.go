package order

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *OrderUsecase) StoreOrder(ctx context.Context, model *models.Order) error {
	return u.orderManager.Store(ctx, model)
}

func (u *OrderUsecase) StoreOrderDetail(ctx context.Context, model *models.OrderDetail) error {
	return u.orderManager.StoreDetail(ctx, model)
}

func (u *OrderUsecase) GetReportOrderByID(ctx context.Context, id int64) (*models.ResponseReportOrder, error) {
	order, err := u.orderManager.GetReportOrderByOrderID(ctx, id)

	if err != nil {
		return nil, err
	}

	detail, err := u.orderManager.GetReportDetailOrderByOrderID(ctx, id)

	if err != nil {
		return nil, err
	}

	return models.ComposeReportData(*order, detail), nil
}

func (u *OrderUsecase) GetIDByPO(ctx context.Context, number string) (*int64, error) {
	return u.orderManager.GetIDByPO(ctx, number)
}
