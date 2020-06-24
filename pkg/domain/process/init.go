package process

import (
	"github.com/BurntSushi/toml"
	
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	"github.com/q8s-io/heimdall/pkg/models"
)

func Init(confPath string) {
	//init runtime
	if _, err := toml.DecodeFile(confPath, &models.Config); err != nil {
		ginext.ErrorLogger(err)
		return
	}
	//init log
	ginext.InitLog()
}
