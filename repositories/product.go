package repositories

import (
	"context"
	"github.com/jatis/oms/models"
)

type ProductManager interface {
	Store(ctx context.Context, model *models.Product) error
	GetByProductName(ctx context.Context, name string) (*models.Product, error)
}
