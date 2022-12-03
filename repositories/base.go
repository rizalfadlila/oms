package repositories

import "context"

type TransactionFunc func(ctx context.Context) error

type BaseManager interface {
	GenerateUUID() string
	WithTransaction(ctx context.Context, fn TransactionFunc) error
}
