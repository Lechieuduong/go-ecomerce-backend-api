package middlewares

import (
	"go-ecomerce-backend-api/m/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func (c * gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorInvalidToken, "")
			c.Abort()
			return 
		}
		c.Next()
	}
}