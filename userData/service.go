package userData

import (
	"github.com/gin-gonic/gin"
	"hybbscms-api/model"
	"net/http"
)

// 获取所有用户
func userService(ctx *gin.Context) {
	usersRes := userDao()
	if usersRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: usersRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}
