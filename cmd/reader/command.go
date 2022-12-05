package reader

import (
	"github.com/jatis/oms/init/assembler"
	"github.com/jatis/oms/lib/log"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:              "reader",
		Short:            "File reader",
		Long:             "Read data from csv file",
		PersistentPreRun: rootPreRun,
		RunE:             runReader,
	}
)

func rootPreRun(cmd *cobra.Command, args []string) {
	log.LogInit()
}

func runReader(cmd *cobra.Command, args []string) error {
	config, _ := cmd.Flags().GetString("config")

	if config == "" {
		config = "files/config"
	}

	source := cmd.PersistentFlags().Lookup("source").Value.String()
	a := assembler.PreRun(config, assembler.AssemblerReader, source).New()

	bootstrapReader(a)

	return nil
}

func bootstrapReader(assembler assembler.AssemblerManager) {
	assembler.RegisterService().
		RegisterHandler().
		Run()
}

func ExecReaderCmd() *cobra.Command {
	return cmd
}
