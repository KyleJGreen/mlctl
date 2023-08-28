package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var targetDir string

func init() {
	RootCmd.AddCommand(templateCmd)
	templateCmd.Flags().StringVarP(&targetDir, "target-dir", "t", "output_directory", "Directory to save templated files")
}

var templateCmd = &cobra.Command{
	Use:   "template [argo|airflow]",
	Short: "Template workflow files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateType := args[0]
		configFile := "./template_config.json" // Update the filename
		templateDir := fmt.Sprintf("workflow_templates/%s", templateType)

		config, err := LoadConfiguration(configFile)
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			os.Exit(1)
		}

		// Create the target directory if it doesn't exist
		if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
			fmt.Println("Error creating target directory:", err)
			os.Exit(1)
		}

		if err := TemplateFiles(templateDir, targetDir, config); err != nil {
			fmt.Println("Error templating files:", err)
			os.Exit(1)
		}

		fmt.Println("Files templated successfully!")
	},
}
