package wechat

import (
	"net/url"

	"fmt"

	"github.com/astaxie/beego/httplib"
)

const (
	QR_SCENE       = "QR_SCENE"       // temporary qrcode
	QR_LIMIT_SCENE = "QR_LIMIT_SCENE" // permanent qrcode
)

type QrCodeAction struct {
	ExpireSeconds uint32 `json:"expire_seconds"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneID  uint32 `json:"scene_id"`
			SceneStr string `json:"scene_str"`
		} `json:"scene"`
	} `json:"action_info"`
}

type QrCodeInfo struct {
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
	Ticket        string `json:"ticket"`
	ExpireSeconds uint32 `json:"expire_seconds"`
	Url           string `json:"url"`
}

func getQrCode(action *QrCodeAction) (*QrCodeInfo, error) {
	params := url.Values{}
	params.Set("access_token", AccessToken())
	req, err := httplib.Post(wechatAPI + "/cgi-bin/qrcode/create?" + params.Encode()).JSONBody(action)
	if err != nil {
		return nil, err
	}
	qrInfo := &QrCodeInfo{}
	err = req.ToJSON(qrInfo)
	if err != nil {
		return nil, err
	}
	if qrInfo.ErrCode != 0 {
		return nil, fmt.Errorf("get qrcode return code %d, %s", qrInfo.ErrCode, qrInfo.ErrMsg)
	}
	return qrInfo, nil
}

func GetTemporaryQrCode(expireSecs uint32, sceneID uint32) (*QrCodeInfo, error) {
	action := &QrCodeAction{}
	action.ExpireSeconds = expireSecs
	action.ActionName = QR_SCENE
	action.ActionInfo.Scene.SceneID = sceneID
	return getQrCode(action)
}

func GetPermanentQrCode(expireSecs uint32, sceneID uint32, sceneStr string) (*QrCodeInfo, error) {
	action := &QrCodeAction{}
	action.ExpireSeconds = expireSecs
	action.ActionName = QR_SCENE
	action.ActionInfo.Scene.SceneID = sceneID
	action.ActionInfo.Scene.SceneStr = sceneStr
	return getQrCode(action)
}
