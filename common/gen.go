package common

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
	c "github.com/zelic91/zen/config"
	"github.com/zelic91/zen/funcs"
)

// This method should be general enough to be reused
// - outputPath: the target path that host the generated files, which is a folder
// - inputPath: the original path of the templates, which is a folder
// - config: the config to be used as the data for rendering templates
func GenerateGeneric(
	outputPath string,
	inputPath string,
	config *c.Config,
) {
	templateMap := map[string]string{}
	fileList := []string{}
	stack := NewStack()

	stack.Push("")

	for {
		if stack.Len() == 0 {
			break
		}

		dirName := stack.Pop().(string)
		dirs, err := config.RootFs.ReadDir(fmt.Sprintf("%s%s", inputPath, dirName))
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

	templates := template.Must(template.New("zen-template").Funcs(sprig.FuncMap()).Funcs(funcs.FuncMap()).ParseFS(
		config.RootFs,
		fileList...,
	))

	for _, tmpl := range templates.Templates() {
		// We are not gonna generate the partials into files
		if strings.HasPrefix(tmpl.Name(), "partial") {
			continue
		}

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
			log.Println(rendered.String())
			log.Fatalf("err formatting Go source %s %v", tmpl.Name(), err)
		}

		filePath.Write(formatted)
	}
}

// This method is used to generate a specific template
// - outputPath: the specific path of the generated file
// - template
func GenerateSpecific(
	outputPath string,
	inputPath string,
	config *c.Config,
) {
	templates := template.Must(template.New("zen-template").Funcs(sprig.FuncMap()).Funcs(funcs.FuncMap()).ParseFS(
		config.RootFs,
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

// TODO: To keep
func MakeTargetPath(outputDir string) error {
	dirName := filepath.Dir(outputDir)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
