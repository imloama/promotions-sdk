package sdk

import (
	"github.com/imloama/promotions-sdk/jd"
)

type Config struct {
	//服务器地址
	ServerURL string
	// appid or appkey
	AppKey string
	// appsecret
	AppSecret string
}

// NewJDClient 创建京东接口
func NewJDClient(config Config) *jd.JDClient {
	return &jd.JDClient{}
}
