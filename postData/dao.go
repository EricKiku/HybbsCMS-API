package postData

import (
	"fmt"
	"gorm.io/gorm"
	"hybbscms-api/model"
	"hybbscms-api/tools"
)

// 获取所有帖子

func getAllPostsDao() model.SysRes {
	var posts []map[string]interface{}
	tx := tools.DB.Raw("SELECT * FROM `user`,post,zone WHERE post.`z_id` = zone.`z_id` AND post.`p_lz` = user.`u_id`").Scan(&posts)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: posts}
	}
}

// 帖子模糊查询
func getPostLikeQueryDao(keyWord string) model.SysRes {
	var sql1 = "SELECT * FROM `user`,post,zone WHERE post.`z_id` = zone.`z_id` AND post.`p_lz` = user.`u_id`"
	var sql2 = " AND (post.p_title like '%" + keyWord + "%' OR post.p_content like '%" + keyWord + "%' OR zone.z_name like '%" + keyWord + "%' OR user.u_nick like '%" + keyWord + "%')"
	var posts []map[string]interface{}
	fmt.Println("sql:", sql1+sql2)
	tx := tools.DB.Raw(sql1 + sql2).Scan(&posts)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: posts}
	}
}

// 获取回复
func getReplyByPId(pId int) model.SysRes {
	var replys []map[string]interface{}
	tx := tools.DB.Raw("SELECT * FROM reply,user WHERE reply.u_id = user.u_id AND reply.p_id = ?", pId).Scan(&replys)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: replys}
	}
}

// 删除帖子
func deletePostDao(pId int) model.SysRes {
	//把对应分区的帖子数量-1
	var zId int
	tx := tools.DB.Raw("SELECT z_id FROM post WHERE p_id = ?", pId).Scan(&zId)
	tx.Exec("UPDATE zone SET z_posts = z_posts -1 WHERE z_id = ?", zId)
	//删除帖子
	tx.Exec("DELETE FROM post WHERE p_id = ?", pId)
	//删除帖子的回复
	tx.Exec("DELETE FROM reply WHERE p_id = ?", pId)

	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200"}
	}
}

// 置顶帖子或是取消置顶帖子
func topOrcancelTopPostDao(pId int, opType bool) model.SysRes {
	var tx *gorm.DB
	if opType {
		tx = tools.DB.Exec("UPDATE post SET p_top = 1 WHERE p_id = ?", pId)
	} else {
		tx = tools.DB.Exec("UPDATE post SET p_top = 0 WHERE p_id = ?", pId)
	}
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200"}
	}
}
