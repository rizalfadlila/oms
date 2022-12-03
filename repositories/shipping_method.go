package repositories

import (
	"context"
	"github.com/jatis/oms/models"
)

type ShippingMethodManager interface {
	Store(ctx context.Context, model *models.ShippingMethod) error
}
