package middleware

import "github.com/gin-gonic/gin"

//BrowserCacheMiddleware 设轩api接口浏览器缓存
func BrowserCacheMiddleware(c *gin.Context) {
	c.Header("Cache-Control", "max-age=3600")
	// $response->withHeader("Cache-Control", 'max-age=' . $this->expireTime)->withHeader('Expires', gmdate("D, d M Y H:i:s", time() + $this->expireTime) . " GMT");
	c.Next()
}
