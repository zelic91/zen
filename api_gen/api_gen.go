package apigen

import (
	"log"
	"strings"

	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

func Gen(config *c.Config, outputDir string) {

	// Make the target path
	if err := common.MakeTargetPath(outputDir); err != nil {
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

	// TODO: Move these into the API
	// Generate mandatory services: auth and user
	generateUserService(outputDir, config)
	generateAuthService(outputDir, config)

	// Generate other services
	generateServices(outputDir, config)

	// Generate config for debugging with VSCode
	generateDebugConfig(outputDir, config)

	// Generate k8s deployment manifests
	generateDeployment(outputDir, config)

	// Generate GitHub workflows
	generateGitHubWorkflow(outputDir, config)

	log.Println("🍺🍺🍺 DONE.")
}

// Generate the root files such as Makefile or sample env
func generateRootFiles(
	outputPath string,
	config *c.Config,
) {

	config.ServiceDatabaseMap = buildServiceDatabaseMap(config)

	common.GenerateGeneric(
		outputPath,
		config.RootTemplatePath+"/root",
		config,
	)
}

func generateConfig(
	outputPath string,
	config *c.Config,
) {
	common.GenerateGeneric(
		outputPath+"/config",
		config.RootTemplatePath+"/config",
		config,
	)
}

func generateCommon(
	outputPath string,
	config *c.Config,
) {
	common.GenerateGeneric(
		outputPath+"/common",
		config.RootTemplatePath+"/common",
		config,
	)
}

func generateCommands(
	outputPath string,
	config *c.Config,
) {

	config.ServiceDatabaseMap = buildServiceDatabaseMap(config)
	config.ServiceOperationMap = buildServiceOperationMap(config)

	common.GenerateSpecific(
		outputPath+"/cmd/root.go",
		config.RootTemplatePath+"/cmd/root.go.tmpl",
		config,
	)

	for name, command := range config.Commands {
		config.CurrentCommand = name
		switch command.Type {
		case "api":
			common.GenerateSpecific(
				outputPath+"/cmd/"+name+".go",
				config.RootTemplatePath+"/cmd/command.api.go.tmpl",
				config,
			)
		case "crawler":
			common.GenerateSpecific(
				outputPath+"/cmd/"+name+".go",
				config.RootTemplatePath+"/cmd/command.crawler.go.tmpl",
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
			common.GenerateGeneric(
				outputPath+"/db/postgres",
				config.RootTemplatePath+"/db/postgres",
				config,
			)
		} else if db.Type == "mongo" {

			serviceDatabaseMap := buildServiceDatabaseMap(config)

			common.GenerateSpecific(
				outputPath+"/db/mongo/util.go",
				config.RootTemplatePath+"/db/mongo/util.go.tmpl",
				config,
			)

			for serviceName, database := range serviceDatabaseMap {
				if database.Type != "mongo" {
					continue
				}
				packageName := strings.ToLower(serviceName)
				config.CurrentPackage = packageName
				service := config.ServiceWithName(serviceName)

				if service == nil {
					continue
				}

				config.CurrentModelName = service.Model
				model := database.ModelWithName(config.CurrentModelName)
				if model == nil {
					continue
				}
				config.CurrentModel = *model
				common.GenerateSpecific(
					outputPath+"/"+strings.ToLower(packageName)+"/model.go",
					config.RootTemplatePath+"/db/mongo/model.go.tmpl",
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

	common.GenerateGeneric(
		outputPath+"/api",
		config.RootTemplatePath+"/api",
		config,
	)
}

func generateAuthService(
	outputPath string,
	config *c.Config,
) {
	common.GenerateSpecific(
		outputPath+"/auth/service.go",
		config.RootTemplatePath+"/service/auth.service.go.tmpl",
		config,
	)
}

func generateUserService(
	outputPath string,
	config *c.Config,
) {
	common.GenerateSpecific(
		outputPath+"/user/repo.go",
		config.RootTemplatePath+"/service/user.repo.postgres.go.tmpl",
		config,
	)
	common.GenerateSpecific(
		outputPath+"/user/service.go",
		config.RootTemplatePath+"/service/user.service.go.tmpl",
		config,
	)
}

func generateServices(
	outputPath string,
	config *c.Config,
) {
	serviceDatabaseMap := buildServiceDatabaseMap(config)
	config.ServiceCrawlerTargetMap = buildServiceCrawlerTargetMap(config)
	for _, service := range config.Services {
		packageName := strings.ToLower(service.Name)
		config.CurrentPackage = packageName

		database := serviceDatabaseMap[service.Name]

		config.CurrentModelName = service.Model
		model := database.ModelWithName(config.CurrentModelName)
		if model == nil {
			continue
		}
		config.CurrentModel = *model

		if database.Type == "postgres" {
			common.GenerateSpecific(
				outputPath+"/"+packageName+"/repo.go",
				config.RootTemplatePath+"/service/repo.postgres.go.tmpl",
				config,
			)
		}

		config.CurrentServiceName = service.Name
		config.CurrentService = service
		common.GenerateSpecific(
			outputPath+"/"+packageName+"/service.go",
			config.RootTemplatePath+"/service/service.go.tmpl",
			config,
		)
	}
}

func generateDebugConfig(
	outputPath string,
	config *c.Config,
) {
	common.GenerateSpecific(
		outputPath+"/.vscode/launch.json",
		config.RootTemplatePath+"/vscode/launch.json.tmpl",
		config,
	)
}

func generateDeployment(outputPath string,
	config *c.Config) {
	common.GenerateGeneric(
		outputPath+"/deployment",
		config.RootTemplatePath+"/deployment",
		config,
	)
}

func generateGitHubWorkflow(
	outputPath string,
	config *c.Config,
) {
	common.GenerateGeneric(
		outputPath+"/.github/workflows",
		config.RootTemplatePath+"/github",
		config,
	)
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

func buildServiceCrawlerTargetMap(config *c.Config) map[string][]c.CrawlerTarget {
	ret := map[string][]c.CrawlerTarget{}

	for _, target := range config.Crawler.Targets {
		serviceName := target.Service
		l := ret[serviceName]
		if l == nil {
			l = []c.CrawlerTarget{target}
		} else {
			l = append(l, target)
		}
		ret[serviceName] = l
	}

	return ret
}

func buildServiceDatabaseMap(config *c.Config) map[string]c.Database {
	ret := map[string]c.Database{}
	for _, service := range config.Services {
		ret[service.Name] = config.Databases[service.Database]
	}
	return ret
}
