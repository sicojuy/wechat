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
    AppID    = "you wechat app ID"
    AppSecret    = "you wechat app secret"
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

    log.Printf("user open ID list: %+v", userList.Data.OpenID)
}
```
