package middle_auth

import (
	"github.com/gin-gonic/gin"
)

func GetAdminId(c *gin.Context) int {
	return c.GetInt(AdminUserKey)
}

func GetSystemId(c *gin.Context) int {
	return c.GetInt(InnerSystemKey)
}

func GetUserId(c *gin.Context) int {
	return c.GetInt(UserAuthKey)
}
