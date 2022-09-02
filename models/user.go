package models

// TokenParam 获取token的参数
type TokenParam struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	AppKey    string `json:"appkey"`
}

// TokenResp token返回值
type TokenResp struct {
	Response
	Data struct {
		ExpireTime string `json:"expire_time"`
		Token      string `json:"token"`
	} `json:"data"`
}
