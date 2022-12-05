package assembler

import (
	"github.com/jatis/oms/handler/reader"
	"github.com/jatis/oms/init/initiator"
	"github.com/jatis/oms/lib/log"
	"strings"
)

type readerAssembler struct {
	*assembler
	handler    reader.ReaderHandler
	workerType string
	fileConfig string
	fileSource string
}

func (a *readerAssembler) New() AssemblerManager {
	a.assembler = &assembler{
		Initiator: initiator.New(a.fileConfig),
	}

	a.setWorkerType()

	return a
}

func (a *readerAssembler) RegisterService() AssemblerManager {
	a.service = a.Initiator.PreInit().InitService()
	return a
}

func (a *readerAssembler) RegisterHandler() AssemblerManager {
	a.handler = reader.New(a.workerType, &reader.Opts{
		Service: a.service,
	})
	return a
}

func (a *readerAssembler) Run() {
	a.handler.Run(a.fileSource)

	a.WaitAndExit()
}

func (a *readerAssembler) WaitAndExit() {
	select {
	case s := <-a.Signals():
		log.Errorf("Signal %v..", s)
	case err := <-a.handler.Error():
		if err != nil {
			log.Errorf("Listening error reader. %v", err)
		}
	}

	a.GracefullyExit()
}

func (a *readerAssembler) setWorkerType() {
	folders := strings.Split(a.fileSource, "/")

	if len(folders) == 0 {
		return
	}

	filename := strings.Split(folders[len(folders)-1], "_")

	if len(filename) == 0 {
		return
	}

	a.workerType = filename[0]
}
