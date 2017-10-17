package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type TemplateMessageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgID   int    `json:"msgid"`
}

func SendTemplateMessage(data []byte) (int, error) {
	params := url.Values{}
	params.Set("access_token", AccessToken())
	url := wechatAPI + "/cgi-bin/message/template/send?" + params.Encode()
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return -1, err
	}
	defer res.Body.Close()
	tmr := &TemplateMessageResponse{}
	err = json.NewDecoder(res.Body).Decode(tmr)
	if err != nil {
		return -1, fmt.Errorf("decode response error: %s", err)
	}
	if tmr.ErrCode != 0 {
		return -1, fmt.Errorf("send template message error, code %d, %s", tmr.ErrCode, tmr.ErrMsg)
	}

	return tmr.MsgID, nil
}
