package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
)

var (
	auditConfig string
	auditCmd    = &cobra.Command{
		Use:   "audit",
		Short: "The audit command allows one to audit selected marketing tools.",
		Long:  `Not available yet`,
	}

	auditInitCmd = &cobra.Command{
		Use:   "init",
		Short: "The init command allows one to generate the initial configuration to use the tools",
		Long:  "Not available yet",
		Run: func(cmd *cobra.Command, args []string) {
			if tool == "ga" {
				initialConfig := googleanalytics.NewConfig()
				initialConfigJSON, err := json.MarshalIndent(initialConfig, "", "\t")
				if err != nil {
					fmt.Println("Unable to generate configuration")
				}
				fmt.Println(string(initialConfigJSON))
			} else if tool == "gtm" {
				fmt.Println("Not available yet")
			} else {
				fmt.Println("This tool is not yet decided to be supported")
			}

		},
	}
)

func getAuditCmd() {
	auditCmd.AddCommand(auditInitCmd)
	auditInitCmd.Flags().StringVar(&tool, "tool", "ga", "Set the tool to be used for audit. The following would be available for use: ga, gtm")
}
