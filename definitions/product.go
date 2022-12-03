package definitions

import (
	"context"
	"github.com/jatis/oms/models"
)

type ProductDefinition interface {
	Store(ctx context.Context, model *models.Product) error
}
