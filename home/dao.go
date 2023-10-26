package home

import (
	"hybbscms-api/model"
	"hybbscms-api/tools"
)

func getZoneCreateMessageDao() model.SysRes {
	var zones []map[string]interface{}
	tx := tools.DB.Raw("SELECT z_createDate,COUNT(*) AS count FROM zone GROUP BY SUBSTRING_INDEX(z_createDate,\" \",1) ORDER BY z_createDate DESC LIMIT 5").Scan(&zones)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: zones}
	}
}

func getPostsOrderByPostsDao() model.SysRes {
	var posts []map[string]interface{}
	tx := tools.DB.Raw("SELECT z_name,z_posts FROM zone ORDER BY z_posts DESC LIMIT 5").Scan(&posts)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: posts}
	}
}

// 获取最近五个日期的签到情况
func getSigninForFiveDao() model.SysRes {
	var userSigninMessage []map[string]interface{}
	tx := tools.DB.Raw("SELECT u_signin_date,COUNT(*) AS `count` FROM `user` GROUP BY SUBSTRING_INDEX(u_signin_date,\" \",1) ORDER BY u_signin_date DESC LIMIT 5").Scan(&userSigninMessage)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: userSigninMessage}
	}
}

// 获取拥有评论最多的三个分区最近的评论热度情况
func getreplyMessageForThreeZoneBestReplyDao() model.SysRes {
	var message []map[string]interface{}
	tx := tools.DB.Raw("SELECT reply.r_date,zone.z_name,COUNT(*) AS `count` FROM reply,post,zone WHERE zone.z_id IN (SELECT z_id FROM (SELECT zone.z_id FROM reply,post,zone WHERE reply.p_id=post.p_id AND post.z_id = zone.z_id GROUP BY zone.z_id ORDER BY COUNT(zone.z_id) DESC LIMIT 3) AS t ) AND reply.p_id=post.p_id AND post.z_id=zone.z_id GROUP BY SUBSTRING_INDEX(reply.r_date,\" \",1) ORDER BY SUBSTRING_INDEX(reply.r_date,\" \",1) DESC ").Scan(&message)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: message}
	}

}

// 获取聊天频率最高的三个频道
func getRoomBestChatForThreeDao() model.SysRes {
	var rooms []map[string]interface{}
	tx := tools.DB.Raw("SELECT room.`r_name`,COUNT(*) AS count FROM chattogroup,room WHERE chattogroup.`ctg_group`=room.`r_id` GROUP BY ctg_group ").Scan(&rooms)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: rooms}
	}
}

// 获取所有用户
func getAdminerDao() model.SysRes {
	var adminers []map[string]interface{}
	tx := tools.DB.Raw("SELECT * FROM adminer").Scan(&adminers)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: adminers}
	}
}

// 获取频道所有聊天数据
func getRoomChatMessageDao() model.SysRes {
	var rooms []map[string]interface{}
	tx := tools.DB.Raw("SELECT room.r_name,COUNT(*) AS `count` FROM chattogroup,room WHERE chattogroup.`ctg_group`= room.`r_id` GROUP BY chattogroup.`ctg_group`").Scan(&rooms)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: rooms}
	}
}

// 获取近期在频道聊天的人
func getChatToGroupUserDao() model.SysRes {
	var users []map[string]interface{}
	tx := tools.DB.Raw("SELECT ctg_id, user.u_nick,ctg_date FROM chattogroup,USER WHERE chattogroup.u_id = user.u_id ORDER BY ctg_id DESC LIMIT 30").Scan(&users)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: users}
	}
}
