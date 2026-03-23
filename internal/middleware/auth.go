package middleware

import (
	"strings"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorMessage(c, 403, "Authorization header is required", nil)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.ErrorMessage(c, 403, "The header format must be: Bearer <token>", nil)
			c.Abort()
			return
		}

		claims, err := utils.ValidateJWT(parts[1], jwtSecret)
		if err != nil {
			utils.ErrorMessage(c, 401, "Token is invalid or has expired", err.Error())
			c.Abort()
			return
		}

		c.Set("userId", claims.Id)	
		c.Set("userEmail", claims.Email)
		c.Set("userName", claims.Name)
		c.Set("userPhone", claims.Phone)
		// c.Set("userRole", claims.Role) 

		c.Next()
	}
}