package shippingmethod

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *ShippingMethodUsecase) Store(ctx context.Context, model *models.ShippingMethod) error {
	return u.shippingMethodManager.Store(ctx, model)
}
