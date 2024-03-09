package apigen

import (
	"log"

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
	generateResources(outputDir, config)

	// Generate config for debugging with VSCode
	generateDebugConfig(outputDir, config)

	// Generate k8s deployment manifests
	generateDeployment(outputDir, config)

	// Generate GitHub workflows
	generateGitHubWorkflow(outputDir, config)

	log.Println("🍺🍺🍺 DONE.")
}
