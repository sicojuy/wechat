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
	MsgID   int64  `json:"msgid"`
}

func SendTemplateMessage(data []byte) (msgID int64, errCode int, errMsg string) {
	params := url.Values{}
	params.Set("access_token", AccessToken())
	url := wechatAPI + "/cgi-bin/message/template/send?" + params.Encode()
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return 0, -1, err.Error()
	}
	defer res.Body.Close()
	tmr := &TemplateMessageResponse{}
	err = json.NewDecoder(res.Body).Decode(tmr)
	if err != nil {
		return 0, -1, fmt.Sprintf("decode response error: %s", err)
	}
	return tmr.MsgID, tmr.ErrCode, tmr.ErrMsg
}
