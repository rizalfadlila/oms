package order

import (
	"context"
	"github.com/jatis/oms/models"
)

func (m module) GetReportOrder(ctx context.Context) (*models.ReportOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (m module) GetReportDetailOrderByOrderID(ctx context.Context, id int64) ([]models.ReportOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}
