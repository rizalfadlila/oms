package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type CustomerDefinition interface {
	Store(ctx context.Context, model *models.Customers) error
	GetIDByEmail(ctx context.Context, email string) (*int64, error)
}
