package oauthhdl

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gofly/app/service/ldauth_service/ldauth_dto"
	"gofly/ldauth"
	"net/http"
	"net/url"
	"time"
)

type ldHandler struct {
}

var LdHdl = &ldHandler{}

// 第一步 其他平台请求到Auth， 设置cookie, 跳转到linux.do authorize

// 第二步 用户在 linux.do 授权， 登录， 注册等

// 第三步 授权完成，跳转到Callback， 使用code置换token, 使用token获取userInfo,  读取redirect_to， 补充userInfo信息，跳转。

func (l *ldHandler) Auth(c *gin.Context) {
	form := new(ldauth_dto.AuthForm)
	err := c.ShouldBind(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = form.Check(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数redirect_to无效", "error": err.Error()})
		return
	}

	rediUrl := fmt.Sprintf("/oauth2/setRedirect?redirect_to=%s", form.RedirectTo)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("host127", "host127", int(time.Second*60*30), "/", "", true, true)
	c.SetCookie("redirect_to", form.RedirectTo, int(time.Second*60*30), "/", "", true, true)
	c.Redirect(http.StatusFound, rediUrl)

}

func (l *ldHandler) SetRedirect(c *gin.Context) {
	form := new(ldauth_dto.AuthForm)
	err := c.ShouldBind(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = form.Check(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数redirect_to无效", "error": err.Error()})
		return
	}

	ldAuthUrl := ldauth.AuthCodeUrl()
	c.Redirect(http.StatusFound, ldAuthUrl)
}

func (l *ldHandler) Callback(c *gin.Context) {
	// 参数接收、 验证
	codeState := new(ldauth.CodeState)
	err := c.ShouldBind(codeState)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = codeState.Check(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取用户信息，token获取用户信息
	userInfo, err := ldauth.User(context.TODO(), codeState)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 回调
	redirectTo, err := c.Cookie("redirect_to")
	if err != nil { // 不存在redirect_to， 止于这步
		c.JSON(http.StatusOK, userInfo)
		return
	}

	redirectUrl, _ := url.Parse(redirectTo)

	qVals := redirectUrl.Query()
	qVals.Set("id", cast.ToString(userInfo.Id))
	qVals.Set("username", userInfo.Username)
	qVals.Set("name", userInfo.Name)
	qVals.Set("active", cast.ToString(userInfo.Active))
	qVals.Set("trust_level", cast.ToString(userInfo.TrustLevel))
	qVals.Set("silenced", cast.ToString(userInfo.Silenced))
	redirectUrl.RawQuery = qVals.Encode()

	c.Redirect(http.StatusFound, redirectUrl.String())
}
