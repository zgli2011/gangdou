package models

import (
	"fmt"
	"gangdou/config"
	"strings"

	"github.com/go-xorm/xorm"
)

// DBEngine : 数据库驱动
func DBEngine() *xorm.Engine {
	var dataSource strings.Builder
	dataSource.WriteString(config.InitConf().Database.Username)
	dataSource.WriteString(":")
	dataSource.WriteString(config.InitConf().Database.Password)
	dataSource.WriteString("@")
	dataSource.WriteString(config.InitConf().Database.Host)
	dataSource.WriteString(":")
	dataSource.WriteString(config.InitConf().Database.Port)
	dataSource.WriteString("/")
	dataSource.WriteString(config.InitConf().Database.DBName)
	fmt.Println(dataSource.String())
	DBClient, err := xorm.NewEngine(config.InitConf().Database.Adapter, dataSource.String())
	if err != nil {
		fmt.Println("database connect failed")
	}
	DBClient.ShowSQL(true)
	return DBClient
}
