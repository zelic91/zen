package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

type commandData struct {
	ModuleName string
	Command    string
	Models     []c.Model
	Databases  []c.Database
	Api        c.Api
	Crawler    c.Crawler
}

func generateCommands(
	outputPath string,
	config *c.Config,
) {
	data := commandData{
		ModuleName: config.ModuleName,
		Models:     config.Models,
		Databases:  config.Databases,
		Api:        config.Api,
		Crawler:    config.Crawler,
	}

	common.GenerateSpecific(
		outputPath+"/cmd/root.go",
		config.RootTemplatePath+"/cmd/root.go.tmpl",
		config,
		config,
	)

	for name, command := range config.Commands {
		data.Command = name
		switch command.Type {
		case "api":
			common.GenerateSpecific(
				outputPath+"/cmd/"+name+".go",
				config.RootTemplatePath+"/cmd/command.api.go.tmpl",
				config,
				data,
			)
		case "crawler":
			common.GenerateSpecific(
				outputPath+"/cmd/"+name+".go",
				config.RootTemplatePath+"/cmd/command.crawler.go.tmpl",
				config,
				data,
			)
		default:
		}
	}
}
