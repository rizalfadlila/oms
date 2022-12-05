package customer

import (
	"context"
)

func (m *module) GetIDByEmail(ctx context.Context, email string) (*int64, error) {
	var (
		id int64
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &id, queryGetIDByEmail, email); err != nil {
		return nil, err
	}

	return &id, nil
}
