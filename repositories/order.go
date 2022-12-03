package repositories

import (
	"context"
	"github.com/jatis/oms/models"
)

type OrderManager interface {
	Store(ctx context.Context, model *models.Order) error
	StoreDetail(ctx context.Context, model *models.OrderDetail) error

	GetReportOrder(ctx context.Context) (*models.ReportOrder, error)
	GetReportDetailOrderByOrderID(ctx context.Context, id int64) ([]models.ReportOrderDetail, error)
}
