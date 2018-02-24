package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	credFile string

	rootCmd = &cobra.Command{
		Use:   "automaton",
		Short: "A tool to help manage marketing software",
		Long: `Automaton is a CLI tool that is meant to manage aspects of
		marketing software`,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Automaton",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Automaton v0.0.1")
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(snapshotCmd)
	getSnapshotCmd()

	rootCmd.AddCommand(guideCmd)
	getGuideCmd()

}

// Execute cli commands
func Execute() {
	rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "config.json" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
