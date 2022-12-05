package employee

import (
	"context"
	"github.com/jatis/oms/errorx"
	"github.com/jatis/oms/models"
)

func (m module) Store(ctx context.Context, model *models.Employee) error {
	_, err := m.GetExecerFromContext(ctx).NamedExecContext(ctx, queryStore, model)
	return errorx.SqlError(err, errorx.SqlTransaction)
}
