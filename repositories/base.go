package repositories

import "context"

type TransactionFunc func(ctx context.Context) error

type BaseManager interface {
	GenerateUUID() int64
	WithTransaction(ctx context.Context, fn TransactionFunc) error
}
