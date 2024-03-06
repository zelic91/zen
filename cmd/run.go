/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	apigen "github.com/zelic91/zen/api_gen"
)

var runCmd = &cobra.Command{
	Use:   "run -c [config-file] -t [target-folder]",
	Short: "Create a new service based on YAML config",
	Long:  `Create a new service based on YAML config`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		to, _ := cmd.Flags().GetString("to")

		config := readConfig(configFile)
		config.RootFs = RootFs
		config.RootTemplatePath = "templates"

		apigen.Gen(config, to)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("config", "c", "zen.yaml", "YAML config for zen")
	runCmd.Flags().StringP("to", "t", ".", "Destination for generated files")
}
