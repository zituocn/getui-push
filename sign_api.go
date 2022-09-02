package getuipush

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zituocn/getui-push/models"
	"time"
)

// getToken 获取过推token
func getToken(appId, appKey, masterSecret string) (token string, err error) {
	sign, timestamp := signature(appKey, masterSecret)
	param := &models.TokenParam{
		Sign:      sign,
		Timestamp: fmt.Sprintf("%d", timestamp),
		AppKey:    appKey,
	}
	bodyByte, err := makeReqBody(param)
	if err != nil {
		return
	}
	b, err := HttpRequest("POST", appId+"/auth", "", bodyByte)
	if err != nil {
		return
	}
	resp := new(models.TokenResp)
	err = json.Unmarshal(b, &resp)
	if resp.Data.Token == "" {
		err = errors.New("返回的token为空")
		return
	}
	token = resp.Data.Token
	return
}

// signature 生成签名方法
func signature(appKey, masterSecret string) (sign string, timestamp int) {
	timestamp = int(time.Now().Unix() * 1000)
	original := fmt.Sprintf("%s%d%s", appKey, timestamp, masterSecret)
	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)
	sign = fmt.Sprintf("%x", sum)
	return
}
