package libutils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

type jwtUtil struct {
}

var Jwt = &jwtUtil{}

type TicketDto struct {
	Authorization string `json:"authorization" form:"authorization" header:"Authorization"`
	Token         string `json:"token" form:"token" header:"token"`
}

// CheckToken 如果是已登录用户， 从token中提取用户ID
// 提取这个方法是为NoRoute中REST使用
func (j *jwtUtil) CheckToken(c *gin.Context, tokenKey string, tokenSecret string) (jwt.MapClaims, error) {
	token, err := j.GetToken(c, tokenKey)
	if err != nil {
		return nil, err
	}
	token = strings.TrimPrefix(token, "Bearer ")

	return j.parseToken(token, tokenSecret)
}

func (j *jwtUtil) GetToken(c *gin.Context, tokenKey string) (string, error) {
	tokenString := c.GetHeader(tokenKey)
	if tokenString != "" {
		return tokenString, nil
	}

	tokenString, _ = c.Cookie(tokenKey)
	if tokenString != "" {
		return tokenString, nil
	}

	form := &TicketDto{}
	_ = c.ShouldBind(form)
	if form.Token != "" { // 优先token
		return form.Token, nil
	}
	if form.Authorization != "" { // 其次authorization
		return form.Authorization, nil
	}

	return "", errors.New("need token")
}

func (j *jwtUtil) parseToken(tokenString string, tokenSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenSecret), nil
	})

	if nil != err {
		return nil, errors.New("token parse failed: " + err.Error())
	}

	if !token.Valid {
		return nil, errors.New("token invalid:" + tokenString)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims invalid")
	}
	return claims, nil
}

func (j *jwtUtil) GenToken(secret string, claims jwt.MapClaims) (string, error) {
	// 生成token https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
	//jwt.MapClaims{"uid": uid, "exp": time.Now().Unix() + 86400*30}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(secret))
}
