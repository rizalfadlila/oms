package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type CustomerDefinition interface {
	Store(ctx context.Context, model *models.Customers) error
}
