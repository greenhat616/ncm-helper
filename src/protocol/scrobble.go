package protocol

import (
	"github.com/a632079/ncm-helper/src/protocol/request"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Scrobble is a func that impl the feature, called “听歌打卡” in Chinese
func (p *NCM) Scrobble(songID string, sourceID string, time string) (err error) {
	if !p.isLogin {
		err = errors.New("未登录")
		return
	}
	data := map[string]interface{}{
		"download": 0,
		"end":      "playend",
		"id":       songID,
		"sourceId": sourceID,
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
