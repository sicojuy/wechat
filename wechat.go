package wechat

import (
	"fmt"
	"time"
)

const (
	wechatAPI = "https://api.weixin.qq.com"
)

var (
	AppID     string
	AppSecret string

	running bool
)

func Run(appID, appSecret string) error {
	if appID == "" || appSecret == "" {
		return fmt.Errorf("you haven't set wechat app ID or secret")
	}
	if running {
		return fmt.Errorf("wechat already running")
	}
	running = true

	AppID = appID
	AppSecret = appSecret

	err := updateAccessToken()
	if err != nil {
		return err
	}

	sleepSecs := TokenExpireIn() / 2
	go func() {
		time.Sleep(time.Duration(sleepSecs) * time.Second)
		err := updateAccessToken()
		if err != nil {
			if TokenExpireAt() < time.Now().Unix() {
				panic(err)
			}
			sleepSecs = 60
		} else {
			sleepSecs = TokenExpireIn() / 2
		}
	}()

	return nil
}
