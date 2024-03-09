package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

func generateGitHubWorkflow(
	outputPath string,
	config *c.Config,
) {
	common.GenerateGeneric(
		outputPath+"/.github/workflows",
		config.RootTemplatePath+"/github",
		config,
		config,
	)
}
