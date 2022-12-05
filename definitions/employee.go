package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type EmployeeDefinition interface {
	Store(ctx context.Context, model *models.Employee) error
	GetIDByWorkPhone(ctx context.Context, workPhone string) (*int64, error)
}
