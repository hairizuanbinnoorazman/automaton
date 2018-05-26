package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/service"
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

	auditRunAuditCmd = &cobra.Command{
		Use:   "runaudit",
		Short: "The runaudit command runs the actual audit command based on the configuration specified",
		Long:  "Not available yet",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, err := ioutil.ReadFile(cfgFile)
			if err != nil {
				errorFeedback := fmt.Sprintf("Unable to load file. %v", err.Error())
				fmt.Println(errorFeedback)
				return
			}

			credFile, err := ioutil.ReadFile(credFile)
			if err != nil {
				errorFeedback := fmt.Sprintf("Unable to load cred file. %v", err.Error())
				fmt.Println(errorFeedback)
				return
			}

			file, err := os.Create(outputFile)
			if err != nil {
				fmt.Println("Error in creation of new file to store output")
				fmt.Println(err.Error())
			}
			_ = bufio.NewWriter(file)

			if tool == "ga" {
				config := googleanalytics.Config{}
				err = json.Unmarshal(configFile, &config)
				if err != nil {
					fmt.Println(fmt.Sprintf("Error in getting the json config. %v", err.Error()))
					return
				}
				client := googleAnalyticsAuth(credFile)
				auditor := googleanalytics.Auditor{
					AccountID:  config.AccountID,
					PropertyID: config.PropertyID,
					ProfileID:  config.ProfileID,
					StartDate:  config.StartDate,
					EndDate:    config.EndDate,
				}
				service := service.Extractor{Client: client}
				results := auditor.Run(service)
				resultsJSON, err := json.MarshalIndent(results, "", "\t")
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(string(resultsJSON))

				// bufferedFile.Flush()
			} else if tool == "gtm" {
				fmt.Println("Not yet implemented")
			} else {
				fmt.Println("Other tools being considered in the future")
			}
		},
	}
)

func getAuditCmd() {
	auditCmd.AddCommand(auditInitCmd)
	auditInitCmd.Flags().StringVar(&tool, "tool", "ga", "Set the tool to be used for audit. The following would be available for use: ga, gtm")

	auditCmd.AddCommand(auditRunAuditCmd)
	auditRunAuditCmd.Flags().StringVar(&tool, "tool", "ga", "Set the tool to be used for audit. The following would be available for use: ga, gtm")
	auditRunAuditCmd.Flags().StringVar(&cfgFile, "config", "config.json", "Set the config file to be used to run the audit")
	auditRunAuditCmd.Flags().StringVar(&credFile, "cred", "cred.json", "Set the cred file to be used to run the audit")
	auditRunAuditCmd.Flags().StringVar(&outputFile, "file", "output.md", "Output of the markdown file if requested. Only works on markdown option")
	auditRunAuditCmd.Flags().StringVar(&outputType, "output", "json", "Set the type of output. Accepts json or markdown")
}
