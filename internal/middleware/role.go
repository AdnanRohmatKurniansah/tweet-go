package middleware

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	RoleAdmin = "admin"
	RoleModerator = "moderator"
	RoleUser = "user"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	allowed := make(map[string]bool, len(allowedRoles))
	for _, r := range allowedRoles {
		allowed[r] = true
	}

	return func(c *gin.Context) {
		val, exists := c.Get("userRole")
		if !exists {
			utils.ErrorMessage(c, 403, "Role information not found", nil)
			return
		}

		role, ok := val.(string)
		if !ok || !allowed[role] {
			utils.ErrorMessage(c, 403, "You do not have permission to access", nil)
			return
		}

		c.Next()
	}
}