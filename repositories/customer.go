package repositories

import (
	"context"
	"github.com/jatis/oms/models"
)

type CustomerManager interface {
	Store(ctx context.Context, model *models.Customers) error
	GetIDByEmail(ctx context.Context, email string) (*int64, error)
}
