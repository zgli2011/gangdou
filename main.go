package main

import (
	"gangdou/config"
	"gangdou/routes"
	"go/controller"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func main() {
	config.InitConf()
	app := iris.New()
	app.Logger().SetLevel("info")
	routes.APIRoute(app)
	routes.InterfaceRoute(app)
	routes.ExternalInterfaceRoute(app)
	app.OnErrorCode(iris.StatusNotFound, routes.ResourceNotFound)
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
}

// ResourceNotFound : 404异常
func ResourceNotFound(ctx iris.Context) {
	response := controller.ResponseBean{}
	response.Msg = "Resource not found"
	// ctx.JSON(&response)
	// 直接返回数据
	ctx.JSON(iris.Map{"msg": "Resource not found"})
	ctx.StatusCode(404)
}

func init() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)
	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.InfoLevel)

}
