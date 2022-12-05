package worker

import (
	"context"
	"fmt"
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/lib/log"
	"github.com/jatis/oms/models"
	"strings"
	"sync"
)

type Job func(ctx context.Context, data []interface{}, usecases *service.Usecases) error

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
		worker:     10,
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
				attemp := 0
				var err = make([]string, 0)
				for attemp <= 5 {
					var (
						errRun error
					)

					func(errRun *error) {
						defer func() {
							if err := recover(); err != nil {
								*errRun = fmt.Errorf("%v", err)
							}
						}()

						*errRun = job(context.Background(), row, svc.UseCases)

					}(&errRun)

					if errRun == nil {
						break
					}

					err = append(err, errRun.Error())
					attemp++
				}

				if len(err) > 0 {
					log.Errorf("worker %s error: %v", w.workerType, strings.Join(err, ", "))
				}

				wg.Done()
			}
		}(workerIndex, w.service, w.rows, w.wg, job)
	}

	log.Infof("worker %s completed execute all job", w.workerType)
}

func customerJob(ctx context.Context, data []interface{}, usecases *service.Usecases) error {
	customer := models.NewCustomerFromRowCSV(data)
	return usecases.CustomerUsecase.Store(ctx, customer)
}

func employeeJob(ctx context.Context, data []interface{}, usecases *service.Usecases) error {
	employee := models.NewEmployeeFromRowCSV(data)
	return usecases.EmployeeUsecase.Store(ctx, employee)
}

func orderJob(ctx context.Context, data []interface{}, usecases *service.Usecases) error {
	order := models.NewOrderFromRowCSV(data)
	return usecases.OrderUsecase.StoreOrder(ctx, order)
}

func orderDetailJob(ctx context.Context, data []interface{}, usecases *service.Usecases) error {
	orderDetail := models.NewOrderDetailFromRowCSV(data)
	return usecases.OrderUsecase.StoreOrderDetail(ctx, orderDetail)
}

func shippingMethodJob(ctx context.Context, data []interface{}, usecases *service.Usecases) error {
	shipping := models.NewShippingMethodFromRowCSV(data)
	return usecases.ShippingMethodUsecase.Store(ctx, shipping)
}

func productJob(ctx context.Context, data []interface{}, usecases *service.Usecases) error {
	product := models.NewProductFromRowCSV(data)
	return usecases.ProductUsecase.Store(ctx, product)
}
