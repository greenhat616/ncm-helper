package protocol

import (
	"encoding/hex"
	"github.com/a632079/ncm-helper/src/protocol/crypto"
	"github.com/a632079/ncm-helper/src/protocol/request"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// SongDetail is a func that impl that get details of songs from netease cloud music
func (p *NCM) SongDetail(songIds []string) (result NeteaseSongDetailResponseData, err error) {
	if !p.isLogin {
		err = notLoginError()
		return
	}
	var buffer []string
	for _, id := range songIds {
		buffer = append(buffer, "{\"id\": "+id+"}")
	}
	data := map[string]interface{}{
		"c":   "[" + strings.Join(buffer, ",") + "]",
		"ids": "[" + strings.Join(songIds, ",") + "]",
	}
	options := request.Options{
		Cookies: p.Cookies,
		Crypto:  "weapi",
	}
	options = appendCustomClientIP(options, p.IP)
	resp, err := request.CreateRequest("POST", "https://music.163.com/weapi/v3/song/detail", data, options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Data, &result)
	if err != nil {
		return
	}
	if result.Code != 200 || len(result.Songs) == 0 {
		log.Debug(resp)
		log.Debug(result)
		err = errors.New("不能获得歌曲信息")
	}
	return
}

// SongURL is a func that get URLs of songs from netease cloud music,
// NOTICE: br is 999000 by default
func (p *NCM) SongURL(songIds []string, br int) (result NeteaseSongURLResponseData, err error) {
	cookies := p.Cookies
	n := true
	for _, cookie := range cookies {
		if cookie.Value == "MUSIC_U" {
			n = false
		}
	}
	if n {
		u := crypto.Util{}
		var k []byte
		k, err = u.GenRandomBytes(16)
		if err != nil {
			return
		}
		cookies = append(cookies, &http.Cookie{
			Name:  "_ntes_nuid",
			Value: hex.EncodeToString(k),
		})
	}
	data := map[string]interface{}{
		"ids": "[" + strings.Join(songIds, ",") + "]",
		"br":  br,
	}
	options := request.Options{
		Cookies: cookies,
		UA:      "pc",
		Crypto:  "linuxapi",
	}
	options = appendCustomClientIP(options, p.IP)
	resp, err := request.CreateRequest("POST", "https://music.163.com/api/song/enhance/player/url", data, options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(resp.Data, &result); err != nil {
		return
	}
	if result.Code != 200 || len(result.Data) == 0 {
		log.Debug(result)
		err = errors.New("不能获得歌曲播放地址")
	}
	return
}
