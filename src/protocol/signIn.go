package protocol

import (
	"errors"
	"github.com/a632079/ncm-helper/src/protocol/request"
	log "github.com/sirupsen/logrus"
)

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
	if p.IP != "" {
		options.IP = p.IP
	}
	resp, err := request.CreateRequest("POST", "https://music.163.com/weapi/point/dailyTask", data, options)
	if err != nil {
		return
	}
	// parse result
	result := make(map[string]interface{})
	err = json.Unmarshal(resp.Data, &result)
	if err != nil {
		return
	}
	c, ok := result["code"]
	if !ok {
		log.Error(resp)
		err = errors.New("code is not exist")
		return
	}
	code, ok := c.(int)
	if !ok {
		log.Error(resp)
		err = errors.New("can't parse code")
		return
	}
	switch code {
	case -2:
		err = errors.New("重复签到")
	case 301:
		err = errors.New("未登录")
	}
	return
}
