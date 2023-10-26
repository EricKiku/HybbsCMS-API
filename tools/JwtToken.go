package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hybbscms-api/model"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从头部信息的Authorization中的到tocken
		tokenString := c.GetHeader("Authorization")
		//如果信息为空
		if tokenString == "" {
			c.JSON(http.StatusOK, model.Res{Status: "500", Msg: "Authorization header not provided"})
			c.Abort()
			return
		}
		//将tokenString解析成想要的*MyClaims格式，自定义的函数解析，下面内容有讲解
		jwt := NewJWt()
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, model.Res{Status: "500", Msg: err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

type JWT struct {
	// 声明签名信息
	SigningKey []byte
}

// NewJWt 初始化jwt对象,固定返回一个为123456的密钥，（基于你的项目实际需要
func NewJWt() *JWT {
	return &JWT{
		[]byte("123456"),
	}
}

// MyClaims 定义自己的负载
type MyClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwt.StandardClaims
}

// CreatToken 进行构造一个token,参数为接受到的负载
func (j *JWT) CreatToken(claims MyClaims) (string, error) {
	// 指定编码的算法为jwt.SigningMethodHS256,返回一个token的指针
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 进行解析一个token
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	//错误处理 ，使更健壮性、易懂性
	if err != nil {
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// ValidationErrorMalformed是一个uint常量，表示token格式错误
				return nil, fmt.Errorf("token格式错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				//ValidationErrorExpired表示token过期
				return nil, fmt.Errorf("token过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				//ValidationErrorNotValidYet 表示token生效时间未到
				return nil, fmt.Errorf("token生效时间未到")
			} else {
				return nil, fmt.Errorf("token 不可用")
			}
		}
	}
	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token无效")
}
