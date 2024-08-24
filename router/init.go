package router

import (
	"github.com/gin-gonic/gin"
	"gofly/app/handler/oauthhdl"
	"gofly/middleware"
)

func InitRouter(engine *gin.Engine) {
	//gin.SetMode(gin.DebugMode)
	engine.GET("/ping", func(gtx *gin.Context) { gtx.String(200, "pong") })
	engine.OPTIONS("/*options_support", middleware.Cors.GinCors())

	oauth(engine)
}

func oauth(engine *gin.Engine) {
	oauthRouter := engine.Group("/oauth2").Use(middleware.Cors.GinCors())
	oauthRouter.GET("/callback", oauthhdl.LdHdl.Callback)

	oauthRouter.GET("/auth", oauthhdl.LdHdl.Auth)
	oauthRouter.GET("/setRedirect", oauthhdl.LdHdl.SetRedirect)
}
