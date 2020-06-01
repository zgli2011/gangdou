package main

import (
	"fmt"
	"gangdou/routes"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

// Configuration :
type Configuration struct {
	db map[string]interface{} `json:yaml:"DB"`
}

func main() {
	app := iris.New()
	cfg := iris.YAML("./setting.yml")
	db, err := xorm.NewEngine("%s", cfg.db["adapter"], "%s://%s:%s@%s:%s/%s", cfg.db["adapter"], cfg.db["password"], cfg.db["user"], cfg.db["host"], cfg.db["port"], cfg.db["name"])
	if err != nil {
		fmt.Println("数据库连接异常：%s", err)
	}
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})
	routes.App(app)
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	addr := cfg.Other["Addr"].(string)
	app.Listen(addr, iris.WithConfiguration(cfg))
}

// NotFound : 404异常，统一返回
func NotFound(ctx iris.Context) {
	ctx.Writef("source not found")
}
