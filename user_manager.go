package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	url := wechatAPI + "/cgi-bin/user/info?" + params.Encode()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	userInfo := &UserInfo{}
	err = json.NewDecoder(res.Body).Decode(userInfo)
	if err != nil {
		return nil, fmt.Errorf("decode response error: %s", err)
	}
	return userInfo, nil
}

func GetUserList() (*UserList, error) {
	params := url.Values{}
	params.Set("access_token", AccessToken())
	url := wechatAPI + "/cgi-bin/user/get?" + params.Encode()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	userList := &UserList{}
	err = json.NewDecoder(res.Body).Decode(userList)
	if err != nil {
		return nil, fmt.Errorf("decode response error: %s", err)
	}
	return userList, nil
}
