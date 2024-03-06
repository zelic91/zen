/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create new backend service project",
	Long:  `This command will only create a new directory with a sample zen.yaml file to start with.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		new(args[0])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func new(name string) {
	config := c.Config{
		ModuleName:       name,
		RootFs:           RootFs,
		RootTemplatePath: "templates",
	}

	// Make the target path
	if err := common.MakeTargetPath(name); err != nil {
		log.Fatal(err)
	}

	common.GenerateSpecific(
		name+"/zen.yaml",
		config.RootTemplatePath+"/project/zen.yaml.tmpl",
		&config,
	)
}
