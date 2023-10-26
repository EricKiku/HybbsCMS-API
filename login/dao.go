package login

import (
	"fmt"
	"hybbscms-api/model"
	"hybbscms-api/tools"
)

func LoginDao(code string, card string, pwd string) model.SysRes {
	var adminer []map[string]interface{}
	tools.DB.Raw("SELECT u_id,u_name,u_code FROM adminer WHERE u_code = ? AND u_card = ? AND u_password = ? ", code, card, pwd).Scan(&adminer)
	fmt.Printf("len(aminer):%v", len(adminer))
	// 如果数组长度为0，则返回失败
	if len(adminer) == 0 {
		return model.SysRes{Status: "500"}
	} else {
		return model.SysRes{Status: "200", Data: adminer}
	}
}

// 获取所有code
func CodesDao() model.SysRes {
	var codes []map[string]interface{}
	tools.DB.Raw("SELECT u_name,u_code,u_permissions FROM adminer").Scan(&codes)
	return model.SysRes{Status: "200", Data: codes}
}
