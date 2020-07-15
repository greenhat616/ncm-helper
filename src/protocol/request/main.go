package request

import (
	"github.com/a632079/ncm-helper/src/protocol/crypto"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

func chooseUserAgent(ua *string) string {
	index := 0
	rand.Seed(time.Now().UnixNano())
	if *ua == "mobile" {
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
	UA      *string
	IP      *string // X-Real-IP IPV4
	Crypto  string
	URL     *string // eapi is needed
}

type APIResponse struct {
	StatusCode int
	Cookies    []*http.Cookie
	Data       []byte
}

func CreateWEAPIRequest(method string, url string, data map[string]interface{}, options Options) (response *APIResponse, err error) {
	headers, url := preFillHeader(method, url, options)
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
	params, encSecKey, err := crypto.WEAPI(d)

	client := resty.New()
	r := client.R().
		SetQueryParams(headers).
		SetCookies(options.Cookies).
		SetBody(map[string]interface{}{
			"params":    params,
			"encSecKey": encSecKey,
		})
	resp, err := performRequest(r, method, url)
	if err != nil {
		return
	}
	return handleResponse(resp)
}

func CreateEAPIRequest(method string, url string, data map[string]interface{}, options Options) (response *APIResponse, err error) {
	// check options
	if options.URL == nil {
		err = errors.New("request failed: url in options is not set")
		return
	}
	_, url = preFillHeader(method, url, options)
	headers := map[string]string{}

	// set headers
	for _, v := range options.Cookies {
		switch v.Name {
		case "__csrf",
			"deviceId",
			"appver",
			"versioncode",
			"mobilename",
			"buildver",
			"resolution",
			"os",
			"channel",
			"MUSIC_U",
			"MUSIC_A":
			headers[v.Name] = v.Value
		}
	}

	// set Default value
	headers = setDefaultValue(headers, "appver", "6.1.1")
	headers = setDefaultValue(headers, "versioncode", "140")
	headers = setDefaultValue(headers, "buildver", string(time.Now().Unix()))
	headers = setDefaultValue(headers, "resolution", "1920x1080")
	headers = setDefaultValue(headers, "os", "android")
	headers["requestId"] = genRequestId()

	// encrypted data
	data["headers"] = headers
	raw, err := json.Marshal(data)
	if err != nil {
		return
	}
	params, err := crypto.EAPI(url, raw)
	if err != nil {
		return
	}

	// perform request
	r := resty.New().
		R().
		SetHeaders(headers).
		SetCookies(options.Cookies).
		SetBody(map[string]interface{}{
			"params": params,
		})
	resp, err := performRequest(r, method, url)
	if err != nil {
		return
	}

	// handle Response
	response, err = handleResponse(resp)
	if err != nil {
		return
	}
	decrypted, err := crypto.Decrypt(response.Data)
	if err == nil {
		response.Data = decrypted
	}
	return
}

func CreateLinuxApiRequest(method string, url string, data map[string]interface{}, options Options) (response *APIResponse, err error) {
	headers, url := preFillHeader(method, url, options)
	raw := map[string]interface{}{
		"method": method,
		"url":    url,
		"params": data,
	}
	d, err := json.Marshal(raw)
	if err != nil {
		return
	}
	// encrypt data
	eParams, err := crypto.LinuxAPI(d)
	if err != nil {
		return
	}
	url = "https://music.163.com/api/linux/forward"

	r := resty.New().
		R().
		SetHeaders(headers).
		SetCookies(options.Cookies).
		SetBody(map[string]interface{}{
			"eparams": eParams,
		})
	resp, err := performRequest(r, method, url)
	if err != nil {
		return
	}
	return handleResponse(resp)
}

func CreateRequest(method string, url string, data map[string]interface{}, options Options) (response *APIResponse, err error) {
	switch options.Crypto {
	case "weapi":
		return CreateRequest(method, url, data, options)
	case "linuxapi":
		return CreateLinuxApiRequest(method, url, data, options)
	case "eapi":
		return CreateEAPIRequest(method, url, data, options)
	default:
		err = errors.New("mismatch crypto type")
		return
	}
}
