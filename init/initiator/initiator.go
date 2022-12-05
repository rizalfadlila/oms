package initiator

import (
	"github.com/jatis/oms/config"
	"github.com/jatis/oms/init/service"
)

type (
	preInit func()
)

type (
	InitiatorManager interface {
		PreInit() InitiatorManager
		InitService() *service.Service
	}

	initiator struct {
		config  *config.Main
		basic   *service.Basic
		domain  *service.Domains
		usecase *service.Usecases
		preInit preInit
	}
)

func New(filecfg string) InitiatorManager {
	cfg := initConfig(filecfg)
	i := &initiator{
		config: cfg,
		preInit: func() {
			initAgent()
		},
	}
	return i
}

func (i *initiator) PreInit() InitiatorManager {
	i.newBasic()
	i.newDomain()
	i.newUsecase()
	return i
}

func (i *initiator) InitService() *service.Service {
	s := &service.Service{
		Basic:    i.basic,
		Config:   i.config,
		UseCases: i.usecase,
	}
	return s
}
