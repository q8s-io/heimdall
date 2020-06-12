package process

import (
	"log"
	
	"github.com/BurntSushi/toml"
	
	ge "github.com/70data/heimdall/pkg/infrastructure/gin-extender"
	"github.com/70data/heimdall/pkg/models"
)

func Init(confPath string) {
	// init runtime
	if _, err := toml.DecodeFile(confPath, &models.Config); err != nil {
		log.Println(err)
		return
	}
	log.Println(models.Config)
	// init log
	ge.InitLog()
}
