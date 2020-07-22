package protocol

import "github.com/a632079/ncm-helper/src/protocol/request"

func appendCustomClientIP(input request.Options, ip string) request.Options {
	if ip != "" {
		input.IP = ip
	}
	return input
}
