package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/hairizuanbinnoorazman/automaton/snapshot"
)

var (
	snapshotCmd = &cobra.Command{
		Use:   "snapshot",
		Short: "Use this command to create a snapshot of your GA account",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			cred, err := ioutil.ReadFile(credFile)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			client := googleAnalyticsAuth(cred)

			config, err := ioutil.ReadFile(cfgFile)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
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

func getSnapshotCmd() {
	snapshotCmd.Flags().StringVar(&cfgFile, "config", "config.json", "Default config file is config.yaml")
	snapshotCmd.Flags().StringVar(&credFile, "cred", "cred.json", "Default config file is cred.yaml")
}
