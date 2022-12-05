package repositories

import (
	"context"
	"github.com/jatis/oms/models"
)

type EmployeeManager interface {
	Store(ctx context.Context, model *models.Employee) error
	GetIDByWorkPhone(ctx context.Context, workPhone string) (*int64, error)
}
