package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/devsirose/simplebank/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var RecoveryWithLogger gin.HandlerFunc = func(c *gin.Context) {
	defer func() {
		//entry point when panic
		if r := recover(); r != nil {
			// Log panic + stacktrace
			logger.Log.Error("panic recovered",
				zap.Any("error", r),
				zap.ByteString("stack", debug.Stack()),
				zap.String("path", c.Request.URL.Path),
				zap.String("method", c.Request.Method),
				zap.String("client_ip", c.ClientIP()),
			)

			// Trả 500 JSON
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "internal_server_error",
				"message": "unexpected error",
			})
		}
	}()

	c.Next()
}
