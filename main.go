package main

import (
	"fmt"
	"github.com/a632079/ncm-helper/src/commander"
	"github.com/a632079/ncm-helper/src/config"
	"github.com/a632079/ncm-helper/src/preStart"
	log "github.com/sirupsen/logrus"
	"runtime"
)

var (
	BuildTag    = "Unknown"
	BuildTime   = "Unknown"
	MakeVersion = "Unknown"
	Version     = "development"
)

func init() {
	// global set build information
	config.BuildTag = BuildTag
	config.BuildTime = BuildTime
	config.GoVersion = runtime.Version()
	config.MakeVersion = MakeVersion
	config.Version = Version

	// Parse Flag
	commander.Parse()

	// Init Drivers
	preStart.Do()

	if config.Debug {
		log.Info("[debug] 已启用调试模式")
	}
}

func main() {
	fmt.Println("Hello, World.")
}
