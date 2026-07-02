package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiKeyHeader = "X-API-KEY"
const apiKeyValue = "final-marzo-2026-secure"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(apiKeyHeader)

		if apiKey != apiKeyValue {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
