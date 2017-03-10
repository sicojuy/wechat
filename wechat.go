package wechat

import (
	"fmt"
	"time"
)

const (
	wechatAPI = "https://api.weixin.qq.com"
)

var (
	appID     string
	appSecret string

	running bool
)

func Run(_appID, _appSecret string) error {
	if _appID == "" {
		return fmt.Errorf("app ID is empty")
	}
	if _appSecret == "" {
		return fmt.Errorf("app secret is empty")
	}

	appID = _appID
	appSecret = _appSecret

	if running {
		return fmt.Errorf("wechat already running")
	}
	running = true

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
