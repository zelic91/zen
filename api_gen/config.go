package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

func generateConfig(
	outputPath string,
	config *c.Config,
) {
	common.GenerateGeneric(
		outputPath+"/config",
		config.RootTemplatePath+"/config",
		config,
		config,
	)
}
