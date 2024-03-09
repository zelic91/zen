package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

type databaseData struct {
	ModuleName string
	Databases  []c.Database
}

func generateDatabases(
	outputPath string,
	config *c.Config,
) {
	postgresData := databaseData{
		ModuleName: config.ModuleName,
		Databases:  []c.Database{},
	}

	mongoData := databaseData{
		ModuleName: config.ModuleName,
		Databases:  []c.Database{},
	}

	for _, db := range config.Databases {
		if db.Type == "postgres" {
			postgresData.Databases = append(postgresData.Databases, db)
		} else if db.Type == "mongo" {
			mongoData.Databases = append(mongoData.Databases, db)
		}
	}

	generatePostgres(outputPath, config, postgresData)
	generateMongo(outputPath, config, mongoData)
}

func generatePostgres(
	outputPath string,
	config *c.Config,
	data interface{},
) {
	common.GenerateGeneric(
		outputPath+"/db/postgres",
		config.RootTemplatePath+"/db/postgres",
		config,
		data,
	)
}

func generateMongo(
	outputPath string,
	config *c.Config,
	data interface{},
) {
	common.GenerateSpecific(
		outputPath+"/db/mongo/util.go",
		config.RootTemplatePath+"/db/mongo/util.go.tmpl",
		config,
		data,
	)

	// for serviceName, database := range serviceDatabaseMap {
	// 	if database.Type != "mongo" {
	// 		continue
	// 	}
	// 	packageName := strings.ToLower(serviceName)
	// 	config.CurrentPackage = packageName
	// 	service := config.ServiceWithName(serviceName)

	// 	if service == nil {
	// 		continue
	// 	}

	// 	config.CurrentModelName = service.Model
	// 	model := database.ModelWithName(config.CurrentModelName)
	// 	if model == nil {
	// 		continue
	// 	}
	// 	config.CurrentModel = *model
	// 	common.GenerateSpecific(
	// 		outputPath+"/"+strings.ToLower(packageName)+"/model.go",
	// 		config.RootTemplatePath+"/db/mongo/model.go.tmpl",
	// 		config,
	// 	)
	// }
}
