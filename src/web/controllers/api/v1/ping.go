package v1

import (
	"github.com/a632079/ncm-helper/src/web/util"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	util.Success(c, map[string]interface{}{
		"requestId": requestid.Get(c),
	})
}
