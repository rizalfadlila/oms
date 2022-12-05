package order

import (
	"context"
	"github.com/jatis/oms/errorx"
	"github.com/jatis/oms/models"
)

func (m *module) Store(ctx context.Context, model *models.Order) error {
	_, err := m.GetExecerFromContext(ctx).NamedExecContext(ctx, queryStoreOrder, model)
	return errorx.SqlError(err, errorx.SqlTransaction)
}

func (m *module) StoreDetail(ctx context.Context, model *models.OrderDetail) error {
	_, err := m.GetExecerFromContext(ctx).NamedExecContext(ctx, queryStoreOrderDetail, model)
	return errorx.SqlError(err, errorx.SqlTransaction)
}
