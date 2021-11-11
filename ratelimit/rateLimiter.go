package ratelimit

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			//c.String(http.StatusForbidden, "当前访问人数过多，请稍后再试")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, "当前访问人数过多，请稍后再试")
			//c.Abort()
			return
		}
		c.Next()
	}
}
