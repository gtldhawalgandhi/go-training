package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/token"
)

// Auth checks for valid token in the request
func Auth(t token.Tokener) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		_, err := t.VerifyToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})

			return
		}

		c.Next()
	}
}
