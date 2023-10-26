package login

import (
	"github.com/gin-gonic/gin"
	"hybbscms-api/tools"
)

func LoginRouter(router *gin.Engine) {
	r := router.Group("/lor")
	// 登录
	r.POST("/login", LoginService)
	// 验证Token
	r.GET("/verifyToken", tools.JWTAuth(), VerifyToken)
	// 获取所有编号
	r.GET("/codes", CodesService)
}
