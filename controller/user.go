package controller

import (
	"github.com/kataras/iris/v12"
)

// User : 用户信息
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
}

// GetUser : 查询所有用户信息
func GetUser(ctx iris.Context) {

	
	// page := ctx.Params().GetBoolDefault("page", true)
	// pageNum := ctx.Params().GetIntDefault("pageNum", 1)
	// pageSize := ctx.Params().GetIntDefault("pageSize", 20)
	// if pageSize > 100 {
	// 	pageSize = 100
	// }
	response := ResponseBean{}
	// response.Status = 200
	response.Msg = "success"
	user := User{}
	user.Username = "lizhaogang"
	user.Age = 30
	user.Sex = "男"
	response.Details = user

	ctx.JSON(&response)
	// fmt.Println(page)
	// ctx.Writef("%b, %d, %d", page, pageNum, pageSize)
}
