package commander

import (
	"fmt"
	"github.com/a632079/ncm-helper/src/config"
	flag "github.com/spf13/pflag"
	"os"
)

var h bool

func registerHelpFlag() {
	flag.BoolVarP(&h, "help", "h", false, "查看程序帮助")
}

func handleHelpFlag() {
	if h {
		fmt.Printf(`网易云助手 v%s
使用: ncm [-dhv] [-c filename]

选项：
`, config.Version)
		flag.PrintDefaults()
		os.Exit(0)
	}
}
