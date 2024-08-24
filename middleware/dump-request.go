package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
)

// DumpRequestBody httpDump request
func DumpRequestBody(c *gin.Context) {
	res, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println(string(res))
	c.Next()
}

func DumpRequest(c *gin.Context) {
	res, _ := httputil.DumpRequest(c.Request, false)
	fmt.Println(string(res))
	c.Next()
}
