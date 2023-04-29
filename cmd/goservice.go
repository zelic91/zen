/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/zelic91/zen/common"
)

var (
	packageName string
	directory   string
	templates   *template.Template
)

// goserviceCmd represents the goservice command
var goserviceCmd = &cobra.Command{
	Use:   "goservice",
	Short: "Generate a boilerplate for your service",
	Long:  `This include an implementation on `,
	Run:   GoServiceCmdExec,
}

func init() {
	createCmd.AddCommand(goserviceCmd)
	goserviceCmd.Flags().StringVarP(&packageName, "package", "p", "github.com/zelic91/app", "Package for the service")
	goserviceCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Target directory")
}

func GoServiceCmdExec(cmd *cobra.Command, args []string) {
	var err error
	templateMap := map[string]string{}
	fileList := []string{}
	stack := common.NewStack()

	stack.Push("")

	for {
		if stack.Len() == 0 {
			break
		}

		dirName := stack.Pop().(string)
		dirs, err := RootFs.ReadDir(fmt.Sprintf("templates/goservice%s", dirName))
		if err != nil {
			fmt.Printf("Error parsing templates: %v\n", err)
		}

		for _, dir := range dirs {
			newDirName := fmt.Sprintf("%s/%s", dirName, dir.Name())
			fmt.Println(newDirName)
			if dir.IsDir() {
				stack.Push(newDirName)
			} else {
				templateMap[dir.Name()] = fmt.Sprintf("templates/goservice%s", newDirName)
				fileList = append(fileList, fmt.Sprintf("templates/goservice%s", newDirName))
			}
		}
	}

	templates, err = template.ParseFS(
		RootFs,
		fileList...,
	)
	if err != nil {
		fmt.Printf("Error parsing templates: %v\n", err)
	}

	for _, tmpl := range templates.Templates() {
		fmt.Println(tmpl.Name())
		outFileName := fmt.Sprintf("%s/%s", directory, templateMap[tmpl.Name()])
		outFileName = strings.TrimSuffix(outFileName, ".tmpl")

		var filePath *os.File
		dirName := filepath.Dir(outFileName)
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Printf("error creating parent dirs %v", err)
			return
		}

		filePath, err = os.Create(outFileName)
		if err != nil {
			fmt.Printf("error creating output file %v", err)
			return
		}

		defer filePath.Close()

		if err := templates.ExecuteTemplate(filePath, tmpl.Name(), nil); err != nil {
			log.Fatal(err)
		}
	}
}
