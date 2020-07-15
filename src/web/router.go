package web

import (
	"github.com/a632079/ncm-helper/src/config"
	apiV1 "github.com/a632079/ncm-helper/src/web/controllers/api/v1"
	"github.com/a632079/ncm-helper/src/web/middlewares"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/thinkerou/favicon"
	"log"
	"os"
)

func InitWebServer() *gin.Engine {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// load middleware
	r.Use(requestid.New())

	// favicon.ico
	if _, err := os.Stat("../static/resource/favicon.ico"); err == nil {
		r.Use(favicon.New("../static/resource/favicon.ico"))
	} else if _, err := os.Stat("./static/resource/favicon.ico"); err == nil {
		r.Use(favicon.New("./static/resource/favicon.ico"))
	}

	r.Use(middlewares.Cors())
	if !viper.IsSet("server.secret") {
		log.Fatal("[web] can't start server because of the secret is not set.")
	}
	r.Use(middlewares.Session(viper.GetString("server.secret")))

	// Setup router
	r.GET("/", func(context *gin.Context) {
		context.String(200, "Hello, World.")
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", apiV1.Ping)
	}

	return r
}
