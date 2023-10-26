package postData

import "github.com/gin-gonic/gin"

func PostRouter(router *gin.Engine) {
	r := router.Group("/post")
	// 获取所有帖子
	r.GET("/allPost", getAllPostsService)
	// 模糊查询帖子
	r.GET("/getPostLike", getPostLikeQueryService)
	// 获取回复
	r.GET("/replys", getReplyByPIdService)
	// 删除帖子
	r.DELETE("/post", deletePostService)
	// 置顶或取消置顶帖子
	r.PUT("/topOrNoTop", topOrcancelTopPostService)
}
