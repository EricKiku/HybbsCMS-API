package zoneData

import (
	"hybbscms-api/model"
	"hybbscms-api/tools"
)

// 获取所有分区
func getZonesDao() model.SysRes {
	var zones []map[string]interface{}
	tx := tools.DB.Raw("SELECT * FROM zone,user WHERE zone.z_zonelord = user.u_id").Scan(&zones)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: zones}
	}

}

// 获取某个分区的数据
func getZoneMessageDao(zId int) model.SysRes {
	var allFollows int
	var allPosts int // 全部帖子数
	tx := tools.DB.Raw("SELECT COUNT(*) FROM post").Scan(&allPosts)
	tx.Raw("SELECT SUM(z_follows) FROM zone").Scan(&allFollows)
	type tempRes struct {
		allFollows int
		allPosts   int
	}
	var tempres = tempRes{allFollows: allFollows, allPosts: allPosts}
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: tempres}
	}
}

// 删除分区
func DeleteZoneDao(zId int) model.SysRes {
	//删除分区
	tx := tools.DB.Exec("DELETE FROM zone WHERE z_id = ?", zId)
	// 删除帖子
	tx.Exec("DELETE FROM post WHERE z_id = ?", zId)
	// 删除回复
	tx.Exec("DELETE FROM reply WHERE p_id IN (SELECT p_id FROM post WHERE z_id = ?)", zId)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200"}
	}

}

// 更新分区
func UpdateZoneDao(zId int, zName string, zIntroduce string) model.SysRes {
	tx := tools.DB.Exec("UPDATE zone SET z_name = ?,z_introduce = ? WHERE z_id = ?", zName, zIntroduce, zId)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200"}
	}
}

// 模糊查询分区
func GetZonesByLikeDao(keyWords string) model.SysRes {
	var sql = "SELECT * FROM zone,user WHERE  zone.z_zonelord = user.u_id AND (" + "z_name like \"%" + keyWords + "%\" OR z_introduce like \"%" + keyWords + "%\" OR user.u_nick LIKE \"%" + keyWords + "%\")"
	var zones []map[string]interface{}
	tx := tools.DB.Raw(sql).Scan(&zones)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: zones}
	}
}

// 获取该分区的帖子
func getPostsByZoneDao(zId int) model.SysRes {
	var posts []map[string]interface{}
	tx := tools.DB.Raw("SELECT * FROM post,user WHERE post.p_lz = user.u_id AND z_id = ?", zId).Scan(&posts)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: posts}
	}
}

// 获取所有分区的帖子数和关注数
func getZonesPostsAndFollowsDao() model.SysRes {
	var res map[string]interface{}
	tx := tools.DB.Raw("SELECT SUM(z_posts) as posts,SUM(z_follows) as follows FROM zone").Scan(&res)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: res}
	}
}

// 获取指定分区的帖子发布频率
func getPostPublishMessageOfZoneDao(zId int) model.SysRes {
	var res []map[string]interface{}
	tx := tools.DB.Raw("SELECT z_id,p_date,COUNT(p_date) as count FROM post WHERE z_id = ? GROUP BY SUBSTRING_INDEX(p_date,\" \",1)", zId).Scan(&res)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: res}
	}
}

// 置顶分区
func ToTopZoneDao(zId int) model.SysRes {
	tx := tools.DB.Exec("UPDATE zone SET z_top = 1 WHERE z_id = ?", zId)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200"}
	}
}

// 取消
func CancelToTopZoneDao(zId int) model.SysRes {
	tx := tools.DB.Exec("UPDATE zone SET z_top = 0 WHERE z_id = ?", zId)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200"}
	}
}
