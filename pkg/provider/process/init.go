package process

import (
	"log"

	"github.com/BurntSushi/toml"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/qlog"
)

func Init(confPath string) {
	// init runtime
	if _, err := toml.DecodeFile(confPath, &model.Config); err != nil {
		log.Println(err)
		return
	}
	// init log
	qlog.InitLog()
}
