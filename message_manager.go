package wechat

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego/httplib"
)

type TemplateMessageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgID   int    `json:"msgid"`
}

func SendTemplateMessage(tempMsg interface{}) (int, error) {
	params := url.Values{}
	params.Set("access_token", AccessToken())
	req, err := httplib.Post(wechatAPI + "/cgi-bin/message/template/send?" + params.Encode()).JSONBody(tempMsg)
	if err != nil {
		return -1, err
	}
	res := &TemplateMessageResponse{}
	err = req.ToJSON(res)
	if err != nil {
		return -1, err
	}
	if res.ErrCode != 0 {
		return -1, fmt.Errorf("send template message error, code %d, %s", res.ErrCode, res.ErrMsg)
	}

	return res.MsgID, nil
}
