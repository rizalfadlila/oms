package repositories

import (
	"context"
	"github.com/jatis/oms/models"
)

type CustomerManager interface {
	Store(ctx context.Context, model *models.Customers) error
}
