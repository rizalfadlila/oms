package employee

import (
	"context"
	"github.com/jatis/oms/models"
)

func (u *EmployeeUsecase) Store(ctx context.Context, model *models.Employee) error {
	return u.employeeManager.Store(ctx, model)
}

func (u *EmployeeUsecase) GetIDByWorkPhone(ctx context.Context, workPhone string) (*int64, error) {
	return u.employeeManager.GetIDByWorkPhone(ctx, workPhone)
}
