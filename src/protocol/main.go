package protocol

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// NCM is a netease cloud music protocol instance
type NCM struct {
	Cookies     []*http.Cookie
	Email       string
	IP          string // User Client IPV4
	Pass        string
	IsMD5Pass   bool
	Phone       string
	CountryCode string // Phone CountryCode
	isLogin     bool
}

// NewWithEmail is a constructor of NCM by Email
func NewWithEmail(email string, pass string, isMD5Pass bool) *NCM {
	return &NCM{
		Email:     email,
		Pass:      pass,
		IsMD5Pass: isMD5Pass,
	}
}

// NewWithPhone is a constructor of NCM by Phone
func NewWithPhone(phone string, pass string, countryCode string, isMD5Pass bool) *NCM {
	return &NCM{
		Pass:        pass,
		IsMD5Pass:   isMD5Pass,
		Phone:       phone,
		CountryCode: countryCode,
	}
}

// NewWithCookies is a constructor of NCM by Cookies
func NewWithCookies(cookies []*http.Cookie) *NCM {
	return &NCM{
		Cookies: cookies,
	}
}
