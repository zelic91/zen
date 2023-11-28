/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stubCmd represents the stub command
var stubCmd = &cobra.Command{
	Use:   "stub",
	Short: "Generate the stubs from a specific templates",
	Long:  `Generate the stubs from a specific templates`,
	Run:   RunStub,
}

func init() {
	rootCmd.AddCommand(stubCmd)
}

func RunStub(cmd *cobra.Command, args []string) {
	fmt.Printf("stub called: %s \n", args[0])
}
