package controller

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

// GetUser : 查询所有用户信息
func GetUser(ctx iris.Context) {
	page := ctx.Params().GetBoolDefault("page", true)
	pageNum := ctx.Params().GetIntDefault("pageNum", 1)
	pageSize := ctx.Params().GetIntDefault("pageSize", 20)
	if pageSize > 100 {
		pageSize = 100
	}
	fmt.Println(page)
	ctx.Writef("%b, %d, %d", page, pageNum, pageSize)
}
