/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"embed"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootFs embed.FS
)

const version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zen [new|run]",
	Short: "Generate code for your next backend service.",
	Long:  `Generate code for your next backend service.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		println(`
::::::::: :::::::::: ::::    ::: 
     :+:  :+:        :+:+:   :+: 
    +:+   +:+        :+:+:+  +:+ 
   +#+    +#++:++#   +#+ +:+ +#+ 
  +#+     +#+        +#+  +#+#+# 
 #+#      #+#        #+#   #+#+# 
######### ########## ###    #### 
		`)
		println("zen " + version)
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
