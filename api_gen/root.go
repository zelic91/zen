package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

type rootData struct {
	ModuleName string
	Commands   map[string]c.Command
	Databases  []c.Database
	Env        map[string]c.Env
	Api        c.Api
}

// Generate the root files such as Makefile or sample env
func generateRootFiles(
	outputPath string,
	config *c.Config,
) {

	data := rootData{
		ModuleName: config.ModuleName,
		Commands:   config.Commands,
		Databases:  []c.Database{},
		Env:        config.Env,
		Api:        config.Api,
	}

	for _, db := range config.Databases {
		if db.Type == "mongo" {
			data.Databases = append(data.Databases, db)
		}
	}

	common.GenerateGeneric(
		outputPath,
		config.RootTemplatePath+"/root",
		config,
		data,
	)
}
