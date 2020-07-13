package commander

import (
	"github.com/a632079/ncm-helper/src/config"
	flag "github.com/spf13/pflag"
)

func registerDebugFlag () {
	flag.BoolVarP(&config.Debug, "debug", "d", false, "启动调试模式")
}
