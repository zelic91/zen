package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

func generateCommon(
	outputPath string,
	config *c.Config,
) {
	common.GenerateGeneric(
		outputPath+"/common",
		config.RootTemplatePath+"/common",
		config,
		config,
	)
}
