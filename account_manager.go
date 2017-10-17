package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	QR_SCENE           = "QR_SCENE"           // temporary qrcode
	QR_LIMIT_SCENE     = "QR_LIMIT_SCENE"     // permanent qrcode id
	QR_LIMIT_STR_SCENE = "QR_LIMIT_STR_SCENE" // permanent qrcode string
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
	url := wechatAPI + "/cgi-bin/qrcode/create?" + params.Encode()
	data, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	qrInfo := &QrCodeInfo{}
	err = json.NewDecoder(res.Body).Decode(qrInfo)
	if err != nil {
		return nil, fmt.Errorf("decode response error: %s", err)
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

func GetPermanentQrCode(sceneStr string) (*QrCodeInfo, error) {
	action := &QrCodeAction{}
	action.ActionName = QR_LIMIT_STR_SCENE
	action.ActionInfo.Scene.SceneStr = sceneStr
	return getQrCode(action)
}
