package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/q8s-io/heimdall/pkg/models"
)

func Init() {
	mysqlConfig := models.Config.MySQL
	connInfo := mysqlConfig.UserName + ":" + mysqlConfig.PassWord + "@tcp(" + mysqlConfig.Host + ":" + mysqlConfig.Port + ")/" + mysqlConfig.DB + "?charset=utf8&parseTime=True&loc=Local"

	var pool *sqlx.DB
	pool, err := sqlx.Open("mysql", connInfo)
	if err != nil {
		panic(err)
	}

	pool.SetMaxOpenConns(5)
}
