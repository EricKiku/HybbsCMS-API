package home

import "github.com/gin-gonic/gin"

func HomeRouter(router *gin.Engine) {
	r := router.Group("/home")
	// 获取分区数据
	r.GET("/zoneCreateMessageGroupByDate", getZoneCreateMessageService)
	// 获取帖子数据
	r.GET("/getPostBestFive", getPostsOrderByPostsService)
	// 获取用户签到数据
	r.GET("/getUserSignin", getSigninForFiveService)
	// 获取评论最多的分区的近期评论数据
	r.GET("/getReplyMessageForZoneBestReply", getreplyMessageForThreeZoneBestReplyService)
	// 获取聊天频率最高的三个频道
	r.GET("/getRoomBestChat", getRoomBestChatForThreeService)
	// 获取所有用户
	r.GET("/getAdminers", getAdminerService)
	// 获取所有频道的聊天数据
	r.GET("/getRoomsChatMessage", getRoomChatMessageService)
	// 获取聊天用户
	r.GET("/getChatToGroup", getChatToGroupService)
}
