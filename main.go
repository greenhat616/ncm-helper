package main

import (
	"github.com/a632079/ncm-helper/src/commander"
	"github.com/a632079/ncm-helper/src/config"
	"github.com/a632079/ncm-helper/src/prestart"
	"github.com/a632079/ncm-helper/src/web"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
)

var (
	// BuildTag is a commit hash that will be injected in release mode
	BuildTag = "Unknown"
	// BuildTime is a time, when it build, that will be injected in release mode
	BuildTime = "Unknown"
	// MakeVersion is the version of make in release, will be injected in release mode
	MakeVersion = "Unknown"
	// Version is the version of this program, will be injected in release mode
	Version = "development"
)

var r *gin.Engine

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
	prestart.Do()

	if config.Debug {
		log.Info("[debug] 已启用调试模式")
	}

	// init Web Server
	r = web.InitWebServer()
}

func main() {
	// start Server
	r.Run(":" + viper.GetString("server.port"))
}
