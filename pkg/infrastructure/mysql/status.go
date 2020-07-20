package mysql

import (
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

func Status() {
	pingErr := Client.DB().Ping()
	if pingErr != nil {
		xray.ErrMini(pingErr)
	}
}
