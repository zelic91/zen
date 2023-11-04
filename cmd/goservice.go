/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/zelic91/zen/common"
)

type goServiceData struct {
	Module string
}

var (
	moduleName string
	directory  string
	templates  *template.Template
	data       goServiceData
)

// goserviceCmd represents the goservice command
var goserviceCmd = &cobra.Command{
	Use:   "goservice",
	Short: "Generate a boilerplate for your service",
	Long:  `This include an implementation on `,
	Run:   GoServiceCmdExec,
}

const templatePath = "templates/goservice"

func init() {
	runCmd.AddCommand(goserviceCmd)
	goserviceCmd.Flags().StringVarP(&moduleName, "module", "m", "github.com/zelic91/app", "Module name of the service")
	goserviceCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Target directory")

	// initConfig()
}

// func initConfig() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(".")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(fmt.Errorf("fatal error config file: %w", err))
// 	}
// }

func GoServiceCmdExec(cmd *cobra.Command, args []string) {
	data = goServiceData{
		Module: moduleName,
	}

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
		dirs, err := RootFs.ReadDir(fmt.Sprintf("%s%s", templatePath, dirName))
		if err != nil {
			fmt.Printf("Error parsing templates: %v\n", err)
		}

		for _, dir := range dirs {
			newDirName := fmt.Sprintf("%s/%s", dirName, dir.Name())
			if dir.IsDir() {
				stack.Push(newDirName)
			} else {
				templateMap[dir.Name()] = newDirName
				fileList = append(fileList, fmt.Sprintf("%s%s", templatePath, newDirName))
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
		outFileName := fmt.Sprintf("%s/%s", directory, templateMap[tmpl.Name()])

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

		if err := templates.ExecuteTemplate(filePath, tmpl.Name(), data); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("DONE!")
}
