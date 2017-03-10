package wechat

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func VerifySignature(sign, appToken, timestamp, nonce string) error {
	signItems := []string{appToken, timestamp, nonce}
	sort.Sort(sort.StringSlice(signItems))
	signStr := strings.Join(signItems, "")

	mySign := fmt.Sprintf("%x", sha1.Sum([]byte(signStr)))
	if mySign != sign {
		return fmt.Errorf("signature verify failed")
	}
	return nil
}
