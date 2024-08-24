package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gofly/lib/libutils"
	"gofly/middleware/middle_auth"
	"net/http"
)

func jwtTokenWare(tokenKey string, secret string, storeKey string) func(c *gin.Context) {
	// storeKey gin 和 jwt claims中的key
	// tokenKey header头 或者 cookie 包含jwt串的key
	// secret jwt解密密钥

	return func(c *gin.Context) {
		claims, err := libutils.Jwt.CheckToken(c, tokenKey, secret)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": ""})
			c.Abort()
		}

		value := claims[storeKey]
		c.Set(storeKey, cast.ToInt(value))
		c.Next()
	}
}

// InnerToken 内容服务接口认证
//func InnerToken(c *gin.Context) {
//	err := CheckToken(c, InnerJwtKey, InnerJwtSecret, InnerSystemId)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": ""})
//		c.Abort()
//	}
//	c.Next()
//}

// InnerToken 内部api token
func InnerToken() func(c *gin.Context) {
	tokenKey := middle_auth.InnerJwtKey  // read from config
	secret := middle_auth.InnerJwtSecret //  os.Getenv("USER_JWT_SECRET")
	systemKey := middle_auth.InnerSystemKey
	return jwtTokenWare(tokenKey, secret, systemKey)
}

// AdminToken 避免每次调用AdminTokenWare中间件都要传入 tokenKet, secret, userKey, 三个参数， 封装一下， 自行改动。
func AdminToken() func(c *gin.Context) {
	return jwtTokenWare(middle_auth.AdminJwtKey, middle_auth.AdminJwtSecret, middle_auth.AdminUserKey)
}

// UserToken member用户token
func UserToken() func(c *gin.Context) {
	tokenKey := middle_auth.UserJwtKey // read from config
	secret := middle_auth.UserJwtSecret
	userKey := middle_auth.UserAuthKey
	return jwtTokenWare(tokenKey, secret, userKey)
}
