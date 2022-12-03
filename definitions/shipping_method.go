package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type ShippingMethodDefinition interface {
	Store(ctx context.Context, model *models.ShippingMethod) error
}
