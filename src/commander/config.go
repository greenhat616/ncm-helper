package commander

import (
	config "github.com/a632079/ncm-helper/src/config"
	flag "github.com/spf13/pflag"
)

func registerConfigFlag () {
	flag.StringVarP(&config.File, "config", "c", "", "设定档案信息")
}
