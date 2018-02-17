package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hairizuanbinnoorazman/automaton/snapshot"
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

	snapshotCmd = &cobra.Command{
		Use:   "snapshot",
		Short: "Use this command to create a snapshot of your GA account",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			client := googleAnalyticsAuth(credFile)

			config, _ := ioutil.ReadFile(cfgFile)
			type gaConfig struct {
				GaAccountID  string
				GaPropertyID string
				GaViewID     string
			}
			var liveGAConfig gaConfig
			json.Unmarshal(config, &liveGAConfig)

			snapshotData := snapshot.GetSnapshot(client, liveGAConfig.GaAccountID, liveGAConfig.GaPropertyID, liveGAConfig.GaViewID)
			fmt.Println(string(snapshotData))
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(snapshotCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.json", "Default config file is config.yaml")
	rootCmd.PersistentFlags().StringVar(&credFile, "cred", "cred.json", "Default config file is cred.yaml")
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
