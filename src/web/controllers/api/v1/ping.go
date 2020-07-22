package v1

import (
	"github.com/a632079/ncm-helper/src/web/util"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// Ping is a controller func that impl a Pong response,
// intended to notify the client that server is ok
func Ping(c *gin.Context) {
	util.Success(c, map[string]interface{}{
		"requestId": requestid.Get(c),
	})
}
