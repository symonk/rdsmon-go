package cmd

import "github.com/spf13/cobra"

// Monitor an RDS instance for the configured metrics.
var monitorCommand = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor RDS metrics",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCommand.AddCommand(monitorCommand)
	monitorCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rdsmon.yaml)")
}
