package worker

import (
	"context"
	"fmt"
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/lib/log"
	"github.com/jatis/oms/models"
	"sync"
)

type Job func(ctx context.Context, data interface{}, usecases *service.Usecases)

var (
	jobList = map[string]Job{
		"customer":        customerJob,
		"employee":        employeeJob,
		"product":         productJob,
		"shipping_method": shippingMethodJob,
		"order":           orderJob,
		"order_detail":    orderDetailJob,
	}
)

type Worker struct {
	workerType string
	service    *service.Service
	ctx        context.Context
	worker     int
	rows       chan []interface{}
	wg         *sync.WaitGroup
	error      chan error
}

type Opts struct {
	Rows       chan []interface{}
	Wg         *sync.WaitGroup
	Service    *service.Service
	WorkerType string
	Error      chan error
}

func New(o *Opts) *Worker {
	return &Worker{
		workerType: o.WorkerType,
		wg:         o.Wg,
		service:    o.Service,
		rows:       o.Rows,
		error:      o.Error,
	}
}

func (w *Worker) Register() {
	job, ok := jobList[w.workerType]
	if !ok {
		w.error <- fmt.Errorf("invalid worker type :%s", w.workerType)
	}

	for workerIndex := 0; workerIndex <= w.worker; workerIndex++ {
		go func(workerIndex int, svc *service.Service, rows <-chan []interface{}, wg *sync.WaitGroup, job Job) {
			for row := range rows {
				job(context.Background(), row, svc.UseCases)
				wg.Done()
			}
		}(workerIndex, w.service, w.rows, w.wg, job)
	}
}

func (w *Worker) dispatchWorkers() {

}

func customerJob(ctx context.Context, data interface{}, usecases *service.Usecases) {
	customer := models.NewCustomerFromRowCSV(data)
	if err := usecases.CustomerUsecase.Store(ctx, customer); err != nil {
		log.Errorf("error store customer. %v", err)
	}
}

func employeeJob(ctx context.Context, data interface{}, usecases *service.Usecases) {
	employee := models.NewEmployeeFromRowCSV(data)
	if err := usecases.EmployeeUsecase.Store(ctx, employee); err != nil {
		log.Errorf("error store employee. %v", err)
	}
}

func orderJob(ctx context.Context, data interface{}, usecases *service.Usecases) {
	order := models.NewOrderFromRowCSV(data)
	if err := usecases.OrderUsecase.StoreOrder(ctx, order); err != nil {
		log.Errorf("error store order. %v", err)
	}
}

func orderDetailJob(ctx context.Context, data interface{}, usecases *service.Usecases) {
	orderDetail := models.NewOrderDetailFromRowCSV(data)
	if err := usecases.OrderUsecase.StoreOrderDetail(ctx, orderDetail); err != nil {
		log.Errorf("error store order detail. %v", err)
	}
}

func shippingMethodJob(ctx context.Context, data interface{}, usecases *service.Usecases) {
	shipping := models.NewShippingMethodFromRowCSV(data)
	if err := usecases.ShippingMethodUsecase.Store(ctx, shipping); err != nil {
		log.Errorf("error store shipping method. %v", err)
	}
}

func productJob(ctx context.Context, data interface{}, usecases *service.Usecases) {
	prdocut := models.NewProductFromRowCSV(data)
	if err := usecases.ProductUsecase.Store(ctx, prdocut); err != nil {
		log.Errorf("error store product. %v", err)
	}
}
