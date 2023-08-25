/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zelic91/zen/config"
	"gopkg.in/yaml.v3"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new service based on YAML config",
	Long:  `Create a new service based on YAML config`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		to, _ := cmd.Flags().GetString("to")

		create(configFile, to)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("config", "c", "zen.yaml", "YAML config for zen")
	createCmd.Flags().StringP("to", "t", "testgen", "Destination for generated files")
}

func create(configFile string, to string) {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config config.Config
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", config)
}
