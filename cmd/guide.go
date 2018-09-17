package cmd

import (
	"fmt"

	"github.com/hairizuanbinnoorazman/automaton/guide"
	"github.com/spf13/cobra"
)

var (
	guideConfig string
	guideCmd    = &cobra.Command{
		Use:   "guide",
		Short: "Use this set of commands to assist in creating guides on usage of marketing tools",
		Long:  `Not available yet`,
	}

	guideInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Use this command to initialize the required json files for the guide you intend to create",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			guideConfig := guide.NewGuideConfig()
			fmt.Println(string(guideConfig))
		},
	}

	guideGenerateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Use this command to generate the output based on configurations and lists",
		Long:  "The generate command allows the creation of an Implementation Guide. The creation of the guide is dependent on the usage of markdown file and csv files.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating a GTM implementation Guide")
			guide.GenerateGuide(guideConfig)
			fmt.Println("File generated. Check output file that was declared in your config file.")
		},
	}
)

func getGuideCmd() {
	guideCmd.AddCommand(guideInitCmd)
	guideCmd.AddCommand(guideGenerateCmd)
	guideGenerateCmd.Flags().StringVar(&guideConfig, "config", "config.json", "Default config file is config.json")
}
