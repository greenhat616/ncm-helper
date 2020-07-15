package request

import (
	"fmt"
	"github.com/a632079/ncm-helper/src/util/pad"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func performRequest(r *resty.Request, method string, url string) (resp *resty.Response, err error) {
	method = strings.ToUpper(method)
	if method == "GET" {
		resp, err = r.Get(url)
		return
	} else if method == "POST" {
		resp, err = r.Post(url)
		return
	} else {
		err = errors.New("mismatch method type")
		return
	}
}

func handleResponse(resp *resty.Response) (response *APIResponse, err error) {
	if resp.StatusCode() == 200 {
		response = &APIResponse{
			StatusCode: resp.StatusCode(),
			Cookies:    resp.Cookies(),
			Data:       resp.Body(),
		}
		return
	} else {
		err = errors.New(fmt.Sprintf("request failed, status code: "+strconv.Itoa(resp.StatusCode())+", data: %s", resp.Request.Body))
		return
	}
}

func preFillHeader(method string, url string, options Options) (map[string]string, string) {
	headers := map[string]string{}
	re := regexp.MustCompile("/\\w*api/")
	if options.Crypto == "weapi" {
		headers["User-Agent"] = chooseUserAgent(options.UA)
		url = re.ReplaceAllString(url, "weapi")
	} else if options.Crypto == "eapi" {
		// no need UserAgent
		url = re.ReplaceAllString(url, "eapi")
	} else {
		headers["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36"
		url = re.ReplaceAllString(url, "api")
	}
	if strings.ToUpper(method) == "POST" {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	if strings.Contains(url, "music.163.com") {
		headers["Content-Type"] = "https://music.163.com"
	}
	if options.IP != nil {
		headers["X-Real-IP"] = *options.IP
	}
	return headers, url
}

func setDefaultValue(headers map[string]string, name string, value string) map[string]string {
	_, ok := headers[name]
	if !ok {
		headers[name] = value
	}
	return headers
}

func genRequestId() string {
	ms := time.Now().UnixNano() / 1e6
	rand.Seed(time.Now().Unix())
	r := strconv.Itoa(rand.Intn(1000))
	r = pad.Left(r, 4, "0")
	return fmt.Sprintf("%v_%s", ms, r)
}
