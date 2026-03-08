package middleware

import (
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServiceTokenAuth validates requests using a static service token
// passed in the X-Service-Token header. Use this for service-to-service
// authentication where JWT user context is not needed.
func ServiceTokenAuth(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		provided := c.GetHeader("X-Service-Token")
		if provided == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing service token"})
			return
		}

		if subtle.ConstantTimeCompare([]byte(provided), []byte(token)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid service token"})
			return
		}

		c.Set("auth_type", "service_token")
		c.Next()
	}
}
