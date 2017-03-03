package wechat

import (
	"testing"
)

var (
	testAppID     = "" // your wechat app ID
	testAppSecret = "" // your wechat app secret
)

func TestWechat(t *testing.T) {
	err := Run(testAppID, testAppSecret)
	if err != nil {
		t.Fatal("init error: ", err)
	}

	userList, err := GetUserList()
	if err != nil {
		t.Error("get user list error: ", err)
	}

	userCount := len(userList.Data.OpenID)
	t.Logf("get user count: %d", userCount)

	if userCount > 0 {
		userInfo, err := GetUserInfo(userList.Data.OpenID[0])
		if err != nil {
			t.Error("get user info error: ", err)
		}

		t.Logf("get user info: %+v", *userInfo)
	}

	qrInfo, err := GetTemporaryQrCode(120, 123456)
	if err != nil {
		t.Error("get temporary qrcode error: ", err)
	}

	t.Logf("temporary qrcode info: %+v", qrInfo)
}
