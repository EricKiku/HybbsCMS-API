package main

import (
	"github.com/gin-gonic/gin"
	"hybbscms-api/home"
	"hybbscms-api/login"
	"hybbscms-api/postData"
	"hybbscms-api/userData"
	"hybbscms-api/zoneData"
)

func main() {
	router := gin.Default()
	login.LoginRouter(router)
	home.HomeRouter(router)
	zoneData.ZoneDataRouter(router)
	postData.PostRouter(router)
	userData.UserRouter(router)
	router.Run(":9876")
}
