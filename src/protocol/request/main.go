package request

import (
	"encoding/json"
	"github.com/a632079/ncm-helper/src/protocol/crypto"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	userAgentList = []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	}
)

func chooseUserAgent (ua *string) string {
	index := 0
	rand.Seed(time.Now().UnixNano())
	if *ua == "mobile"{
		index = rand.Intn(7)
	} else if *ua == "pc" {
		index = rand.Intn(5) + 8
	} else {
		index = rand.Intn(len(userAgentList))
	}
	return userAgentList[index]
}

type Options struct {
	Cookies []*http.Cookie
	UA *string
}

func CreateWEAPIRequest (method string, url string, data map[string]interface{}, options Options) (err) {
	headers := map[string]string {
		"User-Agent": chooseUserAgent(options.UA),
	}
	if strings.ToUpper(method) == "POST" {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	if strings.Contains(url, "music.163.com") {
		headers["Content-Type"] = "https://music.163.com"
	}
	// set CSRF Token
	for _, v := range options.Cookies {
		if v.Name == "_csrf" {
			data["csrf_token"] = v.Value
		}
	}
	// encrypt data
	d, err := json.Marshal(data)
	if err != nil {
		return
	}
	crypto.WEAPI(d)
	client := resty.New()
	client.
		R().
		SetQueryParams(headers).
		SetCookies(options.Cookies)
}

func CreateEAPIRequest () {

}

func CreateLinuxApiRequest () {

}

