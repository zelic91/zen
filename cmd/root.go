/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"embed"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	c "github.com/zelic91/zen/config"
	"gopkg.in/yaml.v2"
)

var (
	RootFs embed.FS
)

const version = "1.1.3"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zen [new|run]",
	Short: "Generate code for your next backend service.",
	Long:  `Generate code for your next backend service.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		color.Set(color.FgGreen)
		print(`
::::::::: :::::::::: ::::    ::: 
     :+:  :+:        :+:+:   :+: 
    +:+   +:+        :+:+:+  +:+ 
   +#+    +#++:++#   +#+ +:+ +#+ 
  +#+     +#+        +#+  +#+#+# 
 #+#      #+#        #+#   #+#+# 
######### ########## ###    #### `)
		color.Set(color.FgHiMagenta)
		println("v" + version)
		color.Unset()
		println("Generate configs and stubs for your next backend service.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Version of the zen tool")
}

func readConfig(configFile string) *c.Config {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config c.Config
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatal(err)
	}

	// Preprocess the databases and models
	config.DatabaseMap = map[string]c.Database{}
	for _, db := range config.Databases {
		config.DatabaseMap[db.Name] = db
	}

	config.ModelMap = map[string]c.Model{}
	for _, m := range config.Models {
		config.ModelMap[m.Name] = m
	}

	// Prepare the databases
	for index, db := range config.Databases {
		db.Models = []c.Model{}
		for _, modelRef := range db.ModelRefs {
			db.Models = append(db.Models, config.ModelMap[modelRef])
		}
		config.Databases[index] = db
	}

	// Prepare the api
	for index, resource := range config.Api.Resources {
		resource.Database = config.DatabaseMap[resource.DatabaseRef]
		resource.Model = config.ModelMap[resource.ModelRef]
		config.Api.Resources[index] = resource

		if resource.Database.Type == "postgres" {
			config.HasPostgres = true
		} else if resource.Database.Type == "mongo" {
			config.HasMongo = true
		}
	}

	return &config
}
