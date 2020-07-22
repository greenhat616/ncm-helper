package auth

import (
	"github.com/a632079/ncm-helper/src/web/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Check is a gin controller that impl a auth check whether the master password is set
func Check(c *gin.Context) {
	masterKey := viper.GetString("server.auth.master_key")
	if masterKey == "" || util.ValidTokenByContext(c) {
		util.Success(c, map[string]interface{}{})
	} else {
		util.Fail(c, map[string]interface{}{}, 401)
	}
}
