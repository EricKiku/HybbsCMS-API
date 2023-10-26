package zoneData

import "github.com/gin-gonic/gin"

func ZoneDataRouter(router *gin.Engine) {
	r := router.Group("/zoneData")
	r.GET("/zones", getZonesService)
	// 获取分区的数据
	r.GET("/getMessage", getZoneMessageService)
	// 模糊查询分区
	r.GET("/zonesLike", GetZonesByLikeService)
	// 删除分区
	r.DELETE("/deleteZone", DeleteZoneService)
	// 更新分区
	r.PUT("/updateZone", UpdateZoneService)
	// 获取帖子
	r.GET("/posts", getPostsByZoneService)
	// 获取所有分区的关注数和帖子数
	r.GET("/zonesPostsAndFollows", getZonesPostsAndFollowsService)
	// 获取指定分区的帖子发布频率
	r.GET("/postPublishMessage", getPostPublishMessageOfZoneService)
	// 置顶分区
	r.PUT("/toTop", ToTopZoneService)
	// 取消置顶
	r.PUT("/cancelTop", CancelToTopZoneService)
}
