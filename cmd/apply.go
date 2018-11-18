package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	applyCmd = &cobra.Command{
		Use:   "apply",
		Short: "The apply command allows one to set settings onto analytics tools",
		Long:  `Not available yet`,
	}

	applyInitCmd = &cobra.Command{
		Use:   "init",
		Short: "The init command allows one to generate the initial configuration to use the tools",
		Long:  "Not available yet",
	}

	applyRunAuditCmd = &cobra.Command{
		Use:   "run",
		Short: "The run command runs applies the settings to the tool specified here",
		Long:  "Not available yet",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Not yet ready")
		},
	}
)

func getApplyCmd() {
	applyCmd.AddCommand(applyInitCmd)
	applyInitCmd.Flags().StringVar(&tool, "tool", "ga", "Set the tool to be used for audit. The following would be available for use: ga, gtm")

	applyCmd.AddCommand(applyRunAuditCmd)
	applyRunAuditCmd.Flags().StringVar(&tool, "tool", "ga", "Set the tool to be used for audit. The following would be available for use: ga, gtm")
	applyRunAuditCmd.Flags().StringVar(&cfgFile, "config", "config.json", "Set the config file to be used to run the audit")
	applyRunAuditCmd.Flags().StringVar(&credFile, "cred", "cred.json", "Set the cred file to be used to run the audit")
}
