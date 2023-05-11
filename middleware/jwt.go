package middleware

import (
	"booking-flight-system/jwt"
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

func ExtractJWT(service *jwt.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		// Bearer <<token>>
		if header == "" {
			c.Next()
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		email, err := service.ValidateToken(token)
		if err != nil {
			c.Next()
			return
		}
		c.Set("user_email", email)

		ctx := context.WithValue(c.Request.Context(), "user_email", email)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
