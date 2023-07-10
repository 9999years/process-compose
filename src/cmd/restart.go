package cmd

import (
	"github.com/f1bonacc1/process-compose/src/client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart [PROCESS]",
	Short: "Restart a process",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := client.RestartProcesses(pcAddress, port, name)
		if err != nil {
			logFatal(err, "failed to restart process %s", name)
		}
		log.Info().Msgf("Process %s restarted", name)
	},
}

func init() {
	processCmd.AddCommand(restartCmd)
}
