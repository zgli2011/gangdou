package routes

import (
	"gangdou/controller"

	"github.com/kataras/iris/v12"
)

// ResourceNotFound : 404异常，统一返回
func ResourceNotFound(ctx iris.Context) {
	response := controller.ResponseBean{}
	response.Msg = "Resource not found"
	// ctx.JSON(&response)
	// 直接返回数据
	ctx.JSON(iris.Map{"msg": "Resource not found"})
	ctx.StatusCode(404)
}
