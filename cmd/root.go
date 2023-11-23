package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCommand represents the base command when called without any subcommands
var rootCommand = &cobra.Command{
	Use:   "rdsmon-cli",
	Short: "A tool for monitoring RDS server side metrics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initializeConfig)
	rootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rdsmon.yaml)")
}

// Initialise the config, default to $HOME/.rdsmon.yaml if not provided.
func initializeConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Default to $HOME/.rdsmon.yaml
		viper.AddConfigPath(homeDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rdsmon")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using configuration file: ", viper.ConfigFileUsed())
	} else {
		cobra.CheckErr(err)
	}

}
