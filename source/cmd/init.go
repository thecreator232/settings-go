package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "notifire",
	Short: "Notifire for sending notifications",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logrus.WithFields(log.Fields{
			"version": os.Getenv("VERSION"),
		}).Info("Starting up")
	},
}

func init() {
	RootCmd.AddCommand()
}
