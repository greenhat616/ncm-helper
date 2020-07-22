package protocol

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/a632079/ncm-helper/src/protocol/request"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
)

// Login is a func that impl email, phone and cookies login
func (p *NCM) Login() (err error) {
	if p.Phone != "" && p.Pass != "" {
		if p.CountryCode == "" {
			p.CountryCode = "86"
		}
		err = p.phoneLogin(p.Phone, p.CountryCode, p.Pass, p.IsMD5Pass)
	} else if p.Email != "" && p.Pass != "" {
		err = p.emailLogin(p.Email, p.Pass, p.IsMD5Pass)
	} else if p.Cookies != nil {
		err = p.CheckLogin()
	} else {
		err = errors.New("mismatch login type")
	}
	return
}

func (p *NCM) phoneLogin(phone string, countyCode string, password string, isMD5Password bool) (err error) {
	data := map[string]interface{}{
		"phone":         phone,
		"countrycode":   countyCode,
		"rememberLogin": "true",
	}
	if !isMD5Password {
		h := md5.New()
		h.Write([]byte(password))
		password = hex.EncodeToString(h.Sum(nil))
	}
	data["password"] = password
	options := request.Options{
		Cookies: nil,
		UA:      "pc",
		Crypto:  "weapi",
	}
	if p.IP != "" {
		options.IP = p.IP
	}
	resp, err := request.CreateRequest(
		"POST",
		"https://music.163.com/weapi/login/cellphone",
		data,
		options)
	if err != nil {
		return
	}
	// request success
	p.Cookies = resp.Cookies
	p.isLogin = true
	return
}

func (p *NCM) emailLogin(email string, password string, isMD5Password bool) (err error) {
	cookies := p.Cookies
	cookies = append(cookies, &http.Cookie{
		Name:  "os",
		Value: "pc",
	})
	data := map[string]interface{}{
		"username":      email,
		"rememberLogin": "true",
	}
	if !isMD5Password {
		h := md5.New()
		h.Write([]byte(password))
		password = hex.EncodeToString(h.Sum(nil))
	}
	data["password"] = password
	options := request.Options{
		Cookies: cookies,
		UA:      "pc",
		Crypto:  "weapi",
	}
	if p.IP != "" {
		options.IP = p.IP
	}
	resp, err := request.CreateRequest("POST", "https://music.163.com/weapi/login", data, options)
	if err != nil {
		if resp.StatusCode == 502 { // password or username err
			err = errors.New("用户名或密码错误")
		}
		return
	}
	p.Cookies = resp.Cookies
	p.isLogin = true
	return
}

// CheckLogin is a func that check cookies whether is valid
// TODO: need further test
func (p *NCM) CheckLogin() (err error) {
	resp, err := resty.New().
		SetHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
		}).
		SetCookies(p.Cookies).
		R().
		Get("https://music.163.com")
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		err = fmt.Errorf("status code is not equal 200, actually the code is  %s", err)
		return
	}

	// get detail
	re1 := regexp.MustCompile("GUser\\s*=\\s*([^;]+);")
	re2 := regexp.MustCompile("GBinds\\s*=\\s*([^;]+);")
	len1 := len(re1.FindAllString(resp.String(), -1))
	len2 := len(re2.FindAllString(resp.String(), -1))
	if len1 < 2 || len2 < 2 {
		err = errors.New("can't match data, might don't login")
	} else {
		p.isLogin = true
	}
	return
}
