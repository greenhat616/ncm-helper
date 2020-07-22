package protocol

import "github.com/pkg/errors"

func notLoginError() error {
	return errors.New("未登录")
}
