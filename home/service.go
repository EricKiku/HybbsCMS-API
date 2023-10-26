package home

import (
	"github.com/gin-gonic/gin"
	"hybbscms-api/model"
	"net/http"
)

// 获取分区创建数据按照日期分组
func getZoneCreateMessageService(ctx *gin.Context) {
	zonesRes := getZoneCreateMessageDao()
	if zonesRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: zonesRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Data: zonesRes.Data, Msg: "获取失败"})
	}
}

// 获取帖子最多的五个分区
func getPostsOrderByPostsService(ctx *gin.Context) {
	postsRes := getPostsOrderByPostsDao()
	if postsRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: postsRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取最近签到过的五个日期
func getSigninForFiveService(ctx *gin.Context) {
	userRes := getSigninForFiveDao()
	if userRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: userRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取评论最多的三个分区最近的评论热度情况
func getreplyMessageForThreeZoneBestReplyService(ctx *gin.Context) {
	messageRes := getreplyMessageForThreeZoneBestReplyDao()
	if messageRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: messageRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取聊天频率最高的三个频道
func getRoomBestChatForThreeService(ctx *gin.Context) {
	roomsRes := getRoomBestChatForThreeDao()
	if roomsRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: roomsRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取所有用户
func getAdminerService(ctx *gin.Context) {
	adminersRes := getAdminerDao()
	if adminersRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: adminersRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取所有频道的聊天数据
func getRoomChatMessageService(ctx *gin.Context) {
	roomsRes := getRoomChatMessageDao()
	if roomsRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: roomsRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取近期在频道聊天的用户
func getChatToGroupService(ctx *gin.Context) {
	users := getChatToGroupUserDao()
	if users.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: users.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}
