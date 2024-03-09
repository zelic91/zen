package apigen

import (
	"strings"

	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

type resourceData struct {
	ModuleName string
	Resource   c.Resource
	Model      c.Model
	Database   c.Database
}

func generateResources(
	outputPath string,
	config *c.Config,
) {
	for _, resource := range config.Api.Resources {
		modelRef := resource.ModelRef
		model := config.ModelMap[modelRef]
		databaseRef := resource.DatabaseRef
		database := config.DatabaseMap[databaseRef]

		packageName := strings.ToLower(resource.ModelRef)

		data := resourceData{
			ModuleName: config.ModuleName,
			Model:      model,
			Database:   database,
			Resource:   resource,
		}

		if database.Type == "postgres" {
			common.GenerateSpecific(
				outputPath+"/"+packageName+"/repo.go",
				config.RootTemplatePath+"/service/repo.postgres.go.tmpl",
				config,
				data,
			)
		}

		common.GenerateSpecific(
			outputPath+"/"+packageName+"/service.go",
			config.RootTemplatePath+"/service/service.go.tmpl",
			config,
			data,
		)
	}
}
