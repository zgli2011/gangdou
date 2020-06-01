package routes

import (
	"gangdou/controller"

	"github.com/kataras/iris/v12"
)

// App : 路由
func App(application *iris.Application) {
	apis := application.Party("/api_v1")
	{
		apis.Get("/user", controller.GetUser)
	}
	// interfaces := application.Party("/v1", ){

	// }
}
