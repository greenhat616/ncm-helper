package protocol

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type NCM struct {
	Cookies []*http.Cookie
	Email   string
	IP      string // User Client IPV4
	Pass    string
	MD5Pass string
	Phone   string
	isLogin bool
}

func NewWithEmail(email string, pass string, isMD5Pass bool) *NCM {
	if isMD5Pass {
		return &NCM{
			Email:   email,
			MD5Pass: pass,
		}
	} else {
		return &NCM{
			Email: email,
			Pass:  pass,
		}
	}
}

func NewWithPhone(phone string, pass string, isMD5Pass bool) *NCM {
	if isMD5Pass {
		return &NCM{
			Phone:   phone,
			MD5Pass: pass,
		}
	} else {
		return &NCM{
			Phone: phone,
			Pass:  pass,
		}
	}
}

func NewWithCookies(cookies []*http.Cookie) *NCM {
	return &NCM{
		Cookies: cookies,
	}
}
