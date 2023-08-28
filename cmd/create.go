/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/spf13/cobra"
	"github.com/zelic91/zen/common"
	"github.com/zelic91/zen/config"
	"github.com/zelic91/zen/funcs"
	"gopkg.in/yaml.v3"
)

// Root template path points to the directory which hosts
// every templates.
// The templates will be organized into subfolders
// by their functionalities. For example:
// - root: all necessary files for the project roots: Makefile, sample env, etc.
const rootTemplatePath = "templates"

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
	config := readConfig(configFile)

	// Make the target path
	if err := makeTargetPath(outputDir); err != nil {
		log.Fatal(err)
	}

	// Generate root files
	generateRootFiles(outputDir, config)

	// Generate config
	generateConfig(outputDir, config)

	// Generate commands

	// Generate databases
	generateDatabases(outputDir, config)

	// Generate services

	log.Println("DONE.")
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

// Generate the root files such as Makefile or sample env
func generateRootFiles(
	outputPath string,
	config *config.Config,
) {
	generateGeneric(
		outputPath,
		rootTemplatePath+"/root",
		config,
	)
}

func generateConfig(
	outputPath string,
	config *config.Config,
) {
	generateGeneric(
		outputPath+"/config",
		rootTemplatePath+"/config",
		config,
	)
}

func generateDatabases(
	outputPath string,
	config *config.Config,
) {
	for _, db := range config.Databases {
		if db.IsPostgres() {
			generateGeneric(
				outputPath+"/db/postgres",
				rootTemplatePath+"/db/postgres",
				config,
			)
		} else {
			generateGeneric(
				outputPath+"/db/mongo",
				rootTemplatePath+"/db/mongo",
				config,
			)
		}
	}
}

// This method should be general enough to be reused
// - outputPath: the target path that host the generated files
// - inputPath: the original path of the templates
// - config: the config to be used as the data for rendering templates
func generateGeneric(
	outputPath string,
	inputPath string,
	config *config.Config,
) {
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
			log.Printf("Error parsing templates: %v\n", err)
		}

		for _, dir := range dirs {
			newDirName := fmt.Sprintf("%s/%s", dirName, dir.Name())
			if dir.IsDir() {
				stack.Push(newDirName)
			} else {
				templateMap[dir.Name()] = newDirName
				filePath := inputPath + newDirName
				fileList = append(fileList, filePath)
			}
		}
	}

	templates = template.Must(template.New("zen-template").Funcs(sprig.FuncMap()).Funcs(funcs.FuncMap()).ParseFS(
		RootFs,
		fileList...,
	))

	for _, tmpl := range templates.Templates() {
		strippedOutFileName := strings.TrimSuffix(templateMap[tmpl.Name()], ".tmpl")
		outFileName := fmt.Sprintf("%s/%s", outputPath, strippedOutFileName)

		var filePath *os.File
		dirName := filepath.Dir(outFileName)
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			log.Printf("error creating parent dirs %v", err)
			return
		}

		filePath, err = os.Create(outFileName)
		if err != nil {
			log.Printf("error creating output file %s %v", tmpl.Name(), err)
			return
		}

		defer filePath.Close()

		tmpl = tmpl.Funcs(sprig.FuncMap())

		var rendered bytes.Buffer
		if err := templates.ExecuteTemplate(&rendered, tmpl.Name(), config); err != nil {
			log.Fatal(err)
		}

		if !strings.Contains(tmpl.Name(), ".go") {
			filePath.Write(rendered.Bytes())
			continue
		}

		formatted, err := format.Source(rendered.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		filePath.Write(formatted)
	}
}
