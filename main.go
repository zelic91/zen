/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"

	"github.com/zelic91/zen/cmd"
)

var (
	//go:embed templates
	rootFs embed.FS
)

func main() {
	cmd.RootFs = rootFs
	cmd.Execute()
}
