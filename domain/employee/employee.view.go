package employee

import (
	"context"
)

func (m *module) GetIDByWorkPhone(ctx context.Context, workPhone string) (*int64, error) {
	var (
		id int64
	)

	if err := m.GetQueryerFromContext(ctx).GetContext(ctx, &id, queryGetIDByWorkPhone, workPhone); err != nil {
		return nil, err
	}

	return &id, nil
}
