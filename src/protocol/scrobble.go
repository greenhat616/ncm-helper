package protocol

import (
	"github.com/a632079/ncm-helper/src/protocol/request"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (p *NCM) Scrobble(songId string, sourceId string, time string) (err error) {
	if !p.isLogin {
		err = errors.New("未登录")
		return
	}
	data := map[string]interface{}{
		"download": 0,
		"end":      "playend",
		"id":       songId,
		"sourceId": sourceId,
		"time":     time,
		"type":     "song",
		"wifi":     0,
	}
	options := request.Options{
		Cookies: p.Cookies,
		Crypto:  "weapi",
	}
	if p.IP != "" {
		options.IP = p.IP
	}
	resp, err := request.CreateRequest("POST", "https://music.163.com/weapi/feedback/weblog", data, options)
	if err != nil {
		return
	}
	log.Debug(resp)
	return
}
