package customer

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *CustomerUsecase) Store(ctx context.Context, model *models.Customers) error {
	return u.customerManager.Store(ctx, model)
}

func (u *CustomerUsecase) GetIDByEmail(ctx context.Context, email string) (*int64, error) {
	return u.customerManager.GetIDByEmail(ctx, email)
}
