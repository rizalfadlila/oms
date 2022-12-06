package assembler

import (
	"github.com/jatis/oms/handler/rest"
	"github.com/jatis/oms/init/initiator"
	"github.com/jatis/oms/lib/log"
)

type restAssembler struct {
	*assembler
	handler    rest.RestHandler
	pathConfig string
}

func (a *restAssembler) New() AssemblerManager {
	return &restAssembler{
		assembler: &assembler{
			Initiator: initiator.New(a.pathConfig),
		},
	}
}

func (a *restAssembler) RegisterService() AssemblerManager {
	a.service = a.Initiator.PreInit().InitService()
	return a
}

func (a *restAssembler) RegisterHandler() AssemblerManager {
	a.handler = rest.New(&rest.Opts{
		Service: a.service,
	})
	return a
}

func (a *restAssembler) Run() {
	go a.handler.Run()
	a.WaitAndExit()
}

func (a *restAssembler) WaitAndExit() {
	select {
	case s := <-a.Signals():
		log.Errorf("Signal %v..", s)
	case err := <-a.handler.Error():
		log.Errorf("Listening error rest server. %v", err)
	}

	a.GracefullyExit()
}
