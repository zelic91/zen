/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/zelic91/zen/config"
)

// dummyCmd represents the dummy command
var dummyCmd = &cobra.Command{
	Use:   "dummy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: RunDummy,
}

func init() {
	rootCmd.AddCommand(dummyCmd)
}

func RunDummy(cmd *cobra.Command, args []string) {
	config := config.Config{}

	t := reflect.TypeOf(config)
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		fmt.Printf("%s\n", t.Field(i).Type)
	}
}
