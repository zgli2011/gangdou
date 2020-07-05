package controller

import (
	"gangdou/dao"

	"github.com/kataras/iris/v12"
)

// User : 用户信息
type User struct {
}

// GetUsers : 查询所有用户信息
func GetUsers(ctx iris.Context) {

	// page := ctx.Params().GetBoolDefault("page", true)
	// pageNum := ctx.Params().GetIntDefault("pageNum", 1)
	// pageSize := ctx.Params().GetIntDefault("pageSize", 20)
	// if pageSize > 100 {
	// 	pageSize = 100
	// }
	id := ctx.Params().GetInt64Default("id", -1)
	response := ResponseBean{}
	response.Status = 200
	response.Msg = "success"
	if id < 0 {
		response.Status = 400
		response.Msg = "param error"
		ctx.JSON(&response)
		return
	}
	response.Details = dao.UserDao.
		ctx.JSON(&response)
	// fmt.Println(page)
	// ctx.Writef("%b, %d, %d", page, pageNum, pageSize)
}
