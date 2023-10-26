package postData

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hybbscms-api/model"
	"net/http"
	"strconv"
)

//获取所有帖子

func getAllPostsService(ctx *gin.Context) {
	postsRes := getAllPostsDao()
	if postsRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: postsRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 模糊查询帖子
func getPostLikeQueryService(ctx *gin.Context) {
	keyWord := ctx.Query("keyWord")
	postsRes := getPostLikeQueryDao(keyWord)
	if postsRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: postsRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取回复
func getReplyByPIdService(ctx *gin.Context) {
	pId, err := strconv.Atoi(ctx.Query("p_id"))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	replysRes := getReplyByPId(pId)
	if replysRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "获取成功", Data: replysRes.Data})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 删除帖子
func deletePostService(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	pId := int(msg["p_id"].(float64))
	res := deletePostDao(pId)
	if res.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "操作成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "操作失败"})
	}
}

// 置顶或取消置顶帖子
func topOrcancelTopPostService(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	pId := int(msg["p_id"].(float64))
	opType := msg["opType"].(bool)
	res := topOrcancelTopPostDao(pId, opType)
	if res.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "操作成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "操作失败"})
	}
}
