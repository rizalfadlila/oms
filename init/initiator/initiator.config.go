package initiator

import (
	"github.com/jatis/oms/config"
	libcfg "github.com/jatis/oms/lib/config"
	"github.com/jatis/oms/lib/log"
)

func initConfig(filecfg string) *config.Main {
	cfg := &config.Main{}

	if err := libcfg.ReadModuleConfig(cfg, filecfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	return cfg
}
