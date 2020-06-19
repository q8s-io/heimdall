package process

import (
	"log"
	
	"github.com/BurntSushi/toml"
	
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	"github.com/q8s-io/heimdall/pkg/models"
)

func Init(confPath string) {
	// init runtime
	if _, err := toml.DecodeFile(confPath, &models.Config); err != nil {
		log.Println(err)
		return
	}
	log.Println(models.Config)
	// init log
	ginext.InitLog()
}
