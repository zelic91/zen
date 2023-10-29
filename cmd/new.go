/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	c "github.com/zelic91/zen/config"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new backend service project",
	Long:  `This command will only create a new directory with a sample zen.yaml file to start with.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		new(name)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("name", "n", "testgen", "Create new backend service project.")
}

func new(name string) {
	config := c.Config{
		ModuleName: name,
	}

	// Make the target path
	if err := makeTargetPath(name); err != nil {
		log.Fatal(err)
	}

	generateSpecific(
		name+"/zen.yaml",
		rootTemplatePath+"/project/zen.yaml.tmpl",
		&config,
	)
}
