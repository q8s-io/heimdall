package mysql

import (
	"log"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"github.com/q8s-io/heimdall/pkg/entity/model"
)

var Client *gorm.DB
var connErr interface{}

func Init() {
	mysqlConfig := model.Config.MySQL
	connInfo := mysqlConfig.UserName + ":" + mysqlConfig.PassWord + "@tcp(" + mysqlConfig.Host + ":" + mysqlConfig.Port + ")/" + mysqlConfig.DB + "?charset=utf8&parseTime=True&loc=Local"

	Client, connErr = gorm.Open("mysql", connInfo)
	if connErr != nil {
		log.Println(connErr)
	}
	// 链接池
	Client.DB().SetMaxOpenConns(5)

	go func() {
		taskConnect := time.NewTicker(3 * time.Second)
		for {
			<-taskConnect.C
			go Status()
		}
	}()
}
