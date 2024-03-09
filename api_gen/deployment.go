package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

func generateDeployment(outputPath string,
	config *c.Config) {
	common.GenerateGeneric(
		outputPath+"/deployment",
		config.RootTemplatePath+"/deployment",
		config,
		config,
	)
}
