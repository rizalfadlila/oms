package cmd

import (
	"github.com/jatis/oms/cmd/reader"
	"github.com/jatis/oms/lib/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "oms",
		Short: "Order Management Service",
	}
)

func Execute() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	registerCronJobCommand()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}

func registerCronJobCommand() {
	command := reader.ExecReaderCmd()

	rootCmd.AddCommand(command)

	command.PersistentFlags().StringP("source", "n", "", "Filepath source")
	command.MarkPersistentFlagRequired("source")
}
