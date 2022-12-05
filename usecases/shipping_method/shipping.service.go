package shippingmethod

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *ShippingMethodUsecase) Store(ctx context.Context, model *models.ShippingMethod) error {
	return u.shippingMethodManager.Store(ctx, model)
}

func (u *ShippingMethodUsecase) GetIDByMethod(ctx context.Context, method string) (*int64, error) {
	return u.shippingMethodManager.GetIDByMethod(ctx, method)
}
