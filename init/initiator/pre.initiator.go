package initiator

import (
	"github.com/google/gops/agent"
	"github.com/jatis/oms/lib/log"
)

func initAgent() {
	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatalln(err)
	}
}
