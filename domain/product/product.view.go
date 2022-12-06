package product

import (
	"context"
	"github.com/jatis/oms/models"
)

func (m *module) GetByProductName(ctx context.Context, name string) (*models.Product, error) {
	var (
		model models.Product
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &model, queryGetByProductName, name); err != nil {
		return nil, err
	}

	return &model, nil
}
