package shippingmethod

import (
	"context"
)

func (m *module) GetIDByMethod(ctx context.Context, method string) (*int64, error) {
	var (
		id int64
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &id, queryGetIDByMethod, method); err != nil {
		return nil, err
	}

	return &id, nil
}
