package mysql

import (
	"log"
)

func Status() {
	pingErr := Client.Ping()
	if pingErr != nil {
		log.Println(pingErr)
	}
}
