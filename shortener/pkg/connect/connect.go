package connect

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

// client 全局的HTTP客户端
// 不是长连接   设置超时时间
var client = &http.Client{
	Transport: &http.Transport{DisableKeepAlives: true},
	Timeout:   2 * time.Second,
}

// Get 判断url是否能请求通
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect client.Get failed", logx.LogField{Key: "err", Value: err.Error()})
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK //别人给我发一个跳转响应这里也不算过
}
