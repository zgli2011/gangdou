package routes

import (
	"gangdou/controller"

	"github.com/kataras/iris/v12"
)

// APIRoute : 路由注册
func APIRoute(app *iris.Application) {
	apiRoute := app.Party("/api_v1")
	{
		apiRoute.Get("/user/{id:int64}", controller.GetUsers)
	}
}

// InterfaceRoute : 路由注册
func InterfaceRoute(app *iris.Application) {

	interfaceRoute := app.Party("/interface_v1")
	{
		userInterface := interfaceRoute.Party("/user")
		{
			userInterface.Get("/user", func(ctx iris.Context) {
				ctx.WriteString(ctx.Path())
			})
		}
		userGroupInterface := interfaceRoute.Party("/user_group")
		{
			userGroupInterface.Get("/user", func(ctx iris.Context) {
				ctx.WriteString(ctx.Path())
			})
		}
	}
}

// ExternalInterfaceRoute : 外部接口路由注册
func ExternalInterfaceRoute(app *iris.Application) {

	interfaceRoute := app.Party("/external_interface_v1")
	{
		userInterface := interfaceRoute.Party("/user")
		{
			userInterface.Get("/user", func(ctx iris.Context) {
				ctx.WriteString(ctx.Path())
			})
		}
		userGroupInterface := interfaceRoute.Party("/user_group")
		{
			userGroupInterface.Get("/user", func(ctx iris.Context) {
				ctx.WriteString(ctx.Path())
			})
		}
	}
}
