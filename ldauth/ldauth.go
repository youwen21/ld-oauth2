package ldauth

import (
	"context"
	"encoding/json"
	"errors"
	"gofly/conf"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type CodeState struct {
	Code  string `json:"code" form:"code"`
	State string `json:"state" form:"state"`
}

func (c *CodeState) Check() error {
	if c.Code == "" {
		return errors.New("code is empty")
	}

	return nil
}

type UserInfo struct {
	Id         int    `json:"id" form:"id"`
	Username   string `json:"username" form:"username"`
	Name       string `json:"name" form:"name"`
	Active     bool   `json:"active" form:"active"`
	TrustLevel int    `json:"trust_level" form:"trust_level"`
	Silenced   bool   `json:"silenced" form:"silenced"`
}

var (
	LinuxDoEndpoint = oauth2.Endpoint{
		AuthURL:   "https://connect.linux.do/oauth2/authorize",
		TokenURL:  "https://connect.linux.do/oauth2/token",
		AuthStyle: oauth2.AuthStyleInHeader,
	}

	UserUrl = "https://connect.linux.do/api/user"
)

func AuthConf() *oauth2.Config {
	authConf := &oauth2.Config{
		ClientID:     conf.GetLDClientId(),
		ClientSecret: conf.GetLDClientSecret(),
		//Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: LinuxDoEndpoint,
	}

	return authConf
}

func AuthCodeUrl() string {
	authConf := AuthConf()
	url := authConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url
}

// Token
func Token(ctx context.Context, code string) (*oauth2.Token, error) {
	authConf := AuthConf()

	httpClient := &http.Client{Timeout: 30 * time.Second} // 30 秒超时
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	return authConf.Exchange(ctx, code)
}

// TokenSource  code置换token, token刷新
func TokenSource(ctx context.Context, code string) (oauth2.TokenSource, error) {
	authConf := AuthConf()

	httpClient := &http.Client{Timeout: 30 * time.Second} // 30 秒超时
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := authConf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	ts := authConf.TokenSource(ctx, tok)
	return ts, nil
}

func User(ctx context.Context, cs *CodeState) (*UserInfo, error) {
	token, err := Token(ctx, cs.Code)
	if err != nil {
		return nil, err
	}

	return GetUserInfo(ctx, token.AccessToken)
}

func GetUserInfo(ctx context.Context, accessToken string) (*UserInfo, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", UserUrl, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	reps, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	userInfo := new(UserInfo)
	err = json.NewDecoder(reps.Body).Decode(userInfo)
	return userInfo, err
}
