package userData

import "github.com/gin-gonic/gin"

func UserRouter(router *gin.Engine) {
	r := router.Group("/user")

	r.GET("/user", userService)
}
