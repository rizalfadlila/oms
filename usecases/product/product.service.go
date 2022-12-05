package product

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *ProductUsecase) Store(ctx context.Context, model *models.Product) error {
	return u.productManager.Store(ctx, model)
}
