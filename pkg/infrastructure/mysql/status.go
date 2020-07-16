package mysql

import (
	"log"
)

func Status() {
	pingErr := Client.DB().Ping()
	if pingErr != nil {
		log.Println(pingErr)
	}
}
