package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type ShippingMethodDefinition interface {
	Store(ctx context.Context, model *models.ShippingMethod) error
	GetIDByMethod(ctx context.Context, method string) (*int64, error)
}
