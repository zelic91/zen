package crawlergen

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
}
