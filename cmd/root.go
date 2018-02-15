package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

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

	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Check on the config values being used",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viper.GetString("testVariable"))
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.json", "Default config file is config.yaml")
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
