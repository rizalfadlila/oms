package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type OrderDefinition interface {
	StoreOrder(ctx context.Context, model *models.Order) error
	StoreOrderDetail(ctx context.Context, model *models.OrderDetail) error
	GetReportOrder(ctx context.Context) ([]models.ResponseReportOrder, error)
}
