package protocol

import "net/http"

type NCM struct {
	Email *string
	Pass *string
	Phone *string
	MD5Pass *string
	Cookies []*http.Cookie
}

func NewWithEmail(email string, pass string, isMD5Pass bool) *NCM {
	if isMD5Pass {
		return &NCM{
			Email: &email,
			MD5Pass: &pass,
		}
	} else {
		return &NCM{
			Email: &email,
			Pass: &pass,
		}
	}
}

func NewWithPhone(phone string, pass string, isMD5Pass bool) *NCM {
	if isMD5Pass {
		return &NCM{
			Phone: &phone,
			MD5Pass: &pass,
		}
	} else {
		return &NCM{
			Phone: &phone,
			Pass: &pass,
		}
	}
}

func NewWithCookies (cookies []*http.Cookie) *NCM {
	return &NCM {
		Cookies: cookies,
	}
}

