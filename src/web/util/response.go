package util

import (
	"github.com/gin-gonic/gin"
	"time"
)

var errorMessageMap = map[int]string{
	404: "not found specific route",
	400: "bad request",
	500: "server error",
}

// Success is a func that do the common situation of responding successful formation
func Success(ctx *gin.Context, data map[string]interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "ok",
		"data":    data,
		"ts":      time.Now().UnixNano() / 1e6,
	})
}

// Fail is a func that do the common situation of responding failed formation
func Fail(ctx *gin.Context, data map[string]interface{}, code int) {
	var status int
	if code <= 0 {
		status = 200
	} else {
		status = code
	}
	msg, ok := errorMessageMap[code]
	if !ok {
		msg = "unknown status code, please contact author."
	}
	ctx.JSON(status, map[string]interface{}{
		"code":    code,
		"message": msg,
		"data":    data,
		"ts":      time.Now().Nanosecond() / 1e6,
	})
}
