package employee

import (
	"context"
	"github.com/jatis/oms/models"
)

func (m *module) Store(ctx context.Context, model *models.Employee) error {
	model.ID = m.GenerateUUID()
	_, err := m.GetExecerFromContext(ctx).NamedExecContext(ctx, queryStore, model)
	return err
}
