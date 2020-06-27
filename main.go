package main

import (
	"gangdou/config"
	"gangdou/routes"

	"github.com/kataras/iris/v12"
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
