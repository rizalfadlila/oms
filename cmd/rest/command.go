package rest

import (
	"github.com/jatis/oms/init/assembler"
	"github.com/jatis/oms/lib/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:              "serve-rest",
		Short:            "OMS Service",
		PersistentPreRun: rootPreRun,
		RunE:             runREST,
	}
)

func rootPreRun(cmd *cobra.Command, args []string) {
	log.LogInit()
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	log.SetLevel("debug")
}

func runREST(cmd *cobra.Command, args []string) error {
	config, _ := cmd.Flags().GetString("config")

	if config == "" {
		config = "files/config/local.yaml"
	}

	bootstrapREST(assembler.PreRun(config, assembler.AssemblerRest).New())
	return nil
}

func bootstrapREST(assembler assembler.AssemblerManager) {
	assembler.RegisterService().
		RegisterHandler().
		Run()
}

func ExecRestCmd() *cobra.Command {
	return cmd
}
