package zoneData

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hybbscms-api/model"
	"net/http"
	"strconv"
)

// 获取所有的分区
func getZonesService(ctx *gin.Context) {
	zonesRes := getZonesDao()
	if zonesRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: zonesRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取分区的数据
func getZoneMessageService(ctx *gin.Context) {
	zId, err := strconv.Atoi(ctx.Query("zId"))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	messagesRes := getZoneMessageDao(zId)
	if messagesRes.Status == "200" {
		fmt.Println(messagesRes.Data)
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: messagesRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 模糊查询分区
func GetZonesByLikeService(ctx *gin.Context) {
	key := ctx.Query("key")
	zonesRes := GetZonesByLikeDao(key)
	if zonesRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Data: zonesRes.Data, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 删除分区
func DeleteZoneService(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
		return
	}
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	var zId = int(msg["z_id"].(float64))
	res := DeleteZoneDao(zId)
	if res.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "操作成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "操作失败"})
	}
}

// 更新分区
func UpdateZoneService(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
		return
	}
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	var zId = int(msg["z_id"].(float64))
	zName := msg["z_name"].(string)
	zIntroduce := msg["z_introduce"].(string)
	res := UpdateZoneDao(zId, zName, zIntroduce)
	if res.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "操作成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "操作失败"})
	}
}

// 获取分区帖子
func getPostsByZoneService(ctx *gin.Context) {
	zId, err := strconv.Atoi(ctx.Query("z_id"))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	postsRes := getPostsByZoneDao(zId)
	if postsRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "获取成功", Data: postsRes.Data})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// // 获取所有分区的帖子数和关注数
func getZonesPostsAndFollowsService(ctx *gin.Context) {
	res := getZonesPostsAndFollowsDao()
	if res.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "获取成功", Data: res.Data})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 获取指定分区的帖子发布频率
func getPostPublishMessageOfZoneService(ctx *gin.Context) {
	zId, err := strconv.Atoi(ctx.Query("z_id"))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	resRes := getPostPublishMessageOfZoneDao(zId)
	if resRes.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "获取成功", Data: resRes.Data})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "获取失败"})
	}
}

// 置顶分区
func ToTopZoneService(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	zId := int(msg["z_id"].(float64))
	res := ToTopZoneDao(zId)
	if res.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "操作成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "操作失败"})
	}
}

// 取消置顶分区
func CancelToTopZoneService(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "参数接收失败"})
	}
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	zId := int(msg["z_id"].(float64))
	res := CancelToTopZoneDao(zId)
	if res.Status == "200" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "操作成功"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "操作失败"})
	}
}
