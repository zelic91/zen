package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

func generateDebugConfig(
	outputPath string,
	config *c.Config,
) {
	common.GenerateSpecific(
		outputPath+"/.vscode/launch.json",
		config.RootTemplatePath+"/vscode/launch.json.tmpl",
		config,
		config,
	)
}
