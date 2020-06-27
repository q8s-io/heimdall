package mysql

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/q8s-io/heimdall/pkg/models"
)

var Client *sqlx.DB
var connErr interface{}

func Init() {
	mysqlConfig := models.Config.MySQL
	connInfo := mysqlConfig.UserName + ":" + mysqlConfig.PassWord + "@tcp(" + mysqlConfig.Host + ":" + mysqlConfig.Port + ")/" + mysqlConfig.DB + "?charset=utf8&parseTime=True&loc=Local"

	Client, connErr = sqlx.Open("mysql", connInfo)
	if connErr != nil {
		log.Println(connErr)
	}

	Client.SetMaxOpenConns(5)

	go func() {
		taskConnect := time.NewTicker(3 * time.Second)
		for {
			<-taskConnect.C
			go Status()
		}
	}()
}
