package mysql

import (
	"log"
)

func InserData(sqlInfo string) error {
	log.Println(sqlInfo)
	_, err := Client.Exec(sqlInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
