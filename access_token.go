package wechat

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/astaxie/beego/httplib"
)

var (
	accessToken   string
	tokenExpireIn int64
	tokenExpireAt int64
	tokenLock     sync.RWMutex
)

type TokenResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpireIn    int64  `json:"expires_in"`
}

func updateAccessToken() error {
	params := url.Values{}
	params.Set("grant_type", "client_credential")
	params.Set("appid", AppID)
	params.Set("secret", AppSecret)

	url := wechatAPI + "/cgi-bin/token?" + params.Encode()
	token := &TokenResponse{}
	err := httplib.Get(url).ToJSON(token)
	if err != nil {
		return fmt.Errorf("get access token error: %s", err)
	}

	if token.ErrCode != 0 {
		return fmt.Errorf("update access token error: %s", token.ErrMsg)
	}

	if token.AccessToken == "" || token.ExpireIn <= 0 {
		return fmt.Errorf("get invalid access token")
	}

	expireAt := time.Now().Unix() + token.ExpireIn

	tokenLock.Lock()
	accessToken = token.AccessToken
	tokenExpireIn = token.ExpireIn
	tokenExpireAt = expireAt
	tokenLock.Unlock()

	return nil
}

func AccessToken() string {
	var token string

	tokenLock.RLock()
	token = accessToken
	tokenLock.RUnlock()

	return token
}

func TokenExpireAt() int64 {
	var expireAt int64

	tokenLock.RLock()
	expireAt = tokenExpireAt
	tokenLock.RUnlock()

	return expireAt
}

func TokenExpireIn() int64 {
	var expireIn int64

	tokenLock.RLock()
	expireIn = tokenExpireIn
	tokenLock.RUnlock()

	return expireIn
}
