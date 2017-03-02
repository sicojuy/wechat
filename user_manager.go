package wechat

import (
	"net/url"

	"github.com/astaxie/beego/httplib"
)

type UserInfo struct {
	Subscribe     int    `json:"subscribe"`
	OpenID        string `json:"openid"`
	NickName      string `json:"nickname"`
	Sex           int    `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	HeadImgUrl    string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`
	UnionID       string `json:"unionid"`
	Remark        string `json:"remark"`
	GroupID       int    `json:"groupid"`
	TagIDList     []int  `json:"tagid_list"`
}

type OpenIDList struct {
	OpenID []string `json:"openid"`
}

type UserList struct {
	Total      int         `json:"total"`
	Count      int         `json:"count"`
	Data       *OpenIDList `json:"data"`
	NextOpenID string      `json:"next_openid"`
}

func GetUserInfo(openID string) (*UserInfo, error) {
	params := url.Values{}
	params.Set("access_token", AccessToken())
	params.Set("openid", openID)
	params.Set("lang", "zh_CN")
	userInfo := &UserInfo{}
	err := httplib.Get(wechatAPI + "/cgi-bin/user/info?" + params.Encode()).ToJSON(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func GetUserList() (*UserList, error) {
	params := url.Values{}
	params.Set("access_token", AccessToken())
	userList := &UserList{}
	err := httplib.Get(wechatAPI + "/cgi-bin/user/get?" + params.Encode()).ToJSON(userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
