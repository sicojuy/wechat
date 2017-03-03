This is a wechat platform SDK

## Quick Start

###### Download and install

    go get github.com/sicojuy/wechat

###### Sample code
```go
package main

import (
    "log"
	"github.com/sicojuy/wechat"
)

const (
	AppID     = "" // "your wechat app ID"
	AppSecret = "" // "your wechat app secret"
)

func main() {
	err := wechat.Run(AppID, AppSecret)
	if err != nil {
		log.Fatal(err)
	}

	userList, err := wechat.GetUserList()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("get users count: %d", len(userList.Data.OpenID))
}
```
