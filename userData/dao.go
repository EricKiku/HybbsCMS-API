package userData

import (
	"hybbscms-api/model"
	"hybbscms-api/tools"
)

// 获取所有用户
func userDao() model.SysRes {
	var users []map[string]interface{}
	tx := tools.DB.Raw("SELECT * FROM user").Scan(&users)
	if tx.Error != nil {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: users}
	}
}
