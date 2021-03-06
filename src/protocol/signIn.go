package protocol

import (
	"github.com/a632079/ncm-helper/src/protocol/request"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// SignIn is a func impl the feature, called “签到” in Chinese
func (p *NCM) SignIn(t int) (err error) {
	if !p.isLogin {
		err = errors.New("未登录")
		return
	}
	data := map[string]interface{}{
		"type": t,
	}
	options := request.Options{
		Cookies: p.Cookies,
		Crypto:  "weapi",
	}
	options = appendCustomClientIP(options, p.IP)
	resp, err := request.CreateRequest("POST", "https://music.163.com/weapi/point/dailyTask", data, options)
	if err != nil {
		return
	}
	// parse result
	var result map[string]interface{}
	err = json.Unmarshal(resp.Data, &result)
	if err != nil {
		return
	}
	code := result["code"].(int)
	switch code {
	case -2:
		err = errors.New("重复签到")
	case 301:
		err = errors.New("未登录")
	default:
		log.Error(result)
		err = errors.New("未知错误")
	}
	return
}
