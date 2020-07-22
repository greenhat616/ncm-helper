package protocol

import (
	"github.com/a632079/ncm-helper/src/protocol/request"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

// CreatePlaylist is a func that perform a request to create a NCM playlist
func (p *NCM) CreatePlaylist(name string, privacy int) (result NCMCreatePlaylistResponseData, err error) {
	if !p.isLogin {
		err = notLoginError()
		return
	}
	data := map[string]interface{}{
		"name":    name,
		"privacy": privacy,
	}
	options := request.Options{
		Cookies: p.Cookies,
		UA:      "pc",
		Crypto:  "weapi",
	}
	options = appendCustomClientIP(options, p.IP)
	resp, err := request.CreateRequest("POST", "https://music.163.com/weapi/playlist/create", data, options)
	if err != nil {
		return
	}
	var d NCMCreatePlaylistResponseData
	if err = json.Unmarshal(resp.Data, &d); err != nil {
		return
	}
	if d.Code != 200 {
		log.Error(resp)
		err = errors.New("创建歌单失败")
		return
	}
	return d, nil
}

// AddSongsToPlaylist is a func that perform a request that require adding songs to a specific playlist
func (p *NCM) AddSongsToPlaylist(songIDs []string, playlistID string) (err error) {
	if !p.isLogin {
		err = notLoginError()
		return
	}
	data := map[string]interface{}{
		"op":       "add",
		"pid":      playlistID,
		"trackIds": "[" + strings.Join(songIDs, ",") + "]",
	}
	options := request.Options{
		Cookies: p.Cookies,
		IP:      "",
		Crypto:  "weapi",
	}
	options = appendCustomClientIP(options, p.IP)
	resp, err := request.CreateRequest("POST", "https://music.163.com/weapi/playlist/manipulate/tracks", data, options)
	if err != nil {
		return
	}
	var d map[string]interface{}
	if err := json.Unmarshal(resp.Data, &d); err != nil {
		return
	}
	if code := d["code"].(int); code != 200 {
		log.Error(resp.Data)
		err = errors.New("添加歌曲失败")
	}
	return
}
