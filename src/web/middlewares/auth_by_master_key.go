package middlewares

import (
	"github.com/a632079/ncm-helper/src/web/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// AuthByMasterKey is a gin middleware that check bearer token if masterKey is set
func AuthByMasterKey() gin.HandlerFunc {
	masterKey := viper.GetString("server.auth.master_key")
	return func(context *gin.Context) {
		if masterKey != "" && !util.ValidTokenByContext(context) {
			util.Fail(context, map[string]interface{}{}, 401)
		} else {
			context.Next()
		}
	}
}
