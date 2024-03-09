package apigen

import (
	"github.com/zelic91/zen/common"
	c "github.com/zelic91/zen/config"
)

type apiData struct {
	ModuleName string
	Api        c.Api
}

func generateApi(
	outputPath string,
	config *c.Config,
) {
	data := apiData{
		ModuleName: config.ModuleName,
		Api:        config.Api,
	}

	common.GenerateGeneric(
		outputPath+"/api",
		config.RootTemplatePath+"/api",
		config,
		data,
	)
}
