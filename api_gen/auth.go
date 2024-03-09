package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

func generateAuthService(
	outputPath string,
	config *c.Config,
) {
	common.GenerateSpecific(
		outputPath+"/auth/service.go",
		config.RootTemplatePath+"/service/auth.service.go.tmpl",
		config,
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
		config,
	)
	common.GenerateSpecific(
		outputPath+"/user/service.go",
		config.RootTemplatePath+"/service/user.service.go.tmpl",
		config,
		config,
	)
}
