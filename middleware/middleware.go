package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IPWhiteList(whitelist map[string]bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		if !whitelist[context.ClientIP()] {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
		}
	}

}
