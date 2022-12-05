package order

import (
	"context"
	"github.com/jatis/oms/models"
)

func (m *module) Store(ctx context.Context, model *models.Order) error {
	model.ID = m.GenerateUUID()
	_, err := m.GetExecerFromContext(ctx).NamedExecContext(ctx, queryStoreOrder, model)
	return err
}

func (m *module) StoreDetail(ctx context.Context, model *models.OrderDetail) error {
	model.ID = m.GenerateUUID()
	_, err := m.GetExecerFromContext(ctx).NamedExecContext(ctx, queryStoreOrderDetail, model)
	return err
}
