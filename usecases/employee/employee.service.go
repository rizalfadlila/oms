package employee

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *EmployeeUsecase) Store(ctx context.Context, model *models.Employee) error {
	return u.employeeManager.Store(ctx, model)
}
