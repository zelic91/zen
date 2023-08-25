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
	"github.com/zelic91/zen/config"
	"gopkg.in/yaml.v3"
)

const rootTemplatePath = "templates"

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

func create(configFile string, outputDir string) {
	// config := readConfig(configFile)

	// Make the target path
	if err := makeTargetPath(outputDir); err != nil {
		log.Fatal(err)
	}

	// Generate root files
	generateGeneric(outputDir, fmt.Sprintf("%s/%s", rootTemplatePath, "root"), "")

	// Generate config

	// Generate commands

	// Generate database

	// Generate services
}

func makeTargetPath(outputDir string) error {
	dirName := filepath.Dir(outputDir)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func readConfig(configFile string) *config.Config {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config config.Config
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatal(err)
	}

	return &config
}

func generateGeneric(outputPath string, inputPath string, overridePath string) {
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
		dirs, err := RootFs.ReadDir(fmt.Sprintf("%s%s", inputPath, dirName))
		if err != nil {
			fmt.Printf("Error parsing templates: %v\n", err)
		}

		for _, dir := range dirs {
			newDirName := fmt.Sprintf("%s/%s", dirName, dir.Name())
			if dir.IsDir() {
				stack.Push(newDirName)
			} else {
				templateMap[dir.Name()] = newDirName
				fileList = append(fileList, fmt.Sprintf("%s%s", inputPath, newDirName))
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
		outFileName := fmt.Sprintf("%s/%s", outputPath, templateMap[tmpl.Name()])

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
}
