package product

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *ProductUsecase) Store(ctx context.Context, model *models.Product) error {
	return u.productManager.Store(ctx, model)
}

func (u *ProductUsecase) GetByProductName(ctx context.Context, name string) (*models.Product, error) {
	return u.productManager.GetByProductName(ctx, name)
}
