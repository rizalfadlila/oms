package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type OrderDefinition interface {
	StoreOrder(ctx context.Context, model *models.Order) error
	StoreOrderDetail(ctx context.Context, model *models.OrderDetail) error
	GetReportOrderByID(ctx context.Context, id int64) (*models.ResponseReportOrder, error)
}
