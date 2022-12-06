package assembler

import (
	"context"
	"github.com/jatis/oms/init/initiator"
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/lib/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	assembler struct {
		Initiator initiator.InitiatorManager

		service *service.Service
	}

	AssemblerManager interface {
		New() AssemblerManager
		RegisterService() AssemblerManager
		RegisterHandler() AssemblerManager
		Run()
	}

	AssemblerType string
)

var (
	AssemblerReader AssemblerType = "reader"
	AssemblerRest   AssemblerType = "rest"
)

func PreRun(path string, typ AssemblerType, args ...string) AssemblerManager {
	switch AssemblerType(typ) {
	case AssemblerReader:
		return &readerAssembler{
			fileConfig: path,
			fileSource: args[0],
		}
	case AssemblerRest:
		return &restAssembler{
			pathConfig: path,
		}
	default:
		log.Fatalln("assembler type not supported")
		return nil
	}
}

func (a *assembler) Signals() <-chan os.Signal {
	done := make(chan os.Signal, 2)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	signal.Notify(done, os.Interrupt, syscall.SIGINT)
	return done
}

func (a *assembler) GracefullyExit() {
	timeout := a.service.Config.Server.Rest.GracefulTimeout
	timeDuration, err := time.ParseDuration(timeout)
	if err != nil {
		timeDuration, _ = time.ParseDuration("15s")
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeDuration)
	defer cancel()

	done := make(chan bool)

	go func() {
		a.service.Close(ctx)
		done <- true
	}()

	select {
	case <-ctx.Done():
		log.Println("timeout waiting all processes to stop")
	case <-done:
		log.Println("all processes stopped successfully")
	}
}
