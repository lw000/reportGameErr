package middleware

import (
	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Infof("ClientIP: %s Referer: %s", c.ClientIP(), c.Request.Referer())
		c.Next()
		c.ClientIP()
	}
}
