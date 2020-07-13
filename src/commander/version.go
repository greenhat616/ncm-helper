package commander

import (
	"fmt"
	config "github.com/a632079/ncm-helper/src/config"
	flag "github.com/spf13/pflag"
	"os"
)

var V bool

func registerVersionFlag() {
	flag.BoolVarP(&V, "version", "v", false, "查看版本信息")
}

func handleVersionFlag() {
	if V {
		fmt.Printf("网易云助手, Authored by a632079\n版本: %s\nMake 版本: %s\nGit Tag: %s\n编译时间: %s\n编译环境: %s\n", config.Version, config.MakeVersion, config.BuildTag, config.BuildTime, config.GoVersion)
		os.Exit(0)
	}
}
