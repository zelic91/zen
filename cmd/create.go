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
	c "github.com/zelic91/zen/config"
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

	// Generate common
	generateCommon(outputDir, config)

	// Generate commands
	generateCommands(outputDir, config)

	// Generate API
	generateApi(outputDir, config)

	// Generate databases
	generateDatabases(outputDir, config)

	// Generate mandatory services: auth and user
	generateAuthService(outputDir, config)
	generateUserService(outputDir, config)

	// Generate services
	generateServices(outputDir, config)

	log.Println("🍺🍺🍺 DONE.")
}

func makeTargetPath(outputDir string) error {
	dirName := filepath.Dir(outputDir)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
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

	return &config
}

// Generate the root files such as Makefile or sample env
func generateRootFiles(
	outputPath string,
	config *c.Config,
) {

	config.ServiceDatabaseMap = buildServiceDatabaseMap(config)

	generateGeneric(
		outputPath,
		rootTemplatePath+"/root",
		config,
	)
}

func generateConfig(
	outputPath string,
	config *c.Config,
) {
	generateGeneric(
		outputPath+"/config",
		rootTemplatePath+"/config",
		config,
	)
}

func generateCommon(
	outputPath string,
	config *c.Config,
) {
	generateGeneric(
		outputPath+"/common",
		rootTemplatePath+"/common",
		config,
	)
}

func generateCommands(
	outputPath string,
	config *c.Config,
) {

	config.ServiceDatabaseMap = buildServiceDatabaseMap(config)
	config.ServiceOperationMap = buildServiceOperationMap(config)

	generateSpecific(
		outputPath+"/cmd/root.go",
		rootTemplatePath+"/cmd/root.go.tmpl",
		config,
	)

	for name, command := range config.Commands {
		config.CurrentCommand = name
		switch command.Type {
		case "api":
			generateSpecific(
				outputPath+"/cmd/"+name+".go",
				rootTemplatePath+"/cmd/command.api.go.tmpl",
				config,
			)
		default:
		}
	}
}

func generateDatabases(
	outputPath string,
	config *c.Config,
) {
	for _, db := range config.Databases {
		if db.Type == "postgres" {
			generateGeneric(
				outputPath+"/db/postgres",
				rootTemplatePath+"/db/postgres",
				config,
			)
		} else if db.Type == "mongo" {

			serviceDatabaseMap := buildServiceDatabaseMap(config)

			generateSpecific(
				outputPath+"/db/mongo/util.go",
				rootTemplatePath+"/db/mongo/util.go.tmpl",
				config,
			)

			for serviceName, database := range serviceDatabaseMap {
				if database.Type != "mongo" {
					continue
				}
				packageName := strings.ToLower(serviceName)
				config.CurrentPackage = packageName
				config.CurrentModelName = config.Services[serviceName].Model
				config.CurrentModel = database.Models[config.CurrentModelName]
				generateSpecific(
					outputPath+"/"+strings.ToLower(packageName)+"/model.go",
					rootTemplatePath+"/db/mongo/model.go.tmpl",
					config,
				)
			}

		}
	}
}

func generateApi(
	outputPath string,
	config *c.Config,
) {
	config.ServiceOperationMap = buildServiceOperationMap(config)

	generateGeneric(
		outputPath+"/api",
		rootTemplatePath+"/api",
		config,
	)
}

func generateAuthService(
	outputPath string,
	config *c.Config,
) {
	generateSpecific(
		outputPath+"/auth/service.go",
		rootTemplatePath+"/service/auth.service.go.tmpl",
		config,
	)
}

func generateUserService(
	outputPath string,
	config *c.Config,
) {
	generateSpecific(
		outputPath+"/user/repo.go",
		rootTemplatePath+"/service/user.repo.postgres.go.tmpl",
		config,
	)
	generateSpecific(
		outputPath+"/user/service.go",
		rootTemplatePath+"/service/user.service.go.tmpl",
		config,
	)
}

func generateServices(
	outputPath string,
	config *c.Config,
) {
	serviceDatabaseMap := buildServiceDatabaseMap(config)
	for serviceName, service := range config.Services {
		packageName := strings.ToLower(serviceName)
		config.CurrentPackage = packageName

		database := serviceDatabaseMap[serviceName]

		config.CurrentModelName = service.Model
		config.CurrentModel = database.Models[config.CurrentModelName]

		if database.Type == "postgres" {
			generateSpecific(
				outputPath+"/"+packageName+"/repo.go",
				rootTemplatePath+"/service/repo.postgres.go.tmpl",
				config,
			)
		}

		config.CurrentServiceName = serviceName
		config.CurrentService = service
		generateSpecific(
			outputPath+"/"+packageName+"/service.go",
			rootTemplatePath+"/service/service.go.tmpl",
			config,
		)
	}
}

// This method should be general enough to be reused
// - outputPath: the target path that host the generated files, which is a folder
// - inputPath: the original path of the templates, which is a folder
// - config: the config to be used as the data for rendering templates
func generateGeneric(
	outputPath string,
	inputPath string,
	config *c.Config,
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

// This method is used to generate a specific template
// - outputPath: the specific path of the generated file
// - template
func generateSpecific(
	outputPath string,
	inputPath string,
	config *c.Config,
) {
	templates = template.Must(template.New("zen-template").Funcs(sprig.FuncMap()).Funcs(funcs.FuncMap()).ParseFS(
		RootFs,
		inputPath,
	))

	if len(templates.Templates()) > 1 {
		log.Fatalf("invalid inputPath value: %s\n. Abort", inputPath)
		return
	}

	// There must be only one template that match the inputPath
	tmpl := templates.Templates()[0]

	var filePath *os.File
	dirName := filepath.Dir(outputPath)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		log.Printf("error creating parent dirs %v", err)
		return
	}

	filePath, err = os.Create(outputPath)
	if err != nil {
		log.Printf("error creating output file %s %v", tmpl.Name(), err)
		return
	}

	defer filePath.Close()

	tmpl = tmpl.Funcs(sprig.FuncMap())

	var rendered bytes.Buffer
	if err := templates.ExecuteTemplate(&rendered, tmpl.Name(), config); err != nil {
		log.Fatalf("err executing template %v", err)
	}

	if !strings.Contains(tmpl.Name(), ".go") {
		filePath.Write(rendered.Bytes())
		return
	}

	formatted, err := format.Source(rendered.Bytes())
	if err != nil {
		log.Println(rendered.String())
		log.Fatalf("err formatting Go source %s %v", tmpl.Name(), err)
	}

	filePath.Write(formatted)
}

func buildServiceOperationMap(config *c.Config) map[string][]c.ApiPath {
	ret := map[string][]c.ApiPath{}

	for _, path := range config.Api.Paths {
		for _, apiPath := range path {
			serviceName := apiPath.Service
			l := ret[serviceName]
			if l == nil {
				l = []c.ApiPath{apiPath}
			} else {
				l = append(l, apiPath)
			}
			ret[serviceName] = l
		}
	}

	return ret
}

func buildServiceDatabaseMap(config *c.Config) map[string]c.Database {
	ret := map[string]c.Database{}
	for serviceName, service := range config.Services {
		ret[serviceName] = config.Databases[service.Database]
	}
	return ret
}
