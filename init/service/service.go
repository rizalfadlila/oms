package service

import (
	"context"
	"github.com/jatis/oms/definitions"
	"github.com/jatis/oms/lib/database/sql"

	"github.com/jatis/oms/config"
	"github.com/jatis/oms/lib/log"
	"github.com/jatis/oms/repositories"
	"sync"
)

type (
	Basic struct {
		MariaClient *sql.Store
	}

	Domains struct {
		CustomerManager       repositories.CustomerManager
		ProductManager        repositories.ProductManager
		EmployeeManager       repositories.EmployeeManager
		OrderManager          repositories.OrderManager
		ShippingMethodManager repositories.ShippingMethodManager
	}

	Usecases struct {
		CustomerUsecase       definitions.CustomerDefinition
		OrderUsecase          definitions.OrderDefinition
		EmployeeUsecase       definitions.EmployeeDefinition
		ProductUsecase        definitions.ProductDefinition
		ShippingMethodUsecase definitions.ShippingMethodDefinition
	}

	Service struct {
		Config   *config.Main
		UseCases *Usecases
		Basic    *Basic
	}

	StopperFn func(ctx context.Context) error
)

func (s *Service) Close(ctx context.Context) {
	stoppers := []StopperFn{
		func(ctx context.Context) error {
			return s.Basic.MariaClient.GetMaster().Close()
		},
	}

	s.stopper(ctx, stoppers)
}

func (s *Service) stopper(ctx context.Context, resources []StopperFn) {
	wg := sync.WaitGroup{}
	wg.Add(len(resources))

	for i := range resources {
		stopper := resources[i]
		go func() {
			defer wg.Done()
			if err := stopper(ctx); err != nil {
				log.WithError(err).Errorln("failed to stop...")
			}
		}()
	}

	wg.Wait()
}
