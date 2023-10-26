package login

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hybbscms-api/model"
	"hybbscms-api/tools"
	"net/http"
	"time"
)

type Info struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 验证Token是否有效
func VerifyToken(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*tools.MyClaims)
	if claims != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "Token有效"})
	} else {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "Token失效"})
	}
}

// 登陆服务
func LoginService(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Res{Status: "201", Msg: "参数接收失败"})
	}
	var info Info
	var loginer map[string]interface{}
	json.Unmarshal(data, &loginer)
	code := loginer["code"].(string)
	card := loginer["card"].(string)
	pwd := loginer["pwd"].(string)
	//fmt.Printf("code:%v,card:%v,pwd:%v", code, card, pwd)
	adminerRes := LoginDao(code, card, pwd)
	if adminerRes.Status == "500" {
		ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "登录失败"})
	} else if adminerRes.Status == "200" {
		//定义一个jwt，后续函数的应用
		wt := tools.NewJWt()
		claims := tools.MyClaims{
			Name:     info.Name,
			Password: info.Password,
			StandardClaims: jwt.StandardClaims{
				NotBefore: int64(time.Now().Unix() - 1000),  // 签名生效时间
				ExpiresAt: int64(time.Now().Unix() + 36000), // 签名过期时间
				Issuer:    "myProject",                      // 签名颁发者
			},
		}
		// 根据token.go中进行token的生成
		token, err := wt.CreatToken(claims)
		if err != nil {
			ctx.JSON(http.StatusOK, model.Res{Status: "500", Msg: "登录失败"})
		}
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "查询成功", Data: adminerRes.Data, Token: token})
	}

}

// 获取所有编号
func CodesService(ctx *gin.Context) {
	codesRes := CodesDao()
	if codesRes.Status != "500" {
		ctx.JSON(http.StatusOK, model.Res{Status: "200", Msg: "获取成功", Data: codesRes.Data})
	}
}
