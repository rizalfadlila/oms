package customer

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *CustomerUsecase) Store(ctx context.Context, model *models.Customers) error {
	err := u.customerManager.Store(ctx, model)
	return err
}
