package models

import (
	"fmt"
	"gangdou/config"

	"github.com/go-xorm/xorm"
)

var (
	dbClient *xorm.Engine
)

// DBEngine : 数据库驱动
func DBEngine() *xorm.Engine {
	if dbClient != nil {
		return dbClient
	}
	conf := config.InitConf().Database
	driveSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
	engine, err := xorm.NewEngine(conf.Adapter, driveSource)
	if err != nil {
		fmt.Println("database connect failed")
	}
	engine.ShowSQL(true)
	dbClient = engine
	return dbClient
}
