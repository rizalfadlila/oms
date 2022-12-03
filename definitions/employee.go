package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type EmployeeDefinition interface {
	Store(ctx context.Context, model *models.Employee) error
}
